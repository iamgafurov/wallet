package wallet

import (
	"github.com/iamgafurov/wallet/pkg/types"
	"errors"
	)


var ErrAccountNotFound = errors.New("account not found")
var ErrPhoneExist = errors.New("account by this phone exist")

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