package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

var secretKey = "welcome to my blog xixi haha"

func Secret() string {
	return secretKey
}

func CreateToken(id string, username string, role string) (string, error) {
	claims := jwt.MapClaims{
		"id":       id,
		"username": username,
		"role":     role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	result, err := token.SignedString([]byte(secretKey))
	return result, err
}

type ParseToken struct {
	ID       string
	Username string
	Role     string
}

func GetToken(c echo.Context) ParseToken {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	parseToken := ParseToken {
		ID: claims["id"].(string),
		Username: claims["username"].(string),
		Role: claims["role"].(string),
	}
	
	return parseToken
}
