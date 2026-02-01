package main

import (
	"log"
	"multi-layer-architecture/handlers"
	"multi-layer-architecture/services"
)

func main() {
	allServices, err := services.InitServices("some setting")
	if err != nil {
		log.Println("err", err)
		return
	}

	acc, _ := allServices.FcaService.GetFcaAllAccounts(10)
	log.Print("acc", acc)

	app := handlers.SetupHandler(allServices)
	log.Fatal(app.Listen(":3000"))
}
