package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zulanggara/TryGoFiber/configs"
	"github.com/zulanggara/TryGoFiber/routers"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Coders! Welcome to Go programming language.")
	})
	routers.SetupRoutes(app)
	configs.ConnectDB()

	app.Listen(":8080")
}
