package controller

import (
	"github.com/labstack/echo"
	"github.com/globalsign/mgo/bson"

	"MyBlog/model"
	"MyBlog/app"
)

func InitArticleNotAuth(e *echo.Echo) error {
	e.GET("/articles", getAllArticles)
	e.GET("/articles/:id", getArticle)
	return nil
}

func InitArticleAuth(g *echo.Group) error {
	g.POST("/articles", newArticle)
	g.PATCH("/articles/:id", updateArticle)
	g.DELETE("/articles/:id", deleteArticle)
	return nil
}

func getAllArticles(c echo.Context) error {
	collection := app.DB().C(model.ArticleC)

	var articles []model.Article
	err := collection.Find(nil).All(&articles)
	if err != nil {
		return app.ServerError(c, err)
	}

	return app.Ok(c, articles)
}

func getArticle(c echo.Context) error {
	collection := app.DB().C(model.ArticleC)

	id := c.Param("id")
	var article model.Article
	err := collection.FindId(bson.ObjectIdHex(id)).One(&article)
	if err != nil {
		return app.ServerError(c, err)
	}

	return app.Ok(c, article)
}

func newArticle(c echo.Context) error {
	// collection := app.DB().C(model.ArticleC)

	return nil
}

func updateArticle(c echo.Context) error {
	// collection := app.DB().C(model.ArticleC)

	return nil
}

func deleteArticle(c echo.Context) error {
	// collection := app.DB().C(model.ArticleC)

	return nil
}