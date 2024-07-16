package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zulanggara/TryGoFiber/server/handlers"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/languages", handlers.GetAllLanguagesData)
	app.Post("/languages", handlers.CreateNewLanguageData)
	app.Delete("/languages/:id", handlers.DeleteLanguageById)
	app.Put("/languages/:id", handlers.UpdateLanguageDataById)

}
