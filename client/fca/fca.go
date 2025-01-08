package fca

import "multi-layer-architecture/database"

type FcaClientInterface interface {
	GetFcaAccount(id int) (account string, err error)
}

type FcaClient struct {
	FcaDatalyer database.FcaDataLayerInterface
}

func NewFcaClient() FcaClientInterface {
	return FcaClient{}
}

func (f FcaClient) GetFcaAccount(id int) (account string, err error) {
	account, err = f.FcaDatalyer.GetFcaAccount(id)
	return account, err
}
