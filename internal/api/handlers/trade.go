package handlers

import (
	dbmodels "bookswapper/internal/models/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
)

func GetTrades(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		offsetStr := c.Query("offset", "1")
		fmt.Println(offsetStr)
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			errorString := fmt.Sprintf("failed to get offset: %s", err.Error())
			return c.Status(404).JSON(fiber.Map{
				"status": errorString,
			})
		}
		var trades []dbmodels.Trade
		limit := 10

		result := db.Order("id desc").Offset((offset - 1) * limit).Limit(limit).Find(&trades)
		if result.Error != nil {
			errorString := fmt.Sprintf("failed to get 10 trades: %s", result.Error.Error())
			return c.Status(404).JSON(fiber.Map{
				"status": errorString,
			})
		}
		tradeMap := make(map[string][]map[string]string)
		for _, trade := range trades {
			var author dbmodels.User
			result_user := db.First(&author, "id = ?", trade.AuthorId)
			if result_user.Error != nil {
				errorString := fmt.Sprintf("failed to find author: %s", result_user.Error.Error())
				return c.Status(404).JSON(fiber.Map{
					"status": errorString,
				})
			}
			tradeEntry := map[string]string{
				"id":          fmt.Sprint(trade.ID),
				"bookname":    trade.BookName,
				"description": trade.Description,
				"Author_id":   fmt.Sprint(trade.AuthorId),
				"Author_name": author.Username,
			}
			tradeMap["cities"] = append(tradeMap["cities"], tradeEntry)
		}

		return c.JSON(fiber.Map{
			"status": "success",
			"data":   tradeMap,
		})
	}
}
