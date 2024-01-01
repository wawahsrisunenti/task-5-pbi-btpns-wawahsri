package app

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your-secret-key") // Ganti dengan kunci rahasia yang sesuai

// CustomClaims adalah struktur klaim JWT khusus.
type CustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

// GenerateCustomToken menghasilkan token JWT untuk ID pengguna yang diberikan.
func GenerateCustomToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token kedaluwarsa dalam 24 jam

	claims := &CustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseCustomToken mem-parsing token JWT yang diberikan dan mengembalikan klaim khusus.
func ParseCustomToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
