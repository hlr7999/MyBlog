package model

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"golang.org/x/crypto/bcrypt"
)

const UserC = "users"

const (
	AdminRole = "admin"
	UserRole  = "user"
	GuestRole = "guest"
)

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Username string        `bson:"username,omitempty" json:"username,omitempty"`
	Password string        `bson:"password" json:"password"`
	Email    string        `bson:"emial,omitempty" json:"email,omitempty"`
	Role     string        `bson:"role"  json:"role"`
}

func (u *User) ToPartial() UserPartial {
	return UserPartial{
		ID:       u.ID,
		Username: u.Username,
	}
}

type UserPartial struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Username string        `bson:"username,omitempty" json:"username,omitempty"`
	Role     string        `bson:"role"  json:"role"`
}

func EnsureUserIndex(db *mgo.Database) error {
	nameIndex := mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		Background: true,
	}

	emailIndex := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		Background: true,
	}

	err := db.C(UserC).EnsureIndex(nameIndex)
	if err != nil {
		return err
	}

	err = db.C(UserC).EnsureIndex(emailIndex)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) SetPassword(password string) (err error) {
	byts, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(byts)
	return nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) Initialize() {
	if u.ID == "" {
		u.ID = bson.NewObjectId()
	}
	u.Role = UserRole
}

func (u *User) CryptPassword() (err error) {
	if u.Password != "" {
		byts, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(byts)
	}
	return nil
}
