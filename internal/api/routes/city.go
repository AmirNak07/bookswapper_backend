package routes

import (
	"bookswapper/internal/api/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CityRouter(app fiber.Router, db *gorm.DB) {
	app.Get("/utils/allcities", handlers.GetCities(db))
}
