package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pechpijit/Fiber_golang_example_api/middleware"
	"github.com/pechpijit/Fiber_golang_example_api/service"
)

func InitProductRouter(app *fiber.App) {
	productGroup := app.Group("api/v1/product")

	productGroup.Get("/", service.GetProducts)
	productGroup.Get("/:id", service.GetProduct)

	productGroup.Post("/", middleware.JWTProtected(), service.CreateProduct)
	productGroup.Put("/:id", middleware.JWTProtected(), service.UpdateProduct)
	productGroup.Delete("/:id", middleware.JWTProtected(), service.DeleteProduct)
}
