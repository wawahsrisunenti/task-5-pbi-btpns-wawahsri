package app

import (
	"golang.org/x/crypto/bcrypt"
)

// CustomHashPassword meng-hash password yang diberikan menggunakan bcrypt.
func CustomHashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
