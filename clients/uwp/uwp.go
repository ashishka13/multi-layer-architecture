package uwp

import "fmt"

type UwpClientInterface interface {
	GetUwpAccounts(id int) (account string, err error)
}

type UwpClient struct {
	Something string
}

func NewUwpClient(someSetting string) (UwpClientInterface, error) {
	fmt.Println("some setting applied if required", someSetting)
	return UwpClient{Something: someSetting}, nil
}

func (u UwpClient) GetUwpAccounts(id int) (account string, err error) {
	return fmt.Sprint(id, " uwp account"), nil
}
