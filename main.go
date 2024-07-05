package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// init fiber app
	app := fiber.New()

	// map test route
	app.Get("/api/ping", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// start server
	err := app.Listen(":8080")
	// handle error
	if err != nil {
		panic("cannot start server: " + err.Error())
		return
	}
}
