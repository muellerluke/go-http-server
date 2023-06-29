package main

import (
	"go-http-server/configs"
	"go-http-server/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 1024 * 1024 * 10,
	})

	//run database
	configs.ConnectDB()

	//middlewares
	app.Use(logger.New())
	app.Get("/metrics", monitor.New())

	//routes
	routes.UserRoute(app)
	routes.FileRoute(app)

	app.Listen(":3001")
}
