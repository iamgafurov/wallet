package wallet
import (
	"github.com/iamgafurov/wallet/pkg/types"
	"testing"
	)

func TestService_Reject_success(t *testing.T){
	srv := &Service{}
	srv.RegisterAccount(types.Phone("90999390"))
	account,err := srv.FindAccountByID(1)
	if err != nil {
		t.Errorf("Reject(): cant't register account,error = %v",err)
	}
	err = srv.Deposit(account.ID , types.Money(30_000))
	if err != nil{
		t.Errorf("Reject(): deposit,error = %v",err)
	}
	payment,er := srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Reject(): cant't pay,error = %v",err)
	}
	err =srv.Reject(payment.ID)
	if err !=nil {
		t.Errorf("Reject(): cant't reject,error = %v",err)
	}
}
