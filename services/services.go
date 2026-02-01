package services

import (
	"multi-layer-architecture/clients"
	fcaSvc "multi-layer-architecture/services/fca"
	notificationSvc "multi-layer-architecture/services/notification"
	paymentSvc "multi-layer-architecture/services/payment"
)

type Services struct {
	Clients             clients.Clients
	FcaService          fcaSvc.FcaServiceInterface
	PaymentService      paymentSvc.PaymentServiceInterface
	NotificationService notificationSvc.NotificationServiceInterface
}

func InitServices(setting string) (services Services, err error) {
	cls, err := clients.NewClients("apply this setting" + setting)
	if err != nil {
		return services, err
	}

	fcaService := fcaSvc.NewFcaService(cls.Fca, cls.Uwp)

	paymentService := paymentSvc.NewPaymentService(cls.Payment)
	notifService := notificationSvc.NewNotificationService(cls.Notifier)

	services = Services{
		Clients:             cls,
		FcaService:          fcaService,
		PaymentService:      paymentService,
		NotificationService: notifService,
	}

	return services, nil
}
