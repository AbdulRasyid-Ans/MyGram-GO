package middlewares

import (
	"final-project/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerAuthorization := ctx.Request.Header.Get("Authorization")
		if headerAuthorization == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "UNAUTHORIZED REQUEST",
			})
			return
		}

		bearer := strings.HasPrefix(headerAuthorization, "Bearer")
		if !bearer {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "UNAUTHORIZED REQUEST",
			})
			return
		}

		bearerToken := strings.Split(headerAuthorization, "Bearer ")[1]

		verify, err := helpers.ValidateToken(bearerToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
		}

		tokenData := verify.(jwt.MapClaims)

		ctx.Set("id", tokenData["id"])
		ctx.Set("email", tokenData["email"])
		ctx.Next()
	}
}
