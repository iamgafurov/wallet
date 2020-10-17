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

func TestService_ExportAccountHistory_success(t *testing.T){
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
	
	_,err = srv.ExportAccountHistory(int64(1))
	if err !=nil {
		t.Errorf("ExporrtToFile():,error = %v",err)
	}
	//print(payments[0])
	//fmt.Print(*srv.accounts[3],*srv.payments[0])
}


func TestService_HistoryToFiles_success(t *testing.T){
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
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
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
	payments,err := srv.ExportAccountHistory(int64(1))
	if err !=nil {
		t.Errorf("ExporrtToFile():,error = %v",err)
	}
	err = srv.HistoryToFiles(payments,"../",8)
	if err !=nil {
		t.Errorf("ExporrtToFile():,error = %v",err)
	}
	//print(payments[0])
	//fmt.Print(*srv.accounts[3],*srv.payments[0])

}

func TestService_SumPayments_success(t *testing.T){
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
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
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
	_,err = srv.ExportAccountHistory(int64(1))
	if err !=nil {
		t.Errorf("ExporrtToFile():,error = %v",err)
	}
	sum := srv.SumPayments(0)
	if sum != types.Money(800) {
		t.Errorf("Incorrect result, want %v, got: %v",800,sum)
	}
	//print(payments[0])
	//fmt.Print(*srv.accounts[3],*srv.payments[0])

}



func BenchmarkSumPayments(b *testing.B){
	srv := &Service{}
	srv.RegisterAccount(types.Phone("90999390"))
	srv.RegisterAccount(types.Phone("909993210"))
	srv.RegisterAccount(types.Phone("90999323490"))
	srv.RegisterAccount(types.Phone("9099939220"))
	account,err := srv.FindAccountByID(1)
	if err != nil {
		b.Fatalf("Favorite(): cant't register account,error = %v",err)
	}
	err = srv.Deposit(account.ID , types.Money(30_000))
	if err != nil{
		b.Fatalf("Favorite(): deposit,error = %v",err)
	}
	payment,er := srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		b.Fatalf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		b.Fatalf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		b.Fatalf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		b.Fatalf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		b.Fatalf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		b.Fatalf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		b.Fatalf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		b.Fatalf("Favorite(): cant't pay,error = %v",err)
	}
	
	favorite,er :=srv.FavoritePayment(payment.ID,"test name")
	if er !=nil {
		b.Fatalf("Favorite(): cant't Favorite,error = %v",err)
	}
	_,err =srv.FindFavoriteByID(favorite.ID)
	if err !=nil {
		b.Fatalf("Favorite(): cant't Favorite,error = %v",err)
	} 
	_,err = srv.ExportAccountHistory(int64(1))
	if err !=nil {
		b.Fatalf("ExporrtToFile():,error = %v",err)
	}

	for i:=0;i<b.N;i++{
		sum:= srv.SumPayments(0)
		if sum != types.Money(800) {
			b.Fatalf("Incorrect result, want %v, got: %v",800,sum)
		} 
	} 
	
}



func BenchmarkFilterPayments(b *testing.B){
	srv := &Service{}
	srv.RegisterAccount(types.Phone("90999390"))
	srv.RegisterAccount(types.Phone("909993210"))
	srv.RegisterAccount(types.Phone("90999323490"))
	srv.RegisterAccount(types.Phone("9099939220"))
	account,err := srv.FindAccountByID(1)
	if err != nil {
		b.Fatalf("Favorite(): cant't register account,error = %v",err)
	}
	err = srv.Deposit(account.ID , types.Money(30_000))
	if err != nil{
		b.Fatalf("Favorite(): deposit,error = %v",err)
	}
	payment,er := srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		b.Fatalf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		b.Fatalf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		b.Fatalf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		b.Fatalf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		b.Fatalf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		b.Fatalf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		b.Fatalf("Favorite(): cant't pay,error = %v",err)
	}
	payment,er = srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		b.Fatalf("Favorite(): cant't pay,error = %v",err)
	}
	
	favorite,er :=srv.FavoritePayment(payment.ID,"test name")
	if er !=nil {
		b.Fatalf("Favorite(): cant't Favorite,error = %v",err)
	}
	_,err =srv.FindFavoriteByID(favorite.ID)
	if err !=nil {
		b.Fatalf("Favorite(): cant't Favorite,error = %v",err)
	} 
	_,err = srv.ExportAccountHistory(int64(1))
	if err !=nil {
		b.Fatalf("ExporrtToFile():,error = %v",err)
	}

	for i:=0;i<b.N;i++{
		result,_:= srv.FilterPayments(1,3)
		if len(result) != 8 {
			b.Fatalf("Incorrect result, want len %v, got len: %v",8,len(result))
		} 
	} 
	
}