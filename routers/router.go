package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zulanggara/TryGoFiber/server/handlers"
)

func SetupRoutes(app *fiber.App) {

	languageRoute := app.Group("/language")
	languageRoute.Get("/", handlers.GetAllLanguagesData)
	languageRoute.Post("/", handlers.CreateNewLanguageData)
	languageRoute.Delete("/:id", handlers.DeleteLanguageById)
	languageRoute.Put("/:id", handlers.UpdateLanguageDataById)

	userRoute := app.Group("/user")
	userRoute.Get("/list", handlers.GetAllUsers)
	userRoute.Get("/find/:id", handlers.GetUserById)
	userRoute.Post("/create", handlers.CreateNewUser)
	userRoute.Put("/update/:id", handlers.UpdateUserDataById)
	userRoute.Delete("/delete/:id", handlers.DeleteUserDataById)
}
