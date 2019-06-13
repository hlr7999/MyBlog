package model

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const ClassC = "classes"

const (
	FirstLevel = 1
	SecondLevel  = 2
)

type Class struct {
	ID    bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Name  string        `bson:"name,omitempty" json:"name,omitempty"`
	Level int           `bson:"level" json:"level"`
	Child []string      `bson:"child" json:"child"`
}

type ClassPartial struct {
	ID    string `bson:"_id,omitempty" json:"_id,omitempty"`
	Name  string `bson:"name,omitempty" json:"name,omitempty"`
}

func (c *Class) ToPartial() ClassPartial {
	return ClassPartial{
		ID:   c.ID.Hex(),
		Name: c.Name,
	}
}

func EnsureClassIndex(db *mgo.Database) error {
	nameIndex := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		Background: true,
	}

	err := db.C(ClassC).EnsureIndex(nameIndex)
	if err != nil {
		return err
	}

	return nil
}
