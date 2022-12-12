package main

import (
	"admin/src/database"
	"admin/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Connection To Mysql
	database.Connect()
	// Migration
	database.AutoMigrate()

	config := fiber.Config{
		ReadBufferSize: 8192,
	}
	// fiber API
	app := fiber.New(config)

	routes.Setup(app)

	app.Listen(":3000")
}
