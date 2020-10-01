package wallet

import (
	"github.com/iamgafurov/wallet/pkg/types"
	"github.com/google/uuid"
	"errors"
	)


var ErrAccountNotFound = errors.New("account not found")
type Service struct {
	accounts []types.Account
	payments []types.Payment
}

func (s *Service) RegisterAccount(phone types.Phone){
	for _,account := range s.accounts{
		if account.Phone == phone{
			return
		}
	}
	paymentID:= uuid.New().String()
	s.accounts = append(s.accounts, types.Account{
		ID: paymentID,
		Phone: phone,
		Balance: 0,
	})
}

func (s *Service) FindAccoutByID(accountID int64)(*types.Account,error){
	for _,account := range s.accounts {
		if (account).ID == accountID {
			return &account,nil
		}
	}
	return nil, ErrAccountNotFound
}