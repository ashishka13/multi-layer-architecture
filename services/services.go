package services

import (
	"multi-layer-architecture/client/uwp"
	"multi-layer-architecture/services/fca"
)

type Services struct {
	FcaService fca.FcaServiceInterface
	UwpService uwp.UwpClientInterface
}

func InitServices(setting string) (services Services, err error) {

	uwpClient, err := uwp.NewUwpClientInterface("apply this setting")
	services = Services{
		FcaService: fca.NewFcaServiceInterface(),
		UwpService: uwpClient,
	}

	return services, err
}
