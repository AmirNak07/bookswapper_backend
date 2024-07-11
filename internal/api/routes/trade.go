package routes

import (
	"bookswapper/internal/api/handlers"
	"bookswapper/internal/api/middlewares"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func TradeRouter(app fiber.Router, db *gorm.DB) {
	jwt := middlewares.AuthMiddleware("bookswapper")
	app.Get("/trades", handlers.GetTrades(db))

	app.Post("/trade", jwt, handlers.AddTrade(db))
}
