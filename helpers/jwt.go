package helpers

import (
	"time"

	"PBI_BTPN/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// CustomGenerateToken generates a JWT token for the given user ID.
func CustomGenerateToken(userID uuid.UUID) (string, error) {
	secretKey := []byte(CustomGetSecretKey())

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// CustomGetSecretKey retrieves the secret key for JWT signing.
func CustomGetSecretKey() string {
	secretKey := config.GetSecretKey()
	if secretKey == "" {
		secretKey = "defaultsecret"
	}
	return secretKey
}
