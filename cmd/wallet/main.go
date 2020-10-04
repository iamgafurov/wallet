package main

import (
		"github.com/iamgafurov/wallet/pkg/wallet"
		"github.com/iamgafurov/wallet/pkg/types"
	 	"fmt"
		)
func main(){
	srv := &wallet.Service{}
	srv.RegisterAccount(types.Phone("90999390"))
	account,err := srv.FindAccountByID(1)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(account)
	}
	err = srv.Deposit(account.ID , types.Money(30_000))
	if err != nil{
		fmt.Println(err)
	}
	payment,er := srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		fmt.Println(er)
	}
	payment,err =srv.Repeat(payment.ID)
	if err !=nil {
		fmt.Println(err)
	}else{
		fmt.Println(payment)
	}
}