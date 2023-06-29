package routes

import (
	"go-http-server/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	//All routes related to users comes here
	app.Post("/user", controllers.CreateUser)
	app.Get("/user/:userId", controllers.GetAUser)
	app.Get("/users", controllers.GetAllUsers)
}
