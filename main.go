package main

import (
	database2 "bookswapper/models/database"
	"bookswapper/models/requests"
	"bookswapper/utils/database"
	"bookswapper/utils/password"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func main() {
	db, dbErr := database.DatabaseConnection()
	if dbErr != nil {
		panic("failed to connect database" + dbErr.Error())
	}

	// migrate all models
	migrateErr := db.AutoMigrate(&database2.User{})
	if migrateErr != nil {
		panic("failed to migrate database" + migrateErr.Error())
	}

	// init fiber app
	app := fiber.New()

	// map test route
	app.Get("/api/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Post("/api/auth/register", func(c *fiber.Ctx) error {
		data := &requests.User{}
		if reqErr := c.BodyParser(data); reqErr != nil {
			errorString := fmt.Sprintf("invalid json: %s", reqErr.Error())
			return c.Status(400).JSON(fiber.Map{
				"status": errorString,
			})
		}

		hashedPassword, hashErr := password.HashPassword(data.Password)
		if hashErr != nil {
			errorString := fmt.Sprintf("failed to hash password: %s", hashErr.Error())
			return c.Status(400).JSON(fiber.Map{
				"status": errorString,
			})
		}
		user := &database2.User{
			Login:        data.Login,
			PasswordHash: hashedPassword,
			CreatedAt:    time.Now(),
		}
		result := db.Create(&user)
		if result.Error != nil {
			errorString := fmt.Sprintf("failed to create user: %s", result.Error.Error())
			return c.Status(400).JSON(fiber.Map{
				"status": errorString,
			})
		}
		return c.JSON(fiber.Map{
			"status": "user created",
		})
	})

	app.Post("/api/auth/login", func(c *fiber.Ctx) error {
		data := &requests.User{}
		if err := c.BodyParser(data); err != nil {
			errorString := fmt.Sprintf("invalid json: %s", err.Error())
			return c.Status(400).JSON(fiber.Map{
				"status": errorString,
			})
		}
		var user database2.User
		result := db.First(&user, "login = ?", data.Login)
		if result.Error != nil {
			errorString := fmt.Sprintf("failed to find user: %s", result.Error.Error())
			return c.Status(404).JSON(fiber.Map{
				"status": errorString,
			})
		}
		if password.CheckPasswordHash(data.Password, user.PasswordHash) {
			return c.Status(400).JSON(fiber.Map{
				"status": "wrong password",
			})
		}
		return c.JSON(&fiber.Map{
			"status": "user logged in",
		})
	})

	// start server
	fiberError := app.Listen(":8080")
	if fiberError != nil {
		panic("cannot start server: " + fiberError.Error())
		return
	}
}
