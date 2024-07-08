package routes

import (
	"bookswapper/internal/api/handlers"
	"bookswapper/internal/api/middlewares"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProfileRouter(app fiber.Router, db *gorm.DB) {
	jwt := middlewares.AuthMiddleware("bookswapper")
	app.Get("/profiles/me", jwt, handlers.GetMe(db))
}
