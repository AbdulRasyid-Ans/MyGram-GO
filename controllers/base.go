package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func badRequestJsonResponse(ctx *gin.Context, err interface{}) {
	writeJsonResponse(ctx, http.StatusBadRequest, gin.H{
		"error": err,
	})
}

func unauthorizeJsonResponse(ctx *gin.Context, err interface{}) {
	writeJsonResponse(ctx, http.StatusUnauthorized, gin.H{
		"error": err,
	})
}

func internalServerJsonResponse(ctx *gin.Context, err interface{}) {
	writeJsonResponse(ctx, http.StatusInternalServerError, gin.H{
		"error": err,
	})
}

func writeJsonResponse(ctx *gin.Context, status int, payload interface{}) {
	ctx.JSON(status, payload)
}
