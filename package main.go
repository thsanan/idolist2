package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"respone": Actdress{
				ActID:   1,
				ActName: "Marai",
			},
		})
	})

	app.Listen(":3000")
}

type Actdress struct {
	ActID   int    `db:"act_id", json:"act_id"`
	ActName string `db:"act_name", json:"act_name"`
}
