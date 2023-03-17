package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Customer struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	CodeName string `json:"codeName"`
	Location string `json:"location"`
	Balance  string `json:"balance"`
}

func main() {

	app := fiber.New()

	app.Get("/healthCheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	customers := []Customer{}

	app.Post("/api/customer", func(c *fiber.Ctx) error {
		user := &Customer{}

		if err := c.BodyParser(user); err != nil {
			return err
		}
		user.ID = len(customers) + 1

		customers = append(customers, *user)
		return c.JSON(customers)
	})

	log.Fatal(app.Listen(":8080"))
}
