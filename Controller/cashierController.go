package controller

import (
	"fmt"
	models "rest_api/Models"
	db "rest_api/config"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateCashier(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Invalid data",
			})
	}
	if data["name"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier Name is required",
			})
	}
	if data["passcode"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Passcode is required",
			})
	}
	cashier := models.Cashier{
		Name:      data["name"],
		Passcode:  data["passcode"],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db.DB.Create(&cashier)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Cashier Created",
		"data":    cashier,
	})
}

func CashierList(c *fiber.Ctx) error {
	var cashier []models.Cashier
	limit, err := strconv.Atoi(c.Query("limit", "10")) // Default limit to 10 if not provided
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false, "message": "Invalid limit value"})
	}

	skip, err := strconv.Atoi(c.Query("skip", "0")) // Default skip to 0 if not provided
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false, "message": "Invalid skip value"})
	}

	var count int64
	result := db.DB.Limit(limit).Offset(skip).Find(&cashier).Count(&count)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"success": false, "message": "Database query failed", "error": result.Error.Error()})
	}
	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "cashier list api",
			"count":   count,
			"data":    cashier,
		})
}
func GetcashierDetails(c *fiber.Ctx) error {

	cashierId := c.Params("cashierId")
	var cashier models.Cashier
	db.DB.Select("id, name").Where("id=?", cashierId).First(&cashier)
	cashierData := make(map[string]interface{})
	cashierData["cashierID"] = cashier.Id
	cashierData["name"] = cashier.Name
	cashierData["createdAt"] = cashier.CreatedAt
	cashierData["updatedAt"] = cashier.UpdatedAt
	fmt.Println(cashierData)

	if cashier.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Not Found",
			"error":   map[string]interface{}{},
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    cashierData,
	})
}
func UpdateCashier(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier models.Cashier
	db.DB.Find(&cashier, "id = ?", cashierId)
	if cashier.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Not Found",
		})
	}
	var updateCashier models.Cashier
	err := c.BodyParser(&updateCashier)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data",
		})
	}
	if updateCashier.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Name is required",
		})
	}
	cashier.Name = updateCashier.Name
	db.DB.Save(&cashier)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    cashier,
	})
}

func DeleteCashier(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier models.Cashier
	db.DB.Where("id = ?", cashierId).First(&cashier)
	if cashier.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Not Found",
		})
	}
	db.DB.Where("id = ?", cashierId).Delete(&cashier)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Cashier deleted successfully",
	})
}
