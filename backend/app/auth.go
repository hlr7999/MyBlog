package app

import (
	// "errors"
	// "net/http"
	// "strings"

	"github.com/dgrijalva/jwt-go"
	// "github.com/labstack/echo"
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

// func RequireAuth() echo.HandlerFunc {
// 	return func(c *echo.Context) {
// 		c.Header("WWW-Authenticate", "JWT realm="+auth.Realm)

// 		if authenticateUser(c) == false {
// 			c.Abort()
// 			return
// 		}
// 	}
// }

// func RequireRole(allowedRoles ...string) echo.HandlerFunc {
// 	return func(c *echo.Context) {
// 		c.Header("WWW-Authenticate", "JWT realm="+auth.Realm)

// 		if authenticateUser(c) == false {
// 			Unauthorized(c)
// 			return
// 		}

// 		role, exists := c.Get("role")
// 		if !exists {
// 			Unauthorized(c)
// 			return
// 		}

// 		for _, allowedRole := range allowedRoles {
// 			if role == allowedRole {
// 				return
// 			}
// 		}

// 		Unauthorized(c)
// 	}
// }

// func authenticateUser(c *echo.Context) bool {
// 	token, err := parseHeader(c.Request.Header.Get("Authorization"))
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 		return false
// 	}

// 	claims := token.Claims.(jwt.MapClaims)
// 	user := claims["user"].(string)
// 	role := claims["role"].(string)

// 	if user == "" || role == "" {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user or role"})
// 		return false
// 	}

// 	c.Set("user", user)
// 	c.Set("role", role)
// 	return true
// }

// func parseHeader(header string) (*jwt.Token, error) {
// 	if header == "" {
// 		return nil, errors.New("empty authorization header")
// 	}

// 	parts := strings.SplitN(header, " ", 2)
// 	if len(parts) != 2 || parts[0] != "Bearer" {
// 		return nil, errors.New("invalid authorization header")
// 	}

// 	return jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
// 		if token.Method != jwt.SigningMethodHS256 {
// 			return nil, errors.New("invalid singing algorithm")
// 		}

// 		return auth.Secret, nil
// 	})
// }
