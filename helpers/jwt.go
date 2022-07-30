package helpers

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

const jwtSecret = "hackNASAwithHTML"

func GenerateToken(id uint, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"email": email,
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))

	return tokenString, err
}

func ValidateToken(tokenString string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unauthorized token")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, fmt.Errorf("unauthorized token")
	}

	return token.Claims.(jwt.MapClaims), nil
}
