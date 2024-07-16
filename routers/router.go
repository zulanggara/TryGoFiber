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

	app.Get("/user/list", handlers.GetAllUsers)
	app.Get("/user/find/:id", handlers.GetUserById)
	app.Post("/user/create", handlers.CreateNewUser)
	app.Put("/user/update/:id", handlers.UpdateUserDataById)
	app.Delete("/user/delete/:id", handlers.DeleteUserDataById)
}
