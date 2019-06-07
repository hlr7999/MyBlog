package app

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"MyBlog/model"
)

var db *mgo.Database

func InitDB(url string, datebase string) {
	con, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	db = con.DB(datebase)

	// Ensure all index
	if err = model.EnsureUserIndex(db); err != nil {
		panic(err)
	}
}

func InitData() error {
	admin := &model.User{
		Username: "hlr7999",
		Password: "",
		Role:     model.AdminRole,
		Email:    "hlr7999@outlook.com",
	}
	admin.ID = bson.NewObjectId()
	admin.Avatar = "http://localhost/blog/img/avatar/admin.jpg"
	admin.SetPassword("heliren1999")

	err := db.C(model.UserC).Insert(admin)
	if err != nil {
		return err
	}

	return nil
}

func DB() *mgo.Database {
	return db
}
