package controllers

import (
	"final-project/models"
	"final-project/params"
	"fmt"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentController struct {
	db *gorm.DB
}

func NewCommentController(db *gorm.DB) *CommentController {
	return &CommentController{db: db}
}

func (c *CommentController) GetList(ctx *gin.Context) {
	userId, isIdExist := ctx.Get("id")
	if !isIdExist {
		unauthorizeJsonResponse(ctx, "UNAUTHORIZED REQUEST")
		return
	}

	var user models.User
	err := c.db.First(&user, userId).Error
	if err != nil {
		unauthorizeJsonResponse(ctx, "UNAUTHORIZED REQUEST")
		return
	}

	var comments []models.Comment
	err = c.db.Preload("User").Preload("Photo").Find(&comments).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	var response []params.ResponseGetComment
	for _, comment := range comments {
		var userComment params.UserComment
		var photoComment params.PhotoComment

		if comment.User != nil {
			userComment = params.UserComment{
				ID:       comment.User.ID,
				Username: comment.User.Username,
				Email:    comment.User.Email,
			}
		}
		if comment.Photo != nil {
			photoComment = params.PhotoComment{
				ID:       comment.Photo.ID,
				Title:    comment.Photo.Title,
				Caption:  comment.Photo.Caption,
				PhotoUrl: comment.Photo.PhotoUrl,
				UserID:   comment.Photo.UserID,
			}
		}
		response = append(response, params.ResponseGetComment{
			ID:        comment.ID,
			Message:   comment.Message,
			PhotoID:   comment.PhotoID,
			UserID:    comment.UserID,
			UpdatedAt: comment.UpdatedAt,
			CreatedAt: comment.CreatedAt,
			User:      userComment,
			Photo:     photoComment,
		})
	}

	if len(response) == 0 {
		response = make([]params.ResponseGetComment, 0)
	}

	writeJsonResponse(ctx, http.StatusOK, response)
}

func (c *CommentController) Create(ctx *gin.Context) {
	userId, isIdExist := ctx.Get("id")
	if !isIdExist {
		unauthorizeJsonResponse(ctx, "UNAUTHORIZED REQUEST")
		return
	}
	var user models.User
	err := c.db.First(&user, userId).Error
	if err != nil {
		unauthorizeJsonResponse(ctx, "UNAUTHORIZED REQUEST")
		return
	}

	var reqCreate params.RequestCreateComment
	err = ctx.ShouldBindJSON(&reqCreate)
	if err != nil {
		badRequestJsonResponse(ctx, err.Error())
		return
	}
	_, errValidator := govalidator.ValidateStruct(&reqCreate)
	var photo models.Photo
	errPhotoId := c.db.Select("id").First(&photo, reqCreate.PhotoID).Error
	if errPhotoId != nil {
		errMessage := "Photo ID tidak valid"
		if errValidator == nil {
			errValidator = fmt.Errorf("%s", errMessage)
		} else {
			errValidator = fmt.Errorf("%w;%s", errValidator, errMessage)
		}
	}
	if errValidator != nil {
		badRequestJsonResponse(ctx, errValidator.Error())
		return
	}

	newComment := models.Comment{
		UserID:  user.ID,
		PhotoID: reqCreate.PhotoID,
		Message: reqCreate.Message,
	}
	err = c.db.Create(&newComment).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	response := params.ResponseCreateComment{
		ID:        newComment.ID,
		Message:   newComment.Message,
		PhotoID:   newComment.PhotoID,
		UserID:    newComment.UserID,
		CreatedAt: newComment.CreatedAt,
	}

	writeJsonResponse(ctx, http.StatusCreated, response)
}

func (c *CommentController) Update(ctx *gin.Context) {
	commentId, _ := strconv.Atoi(ctx.Param("commentId"))

	var reqUpdate params.RequestUpdateComment
	err := ctx.ShouldBindJSON(&reqUpdate)
	if err != nil {
		badRequestJsonResponse(ctx, err.Error())
		return
	}
	_, errValidator := govalidator.ValidateStruct(&reqUpdate)
	if errValidator != nil {
		badRequestJsonResponse(ctx, errValidator.Error())
		return
	}

	var comment models.Comment
	err = c.db.First(&comment, uint(commentId)).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	updateComment := models.Comment{
		Message: reqUpdate.Message,
	}

	err = c.db.Model(&comment).Updates(updateComment).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	response := params.ResponseUpdateComment{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		UpdatedAt: comment.UpdatedAt,
	}

	writeJsonResponse(ctx, http.StatusOK, response)
}

func (c *CommentController) Delete(ctx *gin.Context) {
	commentId, _ := strconv.Atoi(ctx.Param("commentId"))

	var comment models.Comment
	err := c.db.First(&comment, uint(commentId)).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	err = c.db.Delete(&comment).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	response := params.ResponseMessage{
		Message: "Your comment has been successfully deleted",
	}

	writeJsonResponse(ctx, http.StatusOK, response)
}
