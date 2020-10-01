package wallet

import "wallet/pkg/types"
import "errors"

var ErrAccountNotFound = errors.New("account not found")
type Service struct {
	accounts []types.Account
	payments []types.Payment
	nextAccountID int64
}

func (s *Service) RegisterAccount(phone types.Phone){
	for _,account := range s.accounts{
		if account.Phone == phone{
			return
		}
	}
	s.nextAccountID++
	s.accounts = append(s.accounts, types.Account{
		ID: s.nextAccountID,
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