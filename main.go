package main

import (
	"final-project/database"
	"final-project/middlewares"
	"final-project/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	PORT    = ":4000"
	VERSION = "1.0"
)

func main() {
	router := gin.Default()

	db := database.ConnectDB()

	mainGroup := router.Group(fmt.Sprintf("/api/v%s", VERSION))
	{
		routes.UserRoutes(mainGroup, db)
		mainGroup.Use(middlewares.Auth())
		routes.PhotoRoutes(mainGroup, db)
		routes.CommentRoutes(mainGroup, db)
		routes.SocialMediaRoutes(mainGroup, db)
	}

	router.Run(PORT)
}
