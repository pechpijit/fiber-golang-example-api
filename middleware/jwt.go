package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pechpijit/Fiber_golang_example_api/response"
	"os"
	"time"
)

func JWTProtected() func(ctx *fiber.Ctx) error {
	config := jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET_KEY"))},
		ContextKey:   "jwt",
		ErrorHandler: jwtError,
	}

	return jwtware.New(config)
}

func JWTCheckPermission(ctx *fiber.Ctx, need *string) bool {
	user := ctx.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return claims[*need].(bool)
}

func JWTCheckExpire(ctx *fiber.Ctx) error {
	user := ctx.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	expiresAccessToken := int64(claims["expires"].(float64))

	now := time.Now().Unix()
	if now > expiresAccessToken {
		return response.RespondError(ctx, fiber.StatusUnauthorized, "unauthorized or expired token")
	}

	return ctx.Next()
}

func jwtError(ctx *fiber.Ctx, err error) error {
	// Return status 400 and failed bad request error.
	if err.Error() == "Missing or malformed JWT" {
		return response.RespondError(ctx, fiber.StatusBadRequest, err.Error())
	}

	// Return status 401 and failed authentication error.
	return response.RespondError(ctx, fiber.StatusUnauthorized, err.Error())
}
