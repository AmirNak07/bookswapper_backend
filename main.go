package main

import (
	"bookswapper/api/routes"
	dbmodels "bookswapper/models/database"
	"bookswapper/utils/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db, dbErr := database.Connection()
	if dbErr != nil {
		panic("failed to connect database" + dbErr.Error())
	}

	// migrate all models
	migrateErr := db.AutoMigrate(&dbmodels.User{})
	if migrateErr != nil {
		panic("failed to migrate database" + migrateErr.Error())
	}

	// init fiber app
	app := fiber.New()

	// map api routes
	api := app.Group("/api")
	routes.AuthRouter(api, db)

	// map test route
	app.Get("/api/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// start server
	fiberError := app.Listen(":8080")
	if fiberError != nil {
		panic("cannot start server: " + fiberError.Error())
		return
	}
}
