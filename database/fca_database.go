package database

type FcaDataLayerInterface interface {
	GetFcaAccount(id int) (account string, err error)
}

type FcaDatalyer struct {
}

func NewFcaDatalayerInterface() FcaDataLayerInterface {
	return FcaDatalyer{}
}

func (f FcaDatalyer) GetFcaAccount(id int) (account string, err error) {
	return "", nil
}
