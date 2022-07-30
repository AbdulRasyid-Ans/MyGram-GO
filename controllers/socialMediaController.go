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

type SocialMediaController struct {
	db *gorm.DB
}

func NewSocialMediaController(db *gorm.DB) *SocialMediaController {
	return &SocialMediaController{db: db}
}

func (s *SocialMediaController) GetList(ctx *gin.Context) {
	userId, isIdExist := ctx.Get("id")
	if !isIdExist {
		unauthorizeJsonResponse(ctx, "UNAUTHORIZED REQUEST")
		return
	}

	var user models.User
	err := s.db.First(&user, userId).Error
	if err != nil {
		unauthorizeJsonResponse(ctx, "UNAUTHORIZED REQUEST")
		return
	}

	var socialMedias []models.SocialMedia
	err = s.db.Preload("User").Find(&socialMedias).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	var response params.ResponseGetSocialMedia
	for _, socialMedia := range socialMedias {
		var userSocialMedia params.UserSocialMedia
		if socialMedia.User != nil {
			userSocialMedia = params.UserSocialMedia{
				ID:       socialMedia.User.ID,
				Username: socialMedia.User.Username,
			}
		}
		resSocialMedia := params.SocialMedia{
			ID:             socialMedia.ID,
			Name:           socialMedia.Name,
			SocialMediaUrl: socialMedia.SocialMediaUrl,
			UserID:         socialMedia.UserID,
			CreatedAt:      socialMedia.CreatedAt,
			UpdatedAt:      socialMedia.UpdatedAt,
			User:           userSocialMedia,
		}
		response.SocialMedias = append(response.SocialMedias, resSocialMedia)
	}
	if len(response.SocialMedias) == 0 {
		response.SocialMedias = make([]params.SocialMedia, 0)
	}
	writeJsonResponse(ctx, http.StatusOK, response)
}

func (s *SocialMediaController) Create(ctx *gin.Context) {
	userId, isIdExist := ctx.Get("id")
	if !isIdExist {
		unauthorizeJsonResponse(ctx, "UNAUTHORIZED REQUEST")
		return
	}
	var user models.User
	err := s.db.First(&user, userId).Error
	if err != nil {
		unauthorizeJsonResponse(ctx, "UNAUTHORIZED REQUEST")
		return
	}

	var reqCreate params.RequestCreateOrUpdateSocialMedia
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

	newSocialMedia := models.SocialMedia{
		Name:           reqCreate.Name,
		SocialMediaUrl: reqCreate.SocialMediaUrl,
		UserID:         user.ID,
	}
	err = s.db.Create(&newSocialMedia).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	response := params.ResponseCreateSocialMedia{
		ID:             newSocialMedia.ID,
		Name:           newSocialMedia.Name,
		SocialMediaUrl: newSocialMedia.SocialMediaUrl,
		UserID:         newSocialMedia.UserID,
		CreatedAt:      newSocialMedia.CreatedAt,
	}

	writeJsonResponse(ctx, http.StatusCreated, response)
}

func (s *SocialMediaController) Update(ctx *gin.Context) {
	socialMediaId, _ := strconv.Atoi(ctx.Param("socialMediaId"))

	var reqUpdate params.RequestCreateOrUpdateSocialMedia
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

	var socialMedia models.SocialMedia
	err = s.db.First(&socialMedia, uint(socialMediaId)).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	updateSocialMedia := models.SocialMedia{
		Name:           reqUpdate.Name,
		SocialMediaUrl: reqUpdate.SocialMediaUrl,
	}

	err = s.db.Model(&socialMedia).Updates(updateSocialMedia).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	response := params.ResponseUpdateSocialMedia{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserID:         socialMedia.UserID,
		UpdatedAt:      socialMedia.UpdatedAt,
	}

	writeJsonResponse(ctx, http.StatusOK, response)
}

func (s *SocialMediaController) Delete(ctx *gin.Context) {
	socialMediaId, _ := strconv.Atoi(ctx.Param("socialMediaId"))

	var socialMedia models.SocialMedia
	err := s.db.First(&socialMedia, uint(socialMediaId)).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	err = s.db.Delete(&socialMedia).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	response := params.ResponseMessage{
		Message: "Your social media has been successfully deleted",
	}

	writeJsonResponse(ctx, http.StatusOK, response)
}
