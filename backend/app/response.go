package app

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func Ok(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}

// func Created(c *echo.Context, entity interface{}) {
// 	c.JSON(http.StatusCreated, entity)
// }

// func Deleted(c *echo.Context) {
// 	c.JSON(http.StatusNoContent, nil)
// }

func ServerError(c echo.Context, err error) error {
	Log.Errorf("server error: %v", err)
	return c.JSON(http.StatusInternalServerError, map[string]string{
		"error": err.Error(),
	})
}

func BadRequest(c echo.Context, err string) error {
	Log.Errorf("server error: %s", err)
	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": err,
	})
}

// func Unauthorized(c *echo.Context) {
// 	c.JSON(http.StatusUnauthorized, echo.H{"error": "unauthorized"})
// }

func LoginFail(c echo.Context) error {
	return c.JSON(http.StatusNotFound, map[string]string{
		"error": "username or password error",
	})
}

func RegisterFail(c echo.Context, checkStr string, code int) error {
	return c.JSON(http.StatusConflict, map[string]string{
		"error":     checkStr + " has been registered",
		"errorCode": strconv.Itoa(code),
	})
}
