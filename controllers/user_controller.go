package controllers

import (
	"net/http"

	"PBI_BTPN/app"
	"PBI_BTPN/helpers"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CustomUserController struct {
	DB *gorm.DB
}

func (uc *CustomUserController) RegisterCustomUser(ctx *gin.Context) {
	var newUser app.CustomUser
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := govalidator.ValidateStruct(newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser app.CustomUser
	if err := uc.DB.Where("email = ?", newUser.Email).First(&existingUser).Error; err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Email sudah terdaftar"})
		return
	}

	hashedPassword, err := app.CustomHashPassword(newUser.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	newUser.Password = hashedPassword

	uc.DB.Create(&newUser)

	token, err := helpers.GenerateCustomToken(newUser.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal generate token"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"token": token, "userID": newUser.ID})
}

func (uc *CustomUserController) LoginCustomUser(ctx *gin.Context) {
	var loginUser app.CustomUser
	if err := ctx.ShouldBindJSON(&loginUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := govalidator.ValidateStruct(loginUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser app.CustomUser
	if err := uc.DB.Where("email = ?", loginUser.Email).First(&existingUser).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Email Salah"})
		return
	}

	if !helpers.CheckPasswordHash(loginUser.Password, existingUser.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Password Salah"})
		return
	}

	token, err := helpers.GenerateCustomToken(existingUser.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token, "userID": existingUser.ID})
}

func (uc *CustomUserController) UpdateCustomUser(ctx *gin.Context) {
	userID := ctx.Param("userId")

	var existingUser app.CustomUser
	if err := uc.DB.Where("id = ?", userID).First(&existingUser).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		return
	}

	var updatedUser app.CustomUser
	if err := ctx.ShouldBindJSON(&updatedUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := govalidator.ValidateStruct(updatedUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := app.CustomHashPassword(updatedUser.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal hashing password"})
		return
	}

	existingUser.Username = updatedUser.Username
	existingUser.Email = updatedUser.Email
	existingUser.Password = hashedPassword

	uc.DB.Save(&existingUser)

	ctx.JSON(http.StatusOK, existingUser)
}

func (uc *CustomUserController) DeleteCustomUser(ctx *gin.Context) {
	userID := ctx.Param("userId")

	var existingUser app.CustomUser
	if err := uc.DB.Where("id = ?", userID).First(&existingUser).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		return
	}

	uc.DB.Delete(&existingUser)

	ctx.JSON(http.StatusOK, gin.H{"message": "User berhasil dihapus", "userId": userID})
}

func (uc *CustomUserController) LogoutCustomUser(ctx *gin.Context) {
	// Extract user ID from the token
	userID, err := helpers.ExtractUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	if err := helpers.AddToBlacklist(userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
