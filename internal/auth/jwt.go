package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWTAuthenticator struct {
	PrivateKey string
}

func NewAuthenticator(privateKey string) *JWTAuthenticator {
	return &JWTAuthenticator{
		PrivateKey: privateKey,
	}
}

func (a *JWTAuthenticator) GenerateJWT(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(a.PrivateKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *JWTAuthenticator) VerifyJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrHashUnavailable
		}
		
		return []byte(a.PrivateKey), nil
	})
}
