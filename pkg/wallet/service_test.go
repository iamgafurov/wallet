package wallet
import (
	"github.com/iamgafurov/wallet/pkg/types"
	"testing"
	//"fmt"
	
	)
		
func TestService_Repeat_success(t *testing.T){
	srv := &Service{}
	srv.RegisterAccount(types.Phone("9099343490"))
	account,err := srv.FindAccountByID(1)
	if err != nil {
		t.Errorf("Repeat(): cant't register account,error = %v",err)
	}
	err = srv.Deposit(account.ID , types.Money(30_000))
	if err != nil{
		t.Errorf("Repeat(): cant't register account,error = %v",err)
	}
	payment,er := srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Repeat(): cant't register account,error = %v",err)
	}
	payment,err =srv.Repeat(payment.ID)
	if err !=nil {
		t.Errorf("Repeat(): cant't register account,error = %v",err)
	}
}

func TestService_Repeat_fail(t *testing.T){
	srv := &Service{}
	srv.RegisterAccount(types.Phone("90999390"))
	account,err := srv.FindAccountByID(1)
	if err != nil {
		t.Errorf("Repeat(): cant't register account,error = %v",err)
	}
	err = srv.Deposit(account.ID , types.Money(30_000))
	if err != nil{
		t.Errorf("Repeat(): cant't register account,error = %v",err)
	}
	payment,er := srv.Pay(account.ID,types.Money(25000),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Repeat(): cant't register account,error = %v",err)
	}
	payment,err =srv.Repeat(payment.ID)
	if err ==nil {
		t.Errorf("Repeat(): cant't register account,error = %v",err)
	}
}


func TestService_FavoritePayment_success(t *testing.T){
	srv := &Service{}
	srv.RegisterAccount(types.Phone("90999390"))
	account,err := srv.FindAccountByID(1)
	if err != nil {
		t.Errorf("Favorite(): cant't register account,error = %v",err)
	}
	err = srv.Deposit(account.ID , types.Money(30_000))
	if err != nil{
		t.Errorf("Favorite(): deposit,error = %v",err)
	}
	payment,er := srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Favorite(): cant't pay,error = %v",err)
	}
	_,err =srv.FavoritePayment(payment.ID, "test name")
	if err !=nil {
		t.Errorf("Favorite(): cant't Favorite,error = %v",err)
	}
}

func TestService_PayFromFavorite_success(t *testing.T){
	srv := &Service{}
	srv.RegisterAccount(types.Phone("90999390"))
	account,err := srv.FindAccountByID(1)
	if err != nil {
		t.Errorf("Favorite(): cant't register account,error = %v",err)
	}
	err = srv.Deposit(account.ID , types.Money(30_000))
	if err != nil{
		t.Errorf("Favorite(): deposit,error = %v",err)
	}
	payment,er := srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Favorite(): cant't pay,error = %v",err)
	}
	favorite,er :=srv.FavoritePayment(payment.ID, "test name")
	if er !=nil {
		t.Errorf("Favorite(): cant't Favorite,error = %v",err)
	}
	_,err =srv.PayFromFavorite(favorite.ID)
	if err !=nil {
		t.Errorf("Favorite(): cant't Favorite,error = %v",err)
	} 
}


func TestService_FindFavoriteByID_success(t *testing.T){
	srv := &Service{}
	srv.RegisterAccount(types.Phone("90999390"))
	account,err := srv.FindAccountByID(1)
	if err != nil {
		t.Errorf("Favorite(): cant't register account,error = %v",err)
	}
	err = srv.Deposit(account.ID , types.Money(30_000))
	if err != nil{
		t.Errorf("Favorite(): deposit,error = %v",err)
	}
	payment,er := srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Favorite(): cant't pay,error = %v",err)
	}
	favorite,er :=srv.FavoritePayment(payment.ID,"test name")
	if er !=nil {
		t.Errorf("Favorite(): cant't Favorite,error = %v",err)
	}
	_,err =srv.FindFavoriteByID(favorite.ID)
	if err !=nil {
		t.Errorf("Favorite(): cant't Favorite,error = %v",err)
	} 
}


func TestService_ExportToFile_success(t *testing.T){
	srv := &Service{}
	srv.RegisterAccount(types.Phone("90999390"))
	srv.RegisterAccount(types.Phone("909993210"))
	srv.RegisterAccount(types.Phone("90999323490"))
	//srv.RegisterAccount(types.Phone("9099939220"))
	_,err := srv.FindAccountByID(1)
	if err != nil {
		t.Errorf("Favorite(): cant't register account,error = %v",err)
	}
	
	err = srv.ExportToFile("dump.txt")
	if err !=nil {
		t.Errorf("ExporrtToFile():,error = %v",err)
	}
}


func TestService_ImportFromFile_success(t *testing.T){
	srv := &Service{}
	err := srv.ImportFromFile("dump.txt")
	if err !=nil {
		t.Errorf("ExporrtToFile():,error = %v",err)
	}
	_,err = srv.FindAccountByID(1)

	if err !=nil {
		t.Errorf("ExporrtToFile():,error = %v",err)
	}
}

func TestService_Export_success(t *testing.T){
	srv := &Service{}
	srv.RegisterAccount(types.Phone("90999390"))
	srv.RegisterAccount(types.Phone("909993210"))
	srv.RegisterAccount(types.Phone("90999323490"))
	srv.RegisterAccount(types.Phone("9099939220"))
	account,err := srv.FindAccountByID(1)
	if err != nil {
		t.Errorf("Favorite(): cant't register account,error = %v",err)
	}
	err = srv.Deposit(account.ID , types.Money(30_000))
	if err != nil{
		t.Errorf("Favorite(): deposit,error = %v",err)
	}
	payment,er := srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Favorite(): cant't pay,error = %v",err)
	}
	favorite,er :=srv.FavoritePayment(payment.ID,"test name")
	if er !=nil {
		t.Errorf("Favorite(): cant't Favorite,error = %v",err)
	}
	_,err =srv.FindFavoriteByID(favorite.ID)
	if err !=nil {
		t.Errorf("Favorite(): cant't Favorite,error = %v",err)
	} 
	
	err = srv.Export("../")
	if err !=nil {
		t.Errorf("ExporrtToFile():,error = %v",err)
	}
}

func TestService_Import_success(t *testing.T){
	srv := &Service{}
	srv.RegisterAccount(types.Phone("90999390"))
	srv.RegisterAccount(types.Phone("909993210"))
	srv.RegisterAccount(types.Phone("90999323490"))
	srv.RegisterAccount(types.Phone("9099939220"))
	account,err := srv.FindAccountByID(1)
	if err != nil {
		t.Errorf("Favorite(): cant't register account,error = %v",err)
	}
	err = srv.Deposit(account.ID , types.Money(30_000))
	if err != nil{
		t.Errorf("Favorite(): deposit,error = %v",err)
	}
	payment,er := srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Favorite(): cant't pay,error = %v",err)
	}
	favorite,er :=srv.FavoritePayment(payment.ID,"test name")
	if er !=nil {
		t.Errorf("Favorite(): cant't Favorite,error = %v",err)
	}
	_,err =srv.FindFavoriteByID(favorite.ID)
	if err !=nil {
		t.Errorf("Favorite(): cant't Favorite,error = %v",err)
	} 
	
	err = srv.Import("../")
	if err !=nil {
		t.Errorf("ExporrtToFile():,error = %v",err)
	}
	//fmt.Print(*srv.accounts[3],*srv.payments[0])
}