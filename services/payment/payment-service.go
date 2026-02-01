package payment

import "multi-layer-architecture/clients/payment"

type PaymentServiceInterface interface {
	ChargeAccount(amount float64) (string, error)
}

type PaymentService struct {
	client payment.PaymentClientInterface
}

func NewPaymentService(c payment.PaymentClientInterface) PaymentServiceInterface {
	return PaymentService{client: c}
}

func (p PaymentService) ChargeAccount(amount float64) (string, error) {
	return p.client.Charge(amount)
}
