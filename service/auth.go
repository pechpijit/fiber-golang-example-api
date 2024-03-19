package service

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/pechpijit/Fiber_golang_example_api/models"
	"github.com/pechpijit/Fiber_golang_example_api/response"
	"github.com/pechpijit/Fiber_golang_example_api/utils"
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
		return response.RespondError(ctx, fiber.StatusBadRequest, err.Error())

	}

	fmt.Printf("%s | %s", loginInfo.Email, loginInfo.Password)

	if loginInfo.Email != memberUser.Email || loginInfo.Password != memberUser.Password {
		return fiber.ErrUnauthorized
	}

	t, err := utils.GenerateNewTokens(memberUser.Email)
	if err != nil {
		return response.RespondError(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"message": "Login success.",
		"token":   t,
	})
}
