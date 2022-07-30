package controllers

import (
	"final-project/models"
	"final-project/params"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PhotoController struct {
	db *gorm.DB
}

func NewPhotoController(db *gorm.DB) *PhotoController {
	return &PhotoController{db: db}
}

func (p *PhotoController) GetList(ctx *gin.Context) {
	userId, isIdExist := ctx.Get("id")
	if !isIdExist {
		unauthorizeJsonResponse(ctx, "UNAUTHORIZED REQUEST")
		return
	}

	var user models.User
	err := p.db.First(&user, userId).Error
	if err != nil {
		unauthorizeJsonResponse(ctx, "UNAUTHORIZED REQUEST")
		return
	}

	var photos []models.Photo
	// err = p.db.Preload("Users").Find(&photos, "user_id=?", user.ID).Error
	err = p.db.Preload("User").Find(&photos).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	var response []params.ResponseGetPhoto
	for _, photo := range photos {
		var userPhoto params.UserPhoto
		if photo.User != nil {
			userPhoto = params.UserPhoto{
				Username: photo.User.Username,
				Email:    photo.User.Email,
			}
		}
		response = append(response, params.ResponseGetPhoto{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoUrl:  photo.PhotoUrl,
			UserID:    photo.UserID,
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
			User:      userPhoto,
		})
	}
	if len(response) == 0 {
		response = make([]params.ResponseGetPhoto, 0)
	}
	writeJsonResponse(ctx, http.StatusOK, response)
}

func (p *PhotoController) Create(ctx *gin.Context) {
	userId, isIdExist := ctx.Get("id")
	if !isIdExist {
		unauthorizeJsonResponse(ctx, "UNAUTHORIZED REQUEST")
		return
	}
	var user models.User
	err := p.db.First(&user, userId).Error
	if err != nil {
		unauthorizeJsonResponse(ctx, "UNAUTHORIZED REQUEST")
		return
	}

	var reqCreate params.RequestCreateOrUpdatePhoto
	err = ctx.ShouldBindJSON(&reqCreate)
	if err != nil {
		badRequestJsonResponse(ctx, err.Error())
		return
	}
	_, errValidator := govalidator.ValidateStruct(&reqCreate)
	if errValidator != nil {
		badRequestJsonResponse(ctx, errValidator.Error())
		return
	}

	newPhoto := models.Photo{
		Title:    reqCreate.Title,
		Caption:  reqCreate.Caption,
		PhotoUrl: reqCreate.PhotoUrl,
		UserID:   user.ID,
	}
	err = p.db.Create(&newPhoto).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	response := params.ResponseCreatePhoto{
		ID:        newPhoto.ID,
		Title:     newPhoto.Title,
		Caption:   newPhoto.Caption,
		PhotoUrl:  newPhoto.PhotoUrl,
		UserID:    newPhoto.UserID,
		CreatedAt: newPhoto.CreatedAt,
	}

	writeJsonResponse(ctx, http.StatusCreated, response)
}

func (p *PhotoController) Update(ctx *gin.Context) {
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))

	var reqUpdate params.RequestCreateOrUpdatePhoto
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

	var photo models.Photo
	err = p.db.First(&photo, uint(photoId)).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	updatePhoto := models.Photo{
		Title:    reqUpdate.Title,
		Caption:  reqUpdate.Caption,
		PhotoUrl: reqUpdate.PhotoUrl,
	}

	err = p.db.Model(&photo).Updates(updatePhoto).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	response := params.ResponseUpdatePhoto{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserID:    photo.UserID,
		UpdatedAt: photo.UpdatedAt,
	}

	writeJsonResponse(ctx, http.StatusOK, response)
}

func (p *PhotoController) Delete(ctx *gin.Context) {
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))

	var photo models.Photo
	err := p.db.First(&photo, uint(photoId)).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	err = p.db.Delete(&photo).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	response := params.ResponseMessage{
		Message: "Your photo has been successfully deleted",
	}

	writeJsonResponse(ctx, http.StatusOK, response)
}
