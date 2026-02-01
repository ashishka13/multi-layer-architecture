package clients

import (
	cfca "multi-layer-architecture/clients/fca"
	cnotify "multi-layer-architecture/clients/notification"
	cpay "multi-layer-architecture/clients/payment"
	cuwp "multi-layer-architecture/clients/uwp"
	"multi-layer-architecture/database"
)

type Clients struct {
	Fca      cfca.FcaClientInterface
	Uwp      cuwp.UwpClientInterface
	Payment  cpay.PaymentClientInterface
	Notifier cnotify.Notifier
}

func NewClients(setting string) (Clients, error) {
	dl := database.NewFcaDataLayer()

	fcaClient := cfca.NewFcaClient(dl)
	uwpClient, err := cuwp.NewUwpClient(setting)
	if err != nil {
		return Clients{}, err
	}

	paymentClient := cpay.NewPaymentClient()
	notifier := cnotify.NewNotifier()

	return Clients{
		Fca:      fcaClient,
		Uwp:      uwpClient,
		Payment:  paymentClient,
		Notifier: notifier,
	}, nil
}
