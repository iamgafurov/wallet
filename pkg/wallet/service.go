package wallet

import (
	"github.com/iamgafurov/wallet/pkg/types"
	"github.com/google/uuid"
	"errors"
	)


var ErrAccountNotFound = errors.New("account not found")
var ErrPhoneExist = errors.New("acount exist")
var ErrNotEnoughBalance = errors.New("not enaugh balance")
var ErrAmountMustBePositive = errors.New("amount must be greater than zero")
var ErrPaymentNotFound = errors.New("payment not found")
	

type Service struct {
	accounts []*types.Account
	payments []*types.Payment
}

func (s *Service) RegisterAccount(phone types.Phone)(*types.Account,error){
	nextID := int64(0)
	for _,account := range s.accounts{
		if account.Phone == phone{
			return nil, ErrPhoneExist
		}
	}
	nextID++
	newAccount := &types.Account{
		ID: nextID,
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