package routes

import (
	"final-project/controllers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CommentRoutes(route *gin.RouterGroup, db *gorm.DB) {
	commentController := controllers.NewCommentController(db)
	commentGroup := route.Group("/comments")
	{
		commentGroup.GET("/", commentController.GetList)
		commentGroup.POST("/", commentController.Create)
		commentGroup.PUT("/:commentId", middlewares.CommentAuthorization(db), commentController.Update)
		commentGroup.DELETE("/:commentId", middlewares.CommentAuthorization(db), commentController.Delete)
	}
}
