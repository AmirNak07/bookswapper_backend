package routes

import (
	"bookswapper/internal/api/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func TradeRouter(app fiber.Router, db *gorm.DB) {
	app.Get("/trades", handlers.GetTrades(db))
}
