package data

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
	"github.com/golang/crypto/bcrypt"
)

const (
	Secret      = "WelcomeToHBlog;hlr12138@qq.com"
	ContextUser = "CtxUser"
)

// User

type User struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Username string        `bson:"username" json:"username"`
	Password string        `bson:"password" json:"password"`
	Email    string        `bson:"email,omitempty" json:"email,omitempty"`
}

func (u *User) GenerateID() {
	u.Id = bson.NewObjectId()
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

//----User----

// Claims

type Claims struct {
	jwt.StandardClaims
	UID      bson.ObjectId `bson:"uid" json:"uid"`
	Username string        `bson:"username" json:"username"`
}

//----Claims----
