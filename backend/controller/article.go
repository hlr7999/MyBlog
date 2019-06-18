package controller

import (
	"MyBlog/config"
	"github.com/labstack/echo"
	"github.com/globalsign/mgo/bson"
	"time"
	"os"
	"io"
	"net/http"
	"io/ioutil"

	"MyBlog/model"
	"MyBlog/app"
)

func InitArticleNotAuth(e *echo.Echo) error {
	e.GET("/articles", getAllArticles)
	e.GET("/articles/:id", getArticle)
	e.GET("/articles/class/:id", getArticlesByClass)
	return nil
}

func InitArticleAuth(g *echo.Group) error {
	g.POST("/articles", newArticle)
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

func getArticlesByClass(c echo.Context) error {
	collection := app.DB().C(model.ArticleC)
	cc := app.DB().C(model.ClassC)
	// get class
	id := c.Param("id")
	var class model.Class
	err := cc.FindId(bson.ObjectIdHex(id)).One(&class)
	if err != nil {
		return app.BadRequest(c, err.Error())
	}
	// get articles
	if class.Level == model.FirstLevel {
		query := make([]bson.M, len(class.Child))
		for i, item := range class.Child {
			query[i] = bson.M{"classId": item}
		}
		var articles []model.Article
		if (len(query) > 0) {
			err = collection.Find(bson.M{"$or": query}).All(&articles)
			if err != nil {
				return app.BadRequest(c, err.Error())
			}
		}

		return app.Ok(c, articles)
	} else {
		var articles []model.Article
		err = collection.Find(bson.M{"classId": id}).All(&articles)
		if err != nil {
			return app.BadRequest(c, err.Error())
		}

		return app.Ok(c, articles)
	}
}

func getArticle(c echo.Context) error {
	collection := app.DB().C(model.ArticleC)

	id := c.Param("id")
	var article model.Article
	err := collection.FindId(bson.ObjectIdHex(id)).One(&article)
	if err != nil {
		return app.ServerError(c, err)
	}

	article.BrowseCount += 1
	err = collection.Update(
		bson.M{"_id": article.ID},
		bson.M{"$set": bson.M{"browseCount": article.BrowseCount}},
	)
	if err != nil {
		return app.ServerError(c, err)
	}

	data, err := ioutil.ReadFile(article.Content)
    if err != nil {
        return app.ServerError(c, err)
	}
	article.Content = string(data)

	return app.Ok(c, article)
}

func newArticle(c echo.Context) error {
	token := app.GetToken(c)
	if token.Role != model.AdminRole {
		return app.BadRequest(c, "Bad Request")
	}
	collection := app.DB().C(model.ArticleC)

	//===== basic info
	article := new(model.Article)
	article.ID = bson.NewObjectId()
	article.Title = c.FormValue("title")
	article.Description = c.FormValue("description")
	article.ClassId = c.FormValue("classId")
	article.ClassName = c.FormValue("className")
	now := time.Now()
	article.Year = now.Year()
	article.Month = int(now.Month())
	article.Day = now.Day()

	//===== read content file
	cfile, err := c.FormFile("content")
	if err != nil {
		return app.BadRequest(c, err.Error())
	}
	csrc, err := cfile.Open()
	if err != nil {
		return app.ServerError(c, err)
	}
	defer csrc.Close()
	// contenr destination
	cdstpath := config.FilePath + "articles/" + article.ID.Hex() + ".md"
	cdst, err := os.Create(cdstpath)
	if err != nil {
		return app.ServerError(c, err)
	}
	defer cdst.Close()
	// copy content file
	if _, err = io.Copy(cdst, csrc); err != nil {
		return app.ServerError(c, err)
	}
	article.Content = cdstpath

	//===== read cover file
	ifile, err := c.FormFile("imageCover")
	if err != nil {
		err = collection.Insert(article)
		if err != nil {
			return app.ServerError(c, err)
		}

		return c.String(http.StatusOK, "OK")
	}
	isrc, err := ifile.Open()
	if err != nil {
		return app.ServerError(c, err)
	}
	defer isrc.Close()
	// cover destination
	idst, err := os.Create(config.ImagePath + "articleCover/" + article.ID.Hex() + ".jpg")
	if err != nil {
		return app.ServerError(c, err)
	}
	defer idst.Close()
	// cover content file
	if _, err = io.Copy(idst, isrc); err != nil {
		return app.ServerError(c, err)
	}
	article.ImageCover = config.FrontImagePath + "articleCover/" + article.ID.Hex() + ".jpg"

	//===== database
	err = collection.Insert(article)
	if err != nil {
		return app.ServerError(c, err)
	}

	return c.String(http.StatusOK, "OK")
}

func deleteArticle(c echo.Context) error {
	collection := app.DB().C(model.ArticleC)
	token := app.GetToken(c)
	if token.Role != model.AdminRole {
		return app.BadRequest(c, "Bad Request")
	}

	// get article id
	id := c.Param("id")

	// delete article content and article
	var article model.Article
	err := collection.FindId(bson.ObjectIdHex(id)).One(&article)
	if err != nil {
		return app.ServerError(c, err)
	}
	err = os.Remove(article.Content)
	if err != nil {
		return app.ServerError(c, err)
	}
	err = collection.Remove(
		bson.M{"_id": bson.ObjectIdHex(id)},
	)
	if err != nil {
		return app.ServerError(c, err)
	}

	return c.String(http.StatusOK, "OK")
}
