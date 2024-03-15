package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/pechpijit/Fiber_golang_example_api/response"
	"os"
)

func JWTProtected() func(*fiber.Ctx) error {
	config := jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ContextKey:   "jwt",
		ErrorHandler: jwtError,
	}

	return jwtware.New(config)
}

func jwtError(ctx *fiber.Ctx, err error) error {
	// Return status 400 and failed bad request error.
	if err.Error() == "Missing or malformed JWT" {
		return response.RespondError(ctx, fiber.StatusBadRequest, err.Error())
	}

	// Return status 401 and failed authentication error.
	return response.RespondError(ctx, fiber.StatusUnauthorized, err.Error())
}
