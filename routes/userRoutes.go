package routes

import (
	"final-project/controllers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(route *gin.RouterGroup, db *gorm.DB) {
	userController := controllers.NewUserController(db)
	userGroup := route.Group("/users")
	{
		userGroup.POST("/register", userController.Register)
		userGroup.POST("/login", userController.Login)
		userGroup.Use(middlewares.Auth())
		userGroup.PUT("/", userController.Update)
		userGroup.DELETE("/", userController.Delete)
	}
}
