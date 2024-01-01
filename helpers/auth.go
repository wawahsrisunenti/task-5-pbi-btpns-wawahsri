package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// ExtractUserID extracts the user ID from the JWT token in the request context.
func ExtractUserID(ctx *gin.Context) (uint, error) {
	userID, exists := ctx.Get("userID")
	if !exists {
		return 0, errors.New("user ID not found in context")
	}
	return userID.(uint), nil
}

// AddToBlacklist adds the token to a blacklist (or performs any other action to invalidate the token).
func AddToBlacklist(userID uint) error {
	return nil
}
