package fca

import (
	"fmt"
	cfca "multi-layer-architecture/clients/fca"
	cuwp "multi-layer-architecture/clients/uwp"
)

type FcaServiceInterface interface {
	GetFcaAllAccounts(id int) (account string, err error)
	GetUWPAllAccounts(id int) (account string, err error)
}

type FcaService struct {
	FcaClient cfca.FcaClientInterface
	UwpClient cuwp.UwpClientInterface
}

func NewFcaService(fcaClient cfca.FcaClientInterface, uwpClient cuwp.UwpClientInterface) FcaServiceInterface {
	return FcaService{FcaClient: fcaClient, UwpClient: uwpClient}
}

func (f FcaService) GetFcaAllAccounts(id int) (account string, err error) {
	account, err = f.FcaClient.GetFcaAccount(id)
	fmt.Println("fcaaccount", account, err)
	return account, err
}

func (f FcaService) GetUWPAllAccounts(id int) (account string, err error) {
	uwpaccount, err := f.UwpClient.GetUwpAccounts(id)
	fmt.Println("uwpaccount", uwpaccount, err)
	return uwpaccount, err
}
