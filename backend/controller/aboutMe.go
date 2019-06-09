package controller

import (
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"

	"MyBlog/app"
)

func InitAboutMe(e *echo.Echo) error {
	e.GET("/aboutMe", getAboutMe)
	return nil
}

func getAboutMe(c echo.Context) error {
	data, err := ioutil.ReadFile(app.FilePath + "data/aboutMe/aboutMe.md")
    if err != nil {
        return app.ServerError(c, err)
	}
	dataStr := string(data)
	
	return c.String(http.StatusOK, dataStr)
}
