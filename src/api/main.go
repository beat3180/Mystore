package main

import (
	"api/database"
	"api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// GORMセット
	database.Connect()

	app := fiber.New()
	routes.Setup(app)

	app.Listen(":80")
}
