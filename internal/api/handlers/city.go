package handlers

import (
	dbmodels "bookswapper/internal/models/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetCities(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var cities []dbmodels.City
		result := db.Find(&cities)
		if result.Error != nil {
			errorString := fmt.Sprintf("failed to get all cities: %s", result.Error.Error())
			return c.Status(404).JSON(fiber.Map{
				"status": errorString,
			})
		}
		cityMap := make(map[string][]map[string]string)
		for _, city := range cities {
			cityEntry := map[string]string{
				"id":   fmt.Sprintf("%d", city.ID),
				"city": city.CityName,
			}
			cityMap["cities"] = append(cityMap["cities"], cityEntry)
		}

		return c.JSON(fiber.Map{
			"status": "success",
			"data":   cityMap,
		})
	}
}
