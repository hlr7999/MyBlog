package model

import (
	"github.com/globalsign/mgo/bson"
)

const CommentC = "comments"

type Comment struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	AID      string        `bson:"aid" json:"aid"`
	UID      string        `bson:"uid" json:"uid"`
	Username string        `bson:"username" json:"username"`
	Time     string        `bson:"time" json:"time"`
	Avatar   string        `bson:"avatar" json:"avatar"`
	Content  string        `bson:"content" json:"content"`
}
