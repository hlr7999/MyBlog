package controller

import (
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
	"net/http"
	"MyBlog/db"
	"MyBlog/model"
)

const BaseURL = "/api/users"

func Initialize(e *echo.Group) (err error) {
	b := crud.BasicCRUD{
		BaseURL:       BaseURL,
		GetCollection: db.Global.User,
	}
	b.SetRecordType(&model.User{})
	err = b.Check()
	if err != nil {
		return err
	}
	e.GET(BaseURL, b.All)
	e.POST(BaseURL, Create)
	e.GET(BaseURL+"/:id", b.Get)
	e.PUT(BaseURL+"/:id", b.Update)
	e.DELETE(BaseURL+"/:id", b.Delete)
	return nil
}

func Create(c echo.Context) (err error) {
	collection, closeConn := db.Global.User()
	defer closeConn()

	user := new(model.User)
	err = c.Bind(user)
	if err != nil {
		return err
	}
	n, err := collection.Find(bson.M{"username": user.Username}).Count()
	if err != nil {
		return err
	}
	if n != 0 {
		return c.NoContent(http.StatusConflict)
	}
	user.GenerateID()
	err = user.CryptPassword()
	if err != nil {
		return err
	}
	err = collection.Insert(user)
	if err != nil {
		return err
	}
	user.APIResponsePre()
	return c.JSON(http.StatusOK, user)
}
