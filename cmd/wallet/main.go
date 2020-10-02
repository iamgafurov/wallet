package main

import (
		"github.com/iamgafurov/wallet/pkg/wallet"
		"github.com/iamgafurov/wallet/pkg/types"
	 	"fmt"
		)
func main(){
	srv := &wallet.Service{}
	srv.RegisterAccount(types.Phone("90999390"))
	account,err := srv.FindAccoutByID(1)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(*account)
	}
}