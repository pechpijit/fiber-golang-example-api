package routesApiV1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pechpijit/Fiber_golang_example_api/middleware"
	"github.com/pechpijit/Fiber_golang_example_api/service"
)

func ProductRouter(app fiber.Router) {
	productGroup := app.Group("product")

	productGroup.Get("/", service.GetProducts)
	productGroup.Get("/:id", service.GetProduct)
	productGroup.Post("/", middleware.JWTProtected(), middleware.JWTCheckExpire, service.CreateProduct)
	productGroup.Put("/:id", middleware.JWTProtected(), middleware.JWTCheckExpire, service.UpdateProduct)
	productGroup.Delete("/:id", middleware.JWTProtected(), middleware.JWTCheckExpire, service.DeleteProduct)
}
