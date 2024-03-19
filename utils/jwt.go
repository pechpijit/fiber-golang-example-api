package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/pechpijit/Fiber_golang_example_api/repository"
	"os"
	"strconv"
	"time"
)

func GenerateNewTokens(email string) (string, error) {
	minute, err := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))
	if err != nil || minute <= 0 {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["expires"] = time.Now().Add(time.Minute * time.Duration(minute)).Unix()
	claims[repository.ProductCreateCredential] = true
	claims[repository.ProductUpdateCredential] = true
	claims[repository.ProductDeleteCredential] = true

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	return t, nil
}
