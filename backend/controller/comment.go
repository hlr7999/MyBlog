package controller

import (
	"github.com/labstack/echo"
	"github.com/globalsign/mgo/bson"
	"time"
	"strings"
	"net/http"

	"MyBlog/model"
	"MyBlog/app"
)

func InitCommentNotAuth(e *echo.Echo) error {
	e.GET("/comments/:aid", getComments)
	return nil
}

func InitCommentAuth(g *echo.Group) error {
	g.POST("/comments/:aid", newComment)
	g.DELETE("/comments", deleteComment)
	return nil
}

func getComments(c echo.Context) error {
	collection := app.DB().C(model.CommentC)

	aid := c.Param("aid")
	var comments []model.Comment
	err := collection.Find(bson.M{"aid": aid}).All(&comments)
	if err != nil {
		return app.BadRequest(c, err.Error())
	}

	return app.Ok(c, comments)
}

type Content struct {
	Content string `json:"content"`
}

func newComment(c echo.Context) error {
	collection := app.DB().C(model.CommentC)
	uc := app.DB().C(model.UserC)

	// get user
	token := app.GetToken(c)
	var user model.User
	err := uc.FindId(bson.ObjectIdHex(token.ID)).One(&user)
	if err != nil {
		return app.BadRequest(c, err.Error())
	}

	// new comments
	comment := new(model.Comment)
	comment.ID = bson.NewObjectId()
	comment.AID = c.Param("aid")
	comment.UID = token.ID
	comment.Username = user.Username
	comment.Avatar = user.Avatar
	comment.Time = strings.Split(time.Now().String(), " ")[0]

	var content Content
	err = c.Bind(&content)
	if err != nil {
		return app.BadRequest(c, err.Error())
	}
	comment.Content = content.Content

	// database
	err = collection.Insert(comment)
	if err != nil {
		return app.ServerError(c, err)
	}

	return c.String(http.StatusOK, "OK")
}

type CommentId struct {
	ID string `json:"id"`
}

func deleteComment(c echo.Context) error {
	collection := app.DB().C(model.CommentC)
	token := app.GetToken(c)
	
	// get comment
	var commentId CommentId
	err := c.Bind(&commentId)
	if err != nil {
		return app.BadRequest(c, err.Error())
	}
	id := commentId.ID
	var comment model.Comment
	err = collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&comment)
	if err != nil {
		return app.BadRequest(c, err.Error())
	}
	if comment.UID != token.ID {
		return app.BadRequest(c, "Bad Request")
	}

	// delete comment
	err = collection.Remove(
		bson.M{"_id": bson.ObjectIdHex(id)},
	)
	if err != nil {
		return app.ServerError(c, err)
	}

	return c.String(http.StatusOK, "OK")
}
