package wallet

import (
	"github.com/iamgafurov/wallet/pkg/types"
	"github.com/google/uuid"
	"errors"
	"io"
	"os"
	"log"
	"strconv"
	"encoding/json"
	)


var ErrAccountNotFound = errors.New("account not found")
var ErrPhoneExist = errors.New("acount exist")
var ErrNotEnoughBalance = errors.New("not enaugh balance")
var ErrAmountMustBePositive = errors.New("amount must be greater than zero")
var ErrPaymentNotFound = errors.New("payment not found")
var ErrFavoriteNotFound = errors.New("favorite not found")	

type Service struct {
	accounts []*types.Account
	payments []*types.Payment
	favorites []*types.Favorite
	nextID  int64 
}

func (s *Service) RegisterAccount(phone types.Phone)(*types.Account,error){
	
	for _,account := range s.accounts{
		if account.Phone == phone{
			return nil, ErrPhoneExist
		}
	}
	s.nextID++
	newAccount := &types.Account{
		ID: s.nextID,
		Phone: phone,
		Balance: 0,
	}
	s.accounts = append(s.accounts,newAccount)
	return newAccount,nil
}

func (s *Service) FindAccountByID(accountID int64)(*types.Account,error){
	for _,account := range s.accounts {
		if (account).ID == accountID {
			return account,nil
		}
	}
	return nil, ErrAccountNotFound
}

func (s *Service) Deposit(accountID int64, amount types.Money) error {
	if amount <=0 {
		return ErrAmountMustBePositive
	}

	var account *types.Account

	for _,acc := range s.accounts{
		if acc.ID == accountID {
			account = acc 
			break
		}
	}
	if account == nil {
		return ErrAccountNotFound
	}

	account.Balance += amount
	return nil
}

func (s *Service) Pay(accountID int64, amount types.Money, category types.PaymentCategory) (*types.Payment, error) {
	if amount <=0 {
		return nil, ErrAmountMustBePositive
	}

	var account *types.Account

	for _,acc := range s.accounts{
		if acc.ID == accountID {
			account = acc 
			break
		}
	}
	if account == nil {
		return nil, ErrAccountNotFound
	}
	if account.Balance < amount {
		return nil, ErrNotEnoughBalance
	}
	account.Balance -= amount
	paymentID := uuid.New().String()
	payment := &types.Payment{
		ID: paymentID,
		AccountID: accountID, 
		Amount: amount,
		Category: category,
		Status: types.PaymentStatusInProgress,
	}

	s.payments = append(s.payments,payment)
	return payment,nil 
}
func (s *Service)FindPaymentByID(paymentID string)(*types.Payment,error){
	for _,payment := range s.payments {
		if payment.ID == paymentID {
			return payment,nil
		}
	}
	return nil,ErrPaymentNotFound
}
func (s *Service)Repeat(paymentID string)(*types.Payment,error){
	payment,err := s.FindPaymentByID(paymentID)
	if err !=nil {
		return nil,err
	}
	payment,err = s.Pay(payment.AccountID,payment.Amount,payment.Category)
	if err != nil{
		return nil,err
	} 
	return payment,nil
}

func (s *Service)Reject(paymentID string)error{
	payment,err := s.FindPaymentByID(paymentID)
	if err !=nil {
		return err
	}
	account,er := s.FindAccountByID(payment.AccountID)
	if er != nil {
		return err
	}
	account.Balance += payment.Amount
	payment.Status = types.PaymentStatusFail
	return nil
}

func (s *Service) FavoritePayment(paymentID string, name string)(*types.Favorite,error){
	payment,err := s.FindPaymentByID(paymentID)
	if err != nil {
		return nil,err
	}
	favorite := &types.Favorite {
		ID: 		uuid.New().String(),
		AccountID: payment.AccountID,
		Name :	name,
		Amount:	payment.Amount,
		Category:	payment.Category,
	}
	s.favorites = append(s.favorites,favorite)
	return favorite ,nil
}

func (s *Service) FindFavoriteByID(favoriteID string)(*types.Favorite ,error){
	for _,favorite:= range s.favorites{
		if favorite.ID == favoriteID {
			return favorite,nil
		} 
	}
	return nil, ErrFavoriteNotFound
}

func (s *Service) PayFromFavorite(favoriteID string) (*types.Payment,error){
	favorite,err := s.FindFavoriteByID(favoriteID)
	if err != nil {
		return nil,err
	}
	payment,err := s.Pay(favorite.AccountID,favorite.Amount,favorite.Category)
	if err != nil{
		return nil,err
	}
	return payment,nil
}

func (s *Service) ImportFromFile(path string) error{
	file,err := os.Open(path)
	if err != nil{
		return err
	}	

	content := make([]byte,8096)
	buf := make([]byte,4096)
	for {
		read,err:= file.Read(buf)
		
		if err == io.EOF {
			break
		}
		if err != nil{
			log.Print(err)
			return err
		}
		content = append(content,buf[:read]...)
	}
	var dat map[string]interface{}
	
	err= json.Unmarshal(content,&dat)
	if err != nil {
		print(err,"3234324")
		return err
	}
	log.Print(dat)
	return nil
}

func (s *Service) ExportToFile(path string) error{
	_,err := os.Create(path)
	if err !=nil {
		log.Print(err)
		return err
	}
	file,err := os.OpenFile(path,os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil{
		log.Print(err)
		return err
	}
	data:= ""
	for _,account:= range s.accounts{
		data += strconv.FormatInt(account.ID,10) + ";" + string(account.Phone) + ";" + strconv.FormatInt(int64(account.Balance),10)+ "|"
	}
	_,err = file.Write([]byte(data)) 
	if err != nil {
		log.Print(err)
		return err
	}
	defer func(){
		err := file.Close()
		if err != nil {
			log.Print(err)
		}
	}()
	return nil
}