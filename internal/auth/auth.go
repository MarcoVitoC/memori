package auth

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/golang-jwt/jwt/v5"
)

type Authenticator interface {
	GenerateJWT(claims jwt.MapClaims) (string, error)
}

func GenerateToken(length int) (string, error) {
	bytes := make([]byte, 32)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(bytes), nil
}