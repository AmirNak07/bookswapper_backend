package handlers

import (
	dbmodels "bookswapper/internal/models/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"strconv"
)

func GetMe(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtToken := c.Locals("user").(*jwt.Token)
		claims := jwtToken.Claims.(jwt.MapClaims)
		id := claims["id"].(string)

		var user dbmodels.User
		result_user := db.First(&user, "id = ?", id)
		if result_user.Error != nil {
			errorString := fmt.Sprintf("failed to find user: %s", result_user.Error.Error())
			return c.Status(404).JSON(fiber.Map{
				"status": errorString,
			})
		}
		var city dbmodels.City
		result_city := db.First(&city, "id = ?", user.CityId)
		if result_city.Error != nil {
			errorString := fmt.Sprintf("failed to find city: %s", result_city.Error.Error())
			return c.Status(404).JSON(fiber.Map{
				"status": errorString,
			})
		}
		return c.JSON(fiber.Map{
			"status": "success",
			"data": fiber.Map{
				"login":        user.Login,
				"username":     user.Username,
				"biography":    user.Biography,
				"phone_number": user.PhoneNumber,
				"city":         city.CityName,
				"join_date":    user.CreatedAt,
			},
		})
	}
}

func GetUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		UserIdStr := c.Query("userid", "1")
		UserId, err := strconv.Atoi(UserIdStr)
		if err != nil {
			errorString := fmt.Sprintf("failed to get offset: %s", err.Error())
			return c.Status(404).JSON(fiber.Map{
				"status": errorString,
			})
		}

		var user dbmodels.User
		result_user := db.First(&user, "id = ?", UserId)
		if result_user.Error != nil {
			errorString := fmt.Sprintf("failed to find user: %s", result_user.Error.Error())
			return c.Status(404).JSON(fiber.Map{
				"status": errorString,
			})
		}
		var city dbmodels.City
		result_city := db.First(&city, "id = ?", user.CityId)
		if result_city.Error != nil {
			errorString := fmt.Sprintf("failed to find city: %s", result_city.Error.Error())
			return c.Status(404).JSON(fiber.Map{
				"status": errorString,
			})
		}
		return c.JSON(fiber.Map{
			"status": "success",
			"data": fiber.Map{
				"login":        user.Login,
				"username":     user.Username,
				"biography":    user.Biography,
				"phone_number": user.PhoneNumber,
				"city":         city.CityName,
				"join_date":    user.CreatedAt,
			},
		})
	}
}
