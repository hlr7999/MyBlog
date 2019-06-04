package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"flag"
	"net/http"
	"io/ioutil"
)

func main() {
	var addr string
	flag.StringVar(&addr, "addr", ":10080", "server listens at this addr")
	flag.Parse()

	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/article", GetArticle)
	e.GET("/aboutMe", GetAboutMe)
	e.Logger.Fatal(e.Start(addr))
}

func GetArticle(c echo.Context) (err error) {
	data, err := ioutil.ReadFile("G:/Projects/MyBlog/backend/markdown/article.md")
    if err != nil {
        return err
	}
	dataStr := string(data)
	
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	return c.String(http.StatusOK, dataStr)
}

func GetAboutMe(c echo.Context) (err error) {
	data, err := ioutil.ReadFile("G:/Projects/MyBlog/backend/markdown/aboutMe.md")
    if err != nil {
        return err
	}
	dataStr := string(data)
	
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	return c.String(http.StatusOK, dataStr)
}
