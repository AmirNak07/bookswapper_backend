package main

import (
	"bookswapper/models"
	"bookswapper/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// get database configuration
	postgresHost := utils.GetEnv("POSTGRES_HOST", "localhost")
	postgresPort := utils.GetEnv("POSTGRES_PORT", "5432")
	postgresUser := utils.GetEnv("POSTGRES_USER", "postgres")
	postgresPassword := utils.GetEnv("POSTGRES_PASSWORD", "postgres")
	postgresDatabase := utils.GetEnv("POSTGRES_DB", "bookswapper")

	// create db url
	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		postgresHost, postgresUser, postgresPassword, postgresDatabase, postgresPort)

	// create database connection
	db, dbErr := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if dbErr != nil {
		panic("failed to connect database" + dbErr.Error())
	}

	// migrate all models
	migrateErr := db.AutoMigrate(&models.User{})
	if migrateErr != nil {
		panic("failed to migrate database" + migrateErr.Error())
	}

	// init fiber app
	app := fiber.New()

	// map test route
	app.Get("/api/ping", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/api/auth/register", func(c *fiber.Ctx) error {
		user := models.User{Id: 1, Login: "linuxfight", Password: "1234567890"}
		db.Create(&user)
		return c.SendString("user created")
	})

	app.Get("/api/auth/get", func(c *fiber.Ctx) error {
		var user models.User
		db.First(&user, "login = ?", "linuxfight")
		return c.JSON(&fiber.Map{
			"id":       user.Id,
			"login":    user.Login,
			"password": user.Password,
		})
	})

	// start server
	fiberError := app.Listen(":8080")
	if fiberError != nil {
		panic("cannot start server: " + fiberError.Error())
		return
	}
}
