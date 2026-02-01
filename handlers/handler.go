package handlers

import (
	"strconv"

	"multi-layer-architecture/services"

	"github.com/gofiber/fiber/v2"
)

func SetupHandler(s services.Services) *fiber.App {
	app := fiber.New()

	app.Get("/fca/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, _ := strconv.Atoi(idStr)
		acc, err := s.FcaService.GetFcaAllAccounts(id)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(fiber.Map{"account": acc})
	})

	app.Post("/pay", func(c *fiber.Ctx) error {
		type req struct {
			Amount      float64 `json:"amount"`
			Destination string  `json:"destination"`
		}
		r := new(req)
		if err := c.BodyParser(r); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		tx, err := s.PaymentService.ChargeAccount(r.Amount)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		_ = s.NotificationService.Send(r.Destination, "Payment processed: "+tx)
		return c.JSON(fiber.Map{"transaction": tx})
	})

	return app
}
