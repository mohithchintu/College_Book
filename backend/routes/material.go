package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohithchintu/College_Book/backend/controllers"
)

func MaterialRoutes(app *fiber.App) {
	app.Post("/material", controllers.CreateMaterial)
	app.Get("/material", controllers.GetMaterials)
	app.Get("/material/:id", controllers.GetMaterialByID)

}
