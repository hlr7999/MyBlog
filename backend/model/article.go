package model

import (
	"github.com/globalsign/mgo/bson"
)

const ArticleC = "articles"

type Article struct {
	ID           bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Title        string        `bson:"title" json:"title"`
	Year         int           `bson:"year" json:"year"`
	Month        int           `bson:"month" json:"month"`
	Day          int           `bson:"day" json:"day"`
	BrowseCount  int           `bson:"browseCount" json:"browseCount"`
	CommentCount int           `bson:"commentCount" json:"commentCount"`
	LikeCount    int           `bson:"likeCount" json:"likeCount"`
	CollectCount int           `bson:"collectCount" json:"collectCount"`
	ClassId      string		   `bson:"classId" json:"classId"`
	ClassName    string        `bson:"className" json:"className"`
	Description  string        `bson:"description" json:"description"`
	ImageCover   string        `bson:"imageCover" json:"imageCover"`
	Content      string        `bson:"content" json:"content"`
}
