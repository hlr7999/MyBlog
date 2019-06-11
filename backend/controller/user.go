package controller

import (
	"github.com/labstack/echo"
	"github.com/globalsign/mgo/bson"
	"net/http"
	"os"
	"io"
	"math/rand"
	"time"
	"strings"
	"strconv"

	"MyBlog/model"
	"MyBlog/app"
	"MyBlog/config"
)

func InitUserNotAuth(e *echo.Echo) {
	e.POST("/login", login)
	e.POST("/register", register)
}

func InitUserAuth(g *echo.Group) {
	g.GET("/users/:id", getUser)
	g.GET("/users", getAllUsers)
	g.POST("/uploadAvatar", uploadAvatar)
	g.PATCH("/users/:id", updateUser)
	g.DELETE("/users/:id", deleteUser)
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

	token, err := app.CreateToken(user.ID.Hex(), user.Username, user.Role)
	if err != nil {
		return app.ServerError(c, err)
	}

	return app.Ok(c, map[string]string {
		"token": token,
		"role": user.Role,
		"id": user.ID.Hex(),
		"avatar": user.Avatar,
	})
}

type RegisterRequest struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	Email    string `bson:"email" json:"email"`
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
			checkStr = registerReq.Email
		}
		n, err := collection.Find(bson.M{item: checkStr}).Count()
		if err != nil {
			return app.ServerError(c, err)
		}
		if n != 0 {
			return app.RegisterFail(c, item, i)
		}
	}
	
	user.Username = registerReq.Username
	user.Email = registerReq.Email

	user.Initialize()
	err = user.SetPassword(registerReq.Password)
	if err != nil {
		return app.ServerError(c, err)
	}

	err = collection.Insert(user)
	if err != nil {
		return app.ServerError(c, err)
	}

	userLCList := new(model.UserLCList)
	userLCList.ID = user.ID

	collection = app.DB().C(model.UserLCListC)
	err = collection.Insert(userLCList)
	if err != nil {
		return app.ServerError(c, err)
	}
	
	return app.Ok(c, user.ToPartial())
}

func getUser(c echo.Context) error {
	token := app.GetToken(c)
	collection := app.DB().C(model.UserC)

	id := c.Param("id")
	if id != token.ID {
		return app.BadRequest(c, "Bad Request")
	}

	var user model.User
	err := collection.FindId(bson.ObjectIdHex(id)).One(&user)
	if err != nil {
		return app.ServerError(c, err)
	}

	if user.Username != token.Username || user.Role != token.Role {
		return app.BadRequest(c, "Bad Request")
	}

	return app.Ok(c, user.ToPartial())
}

func getAllUsers(c echo.Context) error {
	token := app.GetToken(c)
	if token.Role != model.AdminRole {
		return app.BadRequest(c, "Bad Request")
	}

	collection := app.DB().C(model.UserC)

	var users []model.User
	err := collection.Find(nil).All(&users)
	if err != nil {
		return app.ServerError(c, err)
	}

	usersPartial := make([]model.UserPartial, len(users) - 1)
	i := 0
	for _, user := range users {
		if user.Role == model.AdminRole {
			continue
		}
		usersPartial[i] = user.ToPartial()
		i++
	}

	return app.Ok(c, users)
}

func uploadAvatar(c echo.Context) error {
	token := app.GetToken(c)
	collection := app.DB().C(model.UserC)
	rand.Seed(time.Now().UnixNano())

	// read file
	file, err := c.FormFile("file")
	if err != nil {
		return app.BadRequest(c, err.Error())
	}
	src, err := file.Open()
	if err != nil {
		return app.ServerError(c, err)
	}
	defer src.Close()

	// Destination
	// origin avatar
	var user model.User
	err = collection.FindId(bson.ObjectIdHex(token.ID)).One(&user)
	if err != nil {
		return app.ServerError(c, err)
	}
	oldAvatars := strings.Split(user.Avatar, "/")
	oldAvatar := oldAvatars[len(oldAvatars) - 1]
	if oldAvatar != "default.jpg" {
		err = os.Remove(config.ImagePath + "avatar/" + oldAvatar)
		if err != nil {
			return app.ServerError(c, err)
		}
	}
	// new avatar
	x := strconv.Itoa(rand.Intn(1000))
	dst, err := os.Create(config.ImagePath + "avatar/" + token.ID + x + ".jpg")
	if err != nil {
		return app.ServerError(c, err)
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return app.ServerError(c, err)
	}

	// database
	avatarPath := config.FrontImagePath + "avatar/" + token.ID + x + ".jpg"
	err = collection.Update(
		bson.M{"_id": bson.ObjectIdHex(token.ID)},
		bson.M{"$set": bson.M{"avatar": avatarPath}},
	)
	if err != nil {
		return app.ServerError(c, err)
	}

	return c.String(http.StatusOK, avatarPath)
}

type UpdateRequest struct {
	Username string `bson:"usernmae" json:"username"`
}

func updateUser(c echo.Context) error {
	token := app.GetToken(c)
	collection := app.DB().C(model.UserC)
	// get user id
	id := c.Param("id")
	if id != token.ID {
		return app.BadRequest(c, "Bad Request")
	}
	// get new user info
	updateReq := new(UpdateRequest)
	err := c.Bind(updateReq)
	if err != nil {
		return app.BadRequest(c, "Bad Request")
	}
	// new username exist
	n, err := collection.Find(bson.M{"username": updateReq.Username}).Count()
	if err != nil {
		return app.ServerError(c, err)
	}
	if n != 0 {
		return app.BadRequest(c, "1")
	}
	// update the user info
	err = collection.Update(
		bson.M{"_id": bson.ObjectIdHex(id)},
		bson.M{"$set": bson.M{"username": updateReq.Username}},
	)
	if err != nil {
		return app.ServerError(c, err)
	}

	return app.Ok(c, updateReq)
}

func deleteUser(c echo.Context) error {
	token := app.GetToken(c)
	if token.Role != model.AdminRole {
		return app.BadRequest(c, "Bad Request")
	}

	collection := app.DB().C(model.UserC)
	// get user id
	id := c.Param("id")

	// delete user
	// not admin
	if id == token.ID {
		return app.BadRequest(c, "Bad Request")
	}
	err := collection.Remove(
		bson.M{"_id": bson.ObjectIdHex(id)},
	)
	if err != nil {
		return app.ServerError(c, err)
	}

	return app.Ok(c, map[string]string{
		"message": "delete success",
	})
}