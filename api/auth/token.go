package auth

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id uint32) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"id":         id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}
