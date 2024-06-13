package main

import (
	"fmt"
	routes "rest_api/Routes"
	db "rest_api/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	fmt.Println("Hello World")
	db.Connect()

	app := fiber.New()
	app.Use(cors.New())

	routes.SetUp(app)
	app.Listen(":3000")

}
