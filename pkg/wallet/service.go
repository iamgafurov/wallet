package wallet

import (
	"github.com/iamgafurov/wallet/pkg/types"
	"github.com/google/uuid"
	"errors"
	"os"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"sync"
	"fmt"
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
	buf := make([]byte,4096)
	read,err:= file.Read(buf)
	if err != nil{
		log.Print(err)
		return err
	}
	accounts := strings.Split(string(buf[:read]),"|")
	accounts = accounts[:len(accounts)-1]
	for _,account := range accounts{
		val := strings.Split(account,";")
		id, err := strconv.ParseInt(val[0],10,64)
		if err != nil {
			return err
		}
		balance,err := strconv.ParseInt(val[2],10,64)
		if err != nil {
			return err
		}
		s.accounts = append(s.accounts,&types.Account{
			ID:      id,
			Phone:   types.Phone(val[1]),
			Balance: types.Money(balance),
		})
	}
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


func (s *Service) SumPayments(gorutines int)types.Money{
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	sum:=int64(0)
	kol:=0
	i:=0
	if gorutines == 0{
		kol= len(s.payments)
	}else{
		kol= int(len(s.payments)/gorutines)
	}
	for i=0;i<gorutines-1;i++{
		wg.Add(1)
		go func (index int){
			defer wg.Done()
			val:=int64(0)
			payments := s.payments[index*kol:(index+1) * kol]
			for _,payment := range payments{
				val +=  int64(payment.Amount)
			}
			mu.Lock()
			sum+= val;
			mu.Unlock()

		}(i)
	}
	wg.Add(1)
		go func (){
			defer wg.Done()
			val:=int64(0)
			payments := s.payments[i*kol:]
			for _,payment := range payments{
				val +=  int64(payment.Amount)
			}
			mu.Lock()
			sum+= val;
			mu.Unlock()

		}()
	wg.Wait()	
	return types.Money(sum)
}

func (s *Service)Export(dir string)error {
	
	if len(s.accounts) != 0 {
		_,err := os.Create(dir +"/accounts.dump")
		if err != nil {
			log.Print(err)
			return err
		}
		accountsData:= ""
		for _,account:= range s.accounts{
			accountsData += strconv.FormatInt(account.ID,10) + ";" + string(account.Phone) + ";" + strconv.FormatInt(int64(account.Balance),10)+ "|"
		}
		err = ioutil.WriteFile(dir + "/accounts.dump", []byte(accountsData), 0666)
	}
	if len(s.favorites) != 0{
		_,err := os.Create(dir + "/favorites.dump")
		if err != nil {
			log.Print(err)
			return err
		}
		favoritesData:= ""
		for _,favorite:= range s.favorites{
			favoritesData += favorite.ID + ";" + strconv.FormatInt(favorite.AccountID,10) + ";" + favorite.Name + ";" + strconv.FormatInt(int64(favorite.Amount),10)+ ";" + string(favorite.Category) + "|"
		}
		err = ioutil.WriteFile(dir + "/favorites.dump", []byte(favoritesData), 0666)
		if err != nil {
			log.Print(err)
			return err
		}
	}
	if len(s.payments)!= 0{
		_,err := os.Create(dir +"/payments.dump")
		if err != nil {
			log.Print(err)
			return err
		}
		
		paymentsData:= ""
		for _,payment:= range s.payments{
			paymentsData += payment.ID +";"+ strconv.FormatInt(payment.AccountID,10)  + ";" + strconv.FormatInt(int64(payment.Amount),10)+ ";" + string(payment.Category)+ ";" + string(payment.Status)+"|"
		}
		err = ioutil.WriteFile(dir + "/payments.dump", []byte(paymentsData), 0666)
		if err != nil {
			log.Print(err)
			return err
		}
	}
	return nil

}


func (s *Service) Import(dir string)error {
	

	accountsData,err := ioutil.ReadFile(dir + "/accounts.dump")
	if err == nil{
		accounts := strings.Split(string(accountsData),"|")
		accounts = accounts[:len(accounts)-1]
		for _,account := range accounts{
			val := strings.Split(account,";")
			id, err := strconv.ParseInt(val[0],10,64)
			if err != nil {
				return err
			}
			balance,err := strconv.ParseInt(val[2],10,64)
			if err != nil {
				return err
			}
			acc,err := s.FindAccountByID(id)
			if err == nil{
				acc.Phone = types.Phone(val[1])
				acc.Balance = types.Money(balance)
			}else {
				s.accounts = append(s.accounts,&types.Account{
					ID:      id,
					Phone:   types.Phone(val[1]),
					Balance: types.Money(balance),
				})
			}
		}
	}

	paymentsData,err := ioutil.ReadFile(dir + "/payments.dump")
	if err == nil{
		payments := strings.Split(string(paymentsData),"|")
		payments = payments[:len(payments)-1]
		for _,payment := range payments{
			val := strings.Split(payment,";")
			id := val[0]
			accountID, err := strconv.ParseInt(val[1],10,64)
			if err != nil {
				return err
			}
			amount,err := strconv.ParseInt(val[2],10,64)
			if err != nil {
				return err
			}
			category := types.PaymentCategory(val[3])
			status := types.Status(val[4])
			pay,err := s.FindPaymentByID(id)
			if err == nil{
				pay.Amount = types.Money(amount)
				pay.Category = category
				pay.Status = status
			}else {
				s.payments = append(s.payments,&types.Payment{
					ID:      id,
					AccountID:   accountID,
					Amount: types.Money(amount),
					Category: category,
					Status: status,
				})
			}
		}
	}


	favoritesData,err := ioutil.ReadFile(dir + "/favorites.dump")
	if err == nil{
		favorites := strings.Split(string(favoritesData),"|")
		favorites = favorites[:len(favorites)-1]
		for _,favorite := range favorites{
			val := strings.Split(favorite,";")
			id := val[0]
			accountID, err := strconv.ParseInt(val[1],10,64)
			if err != nil {
				return err
			}
			name := val[2]
			amount,err := strconv.ParseInt(val[3],10,64)
			if err != nil {
				return err
			}
			category := types.PaymentCategory(val[4])
			pay,err := s.FindFavoriteByID(id)
			if err == nil{
				pay.Amount = types.Money(amount)
				pay.Category = category
				pay.Name = name
			}else {
				s.favorites = append(s.favorites,&types.Favorite{
					ID:      id,
					AccountID:   accountID,
					Name: name,
					Amount: types.Money(amount),
					Category: category,
				})
			}
		}
	}
	return nil

}

func (s *Service) ExportAccountHistory(accountID int64) ([]types.Payment,error){
	result := []types.Payment{}
	for _,payment := range s.payments {
		if payment.AccountID == accountID {
			result = append(result,*payment)
		}
	}
	if len(result) ==0 {
		return nil,ErrAccountNotFound
	}
	return result,nil
}

func (s *Service) HistoryToFiles(payments []types.Payment,dir string ,records int) error{
	num :=1
	i:=0
	if len(payments) == 0{
		return nil
	}
	if len(payments) <= records {
		err:= s.ExportHistoryToFile(payments,dir+"/payments" + ".dump")
		if err != nil {
			log.Print(err)
			return err
		}else {
			return nil
		}
	}
	for i=0;i + records<len(payments);i+=records {
		err:= s.ExportHistoryToFile(payments[i:i+records],dir+"/payments"+strconv.Itoa(num) + ".dump")
		if err != nil {
			log.Print(err)
			return err
		}
		num++
	}
	err:= s.ExportHistoryToFile(payments[i:],dir+"/payments"+strconv.Itoa(num) + ".dump")
		if err != nil {
			log.Print(err)
			return err
		}
	return nil
}

func (s *Service) ExportHistoryToFile(payments []types.Payment, name string)error{
	_,err := os.Create(name)
		if err != nil {
			log.Print(err)
			return err
		}
	paymentsData:= ""
	for _,payment:= range payments{
		paymentsData += payment.ID +";"+ strconv.FormatInt(payment.AccountID,10)  + ";" + strconv.FormatInt(int64(payment.Amount),10)+ ";" + string(payment.Category)+ ";" + string(payment.Status)+"\n"
	}
	err = ioutil.WriteFile(name, []byte(paymentsData), 0666)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil

}


func (s *Service) FilterPayments(accountID int64,gorutines int)([]types.Payment,error){
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	pays := []types.Payment{}
	kol:=0
	i:=0
	if gorutines == 0{
		kol= len(s.payments)
	}else{
		kol= int(len(s.payments)/gorutines)
	}
	for i=0;i<gorutines-1;i++{
		wg.Add(1)
		go func (index int){
			defer wg.Done()
			result:=[]types.Payment{}
			payments := s.payments[index*kol:(index+1) * kol]
			for _,payment := range payments{
				if payment.AccountID == accountID {
					result = append(result,*payment)
				}
			}
			mu.Lock()
			pays = append(pays,result...)
			mu.Unlock()
		}(i)
	}
	wg.Add(1)
		go func (){
			defer wg.Done()
			result:=[]types.Payment{}
			payments := s.payments[i*kol:]
			for _,payment := range payments{
				if payment.AccountID == accountID {
					result = append(result,*payment)
				}
			}
			mu.Lock()
			pays = append(pays,result...)
			mu.Unlock()

		}()
	wg.Wait()	
	if len(pays) == 0{
		return nil,ErrAccountNotFound
	}
	return pays,nil
}


func (s *Service) FilterPaymentsByFn(filter func(payment types.Payment)bool,gorutines int)([]types.Payment,error){
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	pays := []types.Payment{}
	kol:=0
	i:=0 
	if gorutines == 0{
		kol= len(s.payments)
	}else{
		kol= int(len(s.payments)/gorutines)
	}
	for i=0;i<gorutines-1;i++{
		wg.Add(1)
		go func (index int){
			defer wg.Done()
			result:=[]types.Payment{}
			payments := s.payments[index*kol:(index+1) * kol]
			for _,payment := range payments{
				if filter(*payment) {
					result = append(result,*payment)
				}
			}
			mu.Lock()
			pays = append(pays,result...)
			mu.Unlock()
		}(i)
	}
	wg.Add(1)
		go func (){
			defer wg.Done()
			result:=[]types.Payment{}
			payments := s.payments[i*kol:]
			for _,payment := range payments{
				if filter(*payment) {
					result = append(result,*payment)
				}
			}
			mu.Lock()
			pays = append(pays,result...)
			mu.Unlock()

		}()
	wg.Wait()	
	if len(pays) == 0{
		return nil,ErrAccountNotFound
	}
	return pays,nil
}


func (s *Service) SumPaymentsWithProgress() <- chan types.Progress{
	wg := sync.WaitGroup{}
	ch := make(chan types.Progress,1)
	defer close(ch)
	partsLen := 100_000
	i:=0
	var sum types.Progress
	gorutines := int(len(s.payments)/partsLen)
	for i=0;i<gorutines-1;i++{
		wg.Add(1)
		go func(index int){
			defer wg.Done()
			val:=int64(0)
			payments := s.payments[index*partsLen:(index+1) * partsLen]
			for _,payment := range payments{
				val +=  int64(payment.Amount)
			}
			select {
			case sum = <-ch:
				fmt.Println("received Progress", sum)
			default:
				fmt.Println("no Progress received")
			}
			ch <- types.Progress{Part: 0, Result: sum.Result + types.Money(val)}
		}(i)
	}

	wg.Add(1)
		go func(){
			defer wg.Done()
			val:=int64(0)
			payments := s.payments[i*partsLen:]
			for _,payment := range payments{
				val +=  int64(payment.Amount)
			}
			select {
			case sum = <-ch:
				fmt.Println("received Progress", sum)
			default:
				fmt.Println("no Progress received")
			}
			ch <- types.Progress{Part: 0, Result: sum.Result + types.Money(val)}
		}()
	wg.Wait()	
	return ch
}