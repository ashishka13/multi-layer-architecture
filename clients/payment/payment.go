package payment

import "fmt"

type PaymentClientInterface interface {
	Charge(amount float64) (transactionID string, err error)
}

type PaymentClient struct {
}

func NewPaymentClient() PaymentClientInterface {
	return PaymentClient{}
}

func (p PaymentClient) Charge(amount float64) (transactionID string, err error) {
	return fmt.Sprintf("tx-%.0f", amount*100), nil
}
