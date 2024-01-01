package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// ExtractCustomUserID extracts the user ID from the JWT token in the request context.
func ExtractCustomUserID(ctx *gin.Context) (string, error) {
	userID, exists := ctx.Get("userID")
	if !exists {
		return "", errors.New("user ID not found in context")
	}
	return userID.(string), nil
}

// AddToCustomBlacklist adds the token to a blacklist (or performs any other action to invalidate the token).
func AddToCustomBlacklist(userID string) error {
	// Implement your logic to invalidate the token or perform any other necessary actions
	return nil
}
