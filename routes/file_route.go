package routes

import (
	"go-http-server/controllers"

	"github.com/gofiber/fiber/v2"
)

func FileRoute(app *fiber.App) {
	//All routes related to file uploads come here
	app.Post("/file-upload", controllers.UploadHandler)
}
