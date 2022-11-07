package lib

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJwtToken(subject, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Subject:   subject,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secret))
}
