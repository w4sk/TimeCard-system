package main

import (
	"admin/src/database"
	"admin/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	//CORSの設定
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Listen(":3000")
}
