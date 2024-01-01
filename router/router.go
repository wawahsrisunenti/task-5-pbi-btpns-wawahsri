package router

import (
	"PBI_BTPN/controllers"
	"PBI_BTPN/database"
	"PBI_BTPN/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userCtrl := controllers.UserController{DB: db}
	photoCtrl := controllers.PhotoController{DB: db}

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/register", userCtrl.RegisterUser)
		userRoutes.POST("/login", userCtrl.LoginUser)
		userRoutes.PUT("/:userId", middlewares.AuthMiddleware(), userCtrl.UpdateUser)
		userRoutes.DELETE("/:userId", middlewares.AuthMiddleware(), userCtrl.DeleteUser)
		userRoutes.POST("/logout", middlewares.AuthMiddleware(), userCtrl.LogoutUser)
	}

	photoRoutes := r.Group("/photos")
	photoRoutes.Use(middlewares.AuthMiddleware()) // Middleware applied to all photo routes
	{
		photoRoutes.POST("/", photoCtrl.CreatePhoto)
		photoRoutes.GET("/", photoCtrl.GetPhotos)
		photoRoutes.GET("/:photoId", photoCtrl.GetPhoto)
		photoRoutes.PUT("/:photoId", photoCtrl.UpdatePhoto)
		photoRoutes.DELETE("/:photoId", photoCtrl.DeletePhoto)
	}

	return r
}

func Init() *gin.Engine {
	db := database.InitDB()
	return SetupRouter(db)
}
