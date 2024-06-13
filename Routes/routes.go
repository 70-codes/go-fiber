package routes

import (
	Controller "rest_api/Controller"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	app.Post("/cashiers/:cashierId/login", Controller.Login)
	app.Get("/cashiers/:cashierId/logout", Controller.Logout)
	app.Post("/cashiers/:cashierId/passcode", Controller.Passcode)

	// Cashier endpoints
	app.Post("/cashiers", Controller.CreateCashier)
	app.Get("/cashiers", Controller.CashierList)
	app.Get("/cashiers/:cashierId", Controller.GetcashierDetails)
	app.Delete("/cashiers/:cashierId", Controller.DeleteCashier)
	app.Put("/cashiers/:cashierId", Controller.UpdateCashier)

	// Category endpoints
	app.Post("/categories", Controller.CreateCategory)
	app.Get("/categories", Controller.CategoryList)
	app.Get("/categories/:categoryId", Controller.GetCategoryDetails)
	app.Delete("/categories/:categoryId", Controller.DeleteCategory)
	app.Put("/categories/:categoryId", Controller.UpdateCategory)
}
