package main

import (
	"admin/src/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Connection To Mysql
	database.Connect()
	// Migration
	database.AutoMigrate()

	// fiber API
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")
}
