package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pechpijit/Fiber_golang_example_api/service"
)

func LoginRouter(app *fiber.App) {
	app.Post("/login", service.Login)
}
