package routes

import (
	"final-project/controllers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PhotoRoutes(route *gin.RouterGroup, db *gorm.DB) {
	photoController := controllers.NewPhotoController(db)
	photoGroup := route.Group("/photos")
	{
		photoGroup.GET("/", photoController.GetList)
		photoGroup.POST("/", photoController.Create)
		photoGroup.PUT("/:photoId", middlewares.PhotoAuthorization(db), photoController.Update)
		photoGroup.DELETE("/:photoId", middlewares.PhotoAuthorization(db), photoController.Delete)
	}
}
