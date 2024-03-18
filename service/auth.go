package service

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pechpijit/Fiber_golang_example_api/models"
	"os"
	"strconv"
	"time"
)

var memberUser = models.LoginRequest{
	Email:    "super@admin.com",
	Password: "password",
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
// @Param json body models.LoginRequest true "Login info"
func Login(ctx *fiber.Ctx) error {
	loginInfo := new(models.LoginRequest)
	if err := ctx.BodyParser(loginInfo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	fmt.Printf("%s | %s", loginInfo.Email, loginInfo.Password)

	if loginInfo.Email != memberUser.Email || loginInfo.Password != memberUser.Password {
		return fiber.ErrUnauthorized
	}

	minutesCount, err := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))
	if err != nil || minutesCount <= 0 {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	claims := jwt.MapClaims{
		"email": memberUser.Email,
		"role":  "admin",
		"exp":   time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(fiber.Map{
		"message": "Login success.",
		"token":   t,
	})
}
