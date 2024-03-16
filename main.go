package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/pechpijit/Fiber_golang_example_api/configs"
	_ "github.com/pechpijit/Fiber_golang_example_api/docs" // load generated docs
	"github.com/pechpijit/Fiber_golang_example_api/middleware"
	"github.com/pechpijit/Fiber_golang_example_api/routes"
	"github.com/pechpijit/Fiber_golang_example_api/routes/api/v1"
	"github.com/pechpijit/Fiber_golang_example_api/service"
	"log"
	"os"
)

// @title Fiber Example API
// @description This is a sample swagger for Fiber
// @version 1.0
// @host localhost:3000
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app := fiber.New(configs.FiberConfig())
	api := app.Group("/api")
	v1 := api.Group("/v1")

	middleware.FiberMiddleware(app)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error load .env file", err)
	}

	service.AddMockUpData()

	routesApiV1.ProductRouter(v1)

	routes.LoginRouter(app)
	routes.SwaggerRoute(app)
	routes.NotFoundRoute(app)

	err := app.Listen(fmt.Sprintf(
		"%s:%s",
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"),
	))
	if err != nil {
		log.Fatal(err)
	}
}
