package controller

import (
	"github.com/labstack/echo"
	"github.com/globalsign/mgo/bson"

	"MyBlog/model"
	"MyBlog/app"
)

func InitUser(e *echo.Echo) (err error) {
	e.POST("/login", login)
	e.POST("/register", register)
	return nil
}

type LoginRequest struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

func login(c echo.Context) error {
	collection := app.DB().C(model.UserC)

	loginReq := new(LoginRequest)
	err := c.Bind(loginReq)
	if err != nil {
		return app.BadRequest(c, "Bad Request")
	}

	user := new(model.User)
	err = collection.Find(bson.M{"username": loginReq.Username}).One(user)
	if err != nil {
		return app.LoginFail(c)
	}
	
	if !user.ComparePassword(loginReq.Password) {
		return app.LoginFail(c)
	}

	token, err := app.CreateToken(user.ID.String(), user.Username, user.Role)
	if err != nil {
		return app.ServerError(c, err)
	}

	return app.Ok(c, map[string]string {
		"token": token,
		"role": user.Role,
		"id": user.ID.String(),
		"avatar": user.Avatar,
	})
}

type RegisterRequest struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	Email    string `bson:"emial" json:"email"`
}

func register(c echo.Context) error {
	collection := app.DB().C(model.UserC)

	registerReq := new(RegisterRequest)
	err := c.Bind(registerReq)
	if err != nil {
		return app.BadRequest(c, "Bad Request")
	}

	user := new(model.User)

	checkArr := [2]string {"username","email"}
	for i, item := range checkArr {
		var checkStr string
		if i == 0 {
			checkStr = registerReq.Username
		} else {
			checkStr = registerReq.Password
		}
		n, err := collection.Find(bson.M{item: checkStr}).Count()
		if err != nil {
			return app.ServerError(c, err)
		}
		if n != 0 {
			return app.RegisterFail(c, item)
		}
	}
	
	user.Initialize()
	err = user.CryptPassword()
	if err != nil {
		return app.ServerError(c, err)
	}

	err = collection.Insert(user)
	if err != nil {
		return app.ServerError(c, err)
	}
	
	return app.Ok(c, user.ToPartial())
}