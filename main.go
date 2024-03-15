package main

import (
	"fmt"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/swagger"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/pechpijit/Fiber_golang_example_api/docs" // load generated docs
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var memberUser = User{
	Email:    "super@admin.com",
	Password: "password",
}

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
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error load .env file")
	}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	addMockUpData()

	app.Post("/login", login)
	app.Get("/", renderTemplate)

	//app.Use(middleWare)

	productGroup := app.Group("/products")
	productGroup.Get("/", getProducts)
	productGroup.Get("/:id", getProduct)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
	}))
	productGroup.Use(middleWare)

	productGroup.Post("/", createProduct)
	productGroup.Put("/:id", updateProduct)
	productGroup.Delete("/:id", deleteProduct)

	app.Post("/uploadFile", uploadFile)

	app.Get("/getEnv/:name", getEnv)

	app.Listen(":3000")
}

// Handler functions
// login godoc
// @Summary Login
// @Description Login with email and password
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Success 200
// @Router /login [post]
// @Param json body User true "Login info"
func login(ctx *fiber.Ctx) error {
	loginInfo := new(User)
	if err := ctx.BodyParser(loginInfo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	fmt.Printf("%s | %s", loginInfo.Email, loginInfo.Password)

	if loginInfo.Email != memberUser.Email || loginInfo.Password != memberUser.Password {
		return fiber.ErrUnauthorized
	}

	claims := jwt.MapClaims{
		"email": memberUser.Email,
		"role":  "admin",
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(fiber.Map{
		"message": "Login success.",
		"token":   t,
	})
}

func middleWare(ctx *fiber.Ctx) error {
	start := time.Now()
	fmt.Printf("URL: %s, Method: %s, Time: %s\n", ctx.OriginalURL(), ctx.Method(), start)

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role := claims["role"].(string)

	if role != "admin" {
		return fiber.ErrUnauthorized
	}

	return ctx.Next()
}

func getEnv(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	//if val, exit := os.LookupEnv(name); exit {
	//	return ctx.JSON(fiber.Map{
	//		name: val,
	//	})
	//}

	//return ctx.SendStatus(fiber.StatusNotFound)

	return ctx.JSON(fiber.Map{
		name: os.Getenv("handsome"),
	})
}

func renderTemplate(c *fiber.Ctx) error {
	return c.Render("template", fiber.Map{
		"Name": "PechyEiEi",
	})
}

func uploadFile(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = ctx.SaveFile(file, "./upload/"+file.Filename)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.SendString("File upload complete!")
}

func addMockUpData() {
	products = append(products, Product{
		ID:       1,
		Name:     "cc_item_health",
		Price:    500,
		Discount: 10,
	})

	products = append(products, Product{
		ID:       2,
		Name:     "cc_target_farm",
		Price:    900,
		Discount: 15,
	})
}
