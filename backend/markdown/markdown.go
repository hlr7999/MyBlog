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
	e.GET("/markdown", GetMarkdown)
	e.Logger.Fatal(e.Start(addr))
}

func GetMarkdown(c echo.Context) (err error) {
	data, err := ioutil.ReadFile("G:/Projects/MyBlog/frontend/test.md")
    if err != nil {
        return err
	}
	dataStr := string(data)
	
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	return c.String(http.StatusOK, dataStr)
}
