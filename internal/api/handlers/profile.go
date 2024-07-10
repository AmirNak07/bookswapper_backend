package handlers

import (
	dbmodels "bookswapper/internal/models/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func GetMe(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtToken := c.Locals("user").(*jwt.Token)
		claims := jwtToken.Claims.(jwt.MapClaims)
		id := claims["id"].(string)
		println(id)

		var user dbmodels.User
		result := db.First(&user, "id = ?", id)
		if result.Error != nil {
			errorString := fmt.Sprintf("failed to find user: %s", result.Error.Error())
			return c.Status(404).JSON(fiber.Map{
				"status": errorString,
			})
		}

		return c.JSON(fiber.Map{
			"status": "success",
			"data": fiber.Map{
				"login":     user.Login,
				"username":  user.Username,
				"join_date": user.CreatedAt,
			},
		})
	}
}
