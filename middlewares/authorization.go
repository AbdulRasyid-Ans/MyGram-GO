package middlewares

import (
	"final-project/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PhotoAuthorization(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := getUserId(db, ctx)
		if userId == nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "UNAUTHORIZED REQUEST",
			})
			return
		} else {
			userId = uint(userId.(float64))
		}

		photoId, err := strconv.Atoi(ctx.Param("photoId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Parameter",
			})
			return
		}

		var photo models.Photo
		err = db.Select("user_id").First(&photo, uint(photoId)).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Photo data doesn't exist",
			})
			return
		}

		if photo.UserID != userId {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "UNAUTHORIZED REQUEST",
				"message": "you are not allowed to access this data",
			})
		}
		ctx.Next()
	}
}

func CommentAuthorization(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := getUserId(db, ctx)
		if userId == nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "UNAUTHORIZED REQUEST",
			})
			return
		} else {
			userId = uint(userId.(float64))
		}

		commentId, err := strconv.Atoi(ctx.Param("commentId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Parameter",
			})
			return
		}

		var comment models.Comment
		err = db.Select("user_id").First(&comment, uint(commentId)).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Comment data doesn't exist",
			})
			return
		}

		if comment.UserID != userId {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "UNAUTHORIZED REQUEST",
				"message": "you are not allowed to access this data",
			})
		}
		ctx.Next()
	}
}

func SocialMediaAuthorization(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := getUserId(db, ctx)
		if userId == nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "UNAUTHORIZED REQUEST",
			})
			return
		} else {
			userId = uint(userId.(float64))
		}

		socialMediaId, err := strconv.Atoi(ctx.Param("socialMediaId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Parameter",
			})
			return
		}

		var socialMedia models.SocialMedia
		err = db.Select("user_id").First(&socialMedia, uint(socialMediaId)).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Social Media data doesn't exist",
			})
			return
		}

		if socialMedia.UserID != userId {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "UNAUTHORIZED REQUEST",
				"message": "you are not allowed to access this data",
			})
		}
		ctx.Next()
	}
}

func getUserId(db *gorm.DB, ctx *gin.Context) interface{} {
	userId, isIdExist := ctx.Get("id")
	if !isIdExist {
		return nil
	}

	var user models.User
	err := db.Select("id").First(&user, userId).Error
	if err != nil {
		return nil
	}
	return userId
}
