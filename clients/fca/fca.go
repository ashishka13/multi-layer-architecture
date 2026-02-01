package fca

import "multi-layer-architecture/database"

type FcaClientInterface interface {
	GetFcaAccount(id int) (account string, err error)
}

type FcaClient struct {
	FcaDataLayer database.FcaDataLayerInterface
}

func NewFcaClient(d database.FcaDataLayerInterface) FcaClientInterface {
	return FcaClient{FcaDataLayer: d}
}

func (f FcaClient) GetFcaAccount(id int) (account string, err error) {
	account, err = f.FcaDataLayer.GetFcaAccount(id)
	return account, err
}
