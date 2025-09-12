package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(mail string, adm bool) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("server misconfigured: missing JWT_SECRET")
	}

	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"mail": mail,
		"adm":  adm,
		"iat":  now.Unix(),
		"nbf":  now.Unix(),
		"exp":  now.Add(24 * time.Hour).Unix(),
		"iss":  "go-mongo-api",
	})
	return token.SignedString([]byte(secret))
}
