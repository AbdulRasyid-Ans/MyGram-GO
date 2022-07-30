package routes

import (
	"final-project/controllers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SocialMediaRoutes(route *gin.RouterGroup, db *gorm.DB) {
	socialMediaController := controllers.NewSocialMediaController(db)
	socialMediaGroup := route.Group("/socialmedias")
	{
		socialMediaGroup.GET("/", socialMediaController.GetList)
		socialMediaGroup.POST("/", socialMediaController.Create)
		socialMediaGroup.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(db), socialMediaController.Update)
		socialMediaGroup.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(db), socialMediaController.Delete)
	}
}
