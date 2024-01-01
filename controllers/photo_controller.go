package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"PBI_BTPN/app"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PhotoController struct {
	DB *gorm.DB
}

func (pc *PhotoController) CreatePhoto(ctx *gin.Context) {
	var newPhoto app.Photo

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User dengan userID tersebut tidak ada"})
		return
	}

	userIDString, ok := userID.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal Mengubah userID ke string"})
		return
	}

	newPhoto.Title = ctx.PostForm("title")
	newPhoto.Caption = ctx.PostForm("caption")

	if newPhoto.Title == "" || newPhoto.Caption == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Title dan caption wajib diisi"})
		return
	}

	if _, err := govalidator.ValidateStruct(newPhoto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPhoto.UserID = app.UUIDString(userIDString)

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "File wajib dimasukan"})
		return
	}

	filePath := filepath.Join("uploads", string(newPhoto.ID)+filepath.Ext(file.Filename))

	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
		return
	}

	newPhoto.PhotoUrl = filePath

	if err := pc.DB.Create(&newPhoto).Error; err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal meyimpan file ke database"})
		return
	}

	ctx.JSON(http.StatusCreated, newPhoto)
}

func (pc *PhotoController) GetPhotos(ctx *gin.Context) {
	var photos []app.Photo

	pc.DB.Find(&photos)

	ctx.JSON(http.StatusOK, photos)
}

func (pc *PhotoController) GetPhoto(ctx *gin.Context) {
	photoID := ctx.Param("photoId")

	var photo app.Photo
	if err := pc.DB.Where("id = ?", photoID).First(&photo).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Photo tidak ditemukan"})
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

func (pc *PhotoController) UpdatePhoto(ctx *gin.Context) {
	photoID := ctx.Param("photoId")

	var existingPhoto app.Photo
	if err := pc.DB.Where("id = ?", photoID).First(&existingPhoto).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Photo tidak ditemukan"})
		return
	}

	var updatedPhoto app.Photo
	if err := ctx.ShouldBindJSON(&updatedPhoto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := govalidator.ValidateStruct(updatedPhoto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingPhoto.Title = updatedPhoto.Title
	existingPhoto.Caption = updatedPhoto.Caption

	pc.DB.Save(&existingPhoto)

	ctx.JSON(http.StatusOK, existingPhoto)
}

func (pc *PhotoController) DeletePhoto(ctx *gin.Context) {
	photoID := ctx.Param("photoId")

	var existingPhoto app.Photo
	if err := pc.DB.Where("id = ?", photoID).First(&existingPhoto).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Photo tidak ditemukan"})
		return
	}

	pc.DB.Delete(&existingPhoto)
	// ctx.JSON(http.StatusNoContent, nil)
	ctx.JSON(http.StatusNoContent, gin.H{"message": "User berhasil dihapus", "userId": photoID})
}
