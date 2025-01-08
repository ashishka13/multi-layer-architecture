package fca

import (
	"fmt"
	"multi-layer-architecture/client/fca"
	"multi-layer-architecture/client/uwp"
)

type FcaServiceInterface interface {
	GetFcaAllAccounts(id int) (account string, err error)
}

type FcaService struct {
	FcaClient fca.FcaClientInterface
	UwpClient uwp.UwpClientInterface
}

func NewFcaServiceInterface() FcaServiceInterface {
	return FcaService{}
}

func (f FcaService) GetFcaAllAccounts(id int) (account string, err error) {
	uwpaccount, err := f.UwpClient.GetUwpAccounts(id)
	fmt.Println("uwpaccounts", uwpaccount, err, err)

	account, err = f.FcaClient.GetFcaAccount(id)
	return account, err
}
