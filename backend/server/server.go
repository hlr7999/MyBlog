package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"MyBlog/app"
	"MyBlog/controller"
)

var e *echo.Echo

func init() {
    e = echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"*"},
    }))

	controller.InitUser(e)
	controller.InitAboutMe(e)
	controller.InitArticleNotAuth(e)

	g := e.Group("/api")
	g.Use(middleware.JWT([]byte(app.Secret())))

	controller.InitArticleAuth(g)

	app.InitLogger()

	app.Log.Info("Connecting to database")
	app.InitDB("127.0.0.1", "myblog")
}

func Start(bInitData bool, port string) {
	if bInitData {
		app.Log.Info("Init data")
		if app.InitData() != nil {
			app.Log.Error("init data error")
		}
		return
	}

	e.Logger.Fatal(e.Start(port))
}
