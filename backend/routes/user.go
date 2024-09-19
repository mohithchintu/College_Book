package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohithchintu/College_Book/backend/controllers"
)

func UserRoutes(app *fiber.App) {
	app.Post("/users", controllers.CreateUser)
	app.Put("/users/:id/profile", controllers.UpdateUserProfile)
	app.Put("/users/:id/preferences", controllers.UpdateUserPreferences)
	app.Delete("/users/:id", controllers.DeleteUserAccount)
}
