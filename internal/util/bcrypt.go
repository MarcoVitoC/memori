package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string ,error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
