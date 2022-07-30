package controllers

import (
	"final-project/helpers"
	"final-project/models"
	"final-project/params"
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{db: db}
}

func (u *UserController) Register(ctx *gin.Context) {
	var reqRegister params.RequestRegisterUser

	err := ctx.ShouldBindJSON(&reqRegister)
	if err != nil {
		badRequestJsonResponse(ctx, err.Error())
		return
	}

	isValidAge := ageValidation(reqRegister.Age)
	_, errValidator := govalidator.ValidateStruct(&reqRegister)
	if !isValidAge {
		errMessage := "Age harus lebih dari 8"
		if errValidator == nil {
			errValidator = fmt.Errorf("%s", errMessage)
		} else {
			errValidator = fmt.Errorf("%w;Age harus lebih dari 8", errValidator)
		}
	}
	if errValidator != nil {
		badRequestJsonResponse(ctx, errValidator.Error())
		return
	}

	newUser := models.User{
		Age:      reqRegister.Age,
		Email:    reqRegister.Email,
		Password: reqRegister.Password,
		Username: reqRegister.Username,
	}

	err = u.db.Create(&newUser).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	response := params.ResponseRegisterUser{
		Age:      newUser.Age,
		Email:    newUser.Email,
		ID:       newUser.ID,
		Username: newUser.Username,
	}

	writeJsonResponse(ctx, http.StatusCreated, response)
}

func (u *UserController) Login(ctx *gin.Context) {
	var reqLogin params.RequestUserLogin

	err := ctx.ShouldBindJSON(&reqLogin)
	if err != nil {
		badRequestJsonResponse(ctx, err.Error())
		return
	}

	_, errValidator := govalidator.ValidateStruct(&reqLogin)
	if errValidator != nil {
		badRequestJsonResponse(ctx, errValidator.Error())
		return
	}

	var user models.User

	err = u.db.First(&user, "email=?", reqLogin.Email).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			unauthorizeJsonResponse(ctx, "email / password salah")
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	isValidPassword := helpers.VerifyPassword(user.Password, reqLogin.Password)

	if !isValidPassword {
		unauthorizeJsonResponse(ctx, "email / password salah")
		return
	}

	token, err := helpers.GenerateToken(user.ID, user.Email)
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	response := params.ResponseToken{
		Token: token,
	}

	writeJsonResponse(ctx, http.StatusOK, response)
}

func (u *UserController) Update(ctx *gin.Context) {
	var reqUpdate params.RequestUpdateUser

	userId, isIdExist := ctx.Get("id")
	if !isIdExist {
		unauthorizeJsonResponse(ctx, "UNAUTHORIZED REQUEST")
		return
	}

	var user models.User
	err := u.db.First(&user, userId).Error
	if err != nil {
		unauthorizeJsonResponse(ctx, "UNAUTHORIZED REQUEST")
		return
	}

	err = ctx.ShouldBindJSON(&reqUpdate)
	if err != nil {
		badRequestJsonResponse(ctx, err.Error())
		return
	}

	_, errValidator := govalidator.ValidateStruct(&reqUpdate)
	if errValidator != nil {
		badRequestJsonResponse(ctx, errValidator.Error())
		return
	}

	updateUser := models.User{
		Email:    reqUpdate.Email,
		Username: reqUpdate.Username,
	}

	err = u.db.Model(&user).Updates(updateUser).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	response := params.ResponseUpdateUser{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		Age:       user.Age,
		UpdatedAt: user.UpdatedAt,
	}

	writeJsonResponse(ctx, http.StatusOK, response)
}

func (u *UserController) Delete(ctx *gin.Context) {
	userId, isIdExist := ctx.Get("id")
	if !isIdExist {
		unauthorizeJsonResponse(ctx, "UNAUTHORIZED REQUEST")
		return
	}

	var user models.User
	err := u.db.First(&user, userId).Error
	if err != nil {
		unauthorizeJsonResponse(ctx, "UNAUTHORIZED REQUEST")
		return
	}

	err = u.db.Delete(&user).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	response := params.ResponseMessage{
		Message: "Your account has been successfully deleted",
	}

	writeJsonResponse(ctx, http.StatusOK, response)
}

func ageValidation(age int) bool {
	return age > 8
}
