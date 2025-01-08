package uwp

import "fmt"

type UwpClientInterface interface {
	GetUwpAccounts(id int) (account string, err error)
}

type UwpClient struct {
	Something string
}

func NewUwpClientInterface(someSetting string) (uwpClientInterface UwpClientInterface, err error) {
	fmt.Println("some setting applied if required", someSetting)
	return UwpClient{}, nil
}

func (u UwpClient) GetUwpAccounts(id int) (account string, err error) {
	return fmt.Sprint(id, " uwp account"), nil
}
