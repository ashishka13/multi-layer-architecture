package database

import "fmt"

type FcaDataLayerInterface interface {
	GetFcaAccount(id int) (account string, err error)
}

type FcaDataLayer struct{}

func NewFcaDataLayer() FcaDataLayerInterface {
	return FcaDataLayer{}
}

func (f FcaDataLayer) GetFcaAccount(id int) (account string, err error) {
	return "fca-account-" + fmt.Sprint(id), nil
}
