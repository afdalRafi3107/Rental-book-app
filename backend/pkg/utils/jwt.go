package utils

import (
	"rental-book/internal/config"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(id uint) (string, error) {
	config.EnvLoad()
	claims := jwt.MapClaims{
		"user_id":id,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("ikanamasnd"))
}