package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
	"net/http"
	"MyBlog/db"
	"MyBlog/model"
)

const BaseURL = "/api/auth"

func Initialize(e *echo.Echo) error {
	e.POST(BaseURL+"/login", Login)
	e.POST(BaseURL+"/register", Register)
	return nil
}

type LoginForm struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

func Login(c echo.Context) (err error) {
	collection, closeConn := db.GlobalDB.User()
	defer closeConn()
	form := new(LoginForm)
	err = c.Bind(form)
	if err != nil {
		return err
	}

	user := new(data.User)
	err = collection.Find(bson.M{"username": form.Username}).One(user)
	if err == mgo.ErrNotFound {
		err2 := collection.Find(bson.M{"email": form.Username}).One(user)
		if err2 == mgo.ErrNotFound {
			return echo.ErrUnauthorized
		} else if err2 != nil {
			return err2
		}
	} else if err != nil {
		return err
	}

	ok := user.ComparePassword(form.Password)
	if !ok {
		return echo.ErrUnauthorized
	}
	claims := &model.Claims{
		UID:      user.Id,
		Username: user.Username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(data.Secret))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"token": t
	})
}

type RegisterForm struct {
	Username string `bson:"username" json:"username"`
	Email    string `bson:"email" json:"email`
	Password string `bson:"password" json:"password"`
}

func Register(c echo.Context) (err error) {
	collection, closeConn := db.GlobalDB.User()
	defer closeConn()
	form := new(RegisterForm)
	err = c.Bind(form)
	if err != nil {
		return err
	}

	user := new(data.User)
	err = collection.Find(bson.M{"username": req.Username}).One(user)
	if err != mgo.ErrNotFound {
		return c.String(http.StatusFound, "Username has been registered")
		err2 := collection.Find(bson.M{"email": req.Username}).One(user)
		if err2 != mgo.ErrNotFound {
			return c.String(http.StatusFound, "Email has been registered")
	} else if err != nil {
		return err
	}

	user.Username = form.Username
	user.Email = form.Email
	byts, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(byts)
	user.GenerateID()
	err = collection.Insert(user)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Register success!")
}
