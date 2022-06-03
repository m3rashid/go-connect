package models

import "gopkg.in/mgo.v2/bson"

type Bookmarks struct {
	BookmarkID bson.ObjectId       `json:"_id,omitempty" bson:"_id"`
	Post       bson.ObjectId       `json:"post,omitempty" bson:"post"`
	User       bson.ObjectId       `json:"user" bson:"user"`
	CreatedAt  bson.MongoTimestamp `json:"createdAt" bson:"createdAt"`
}
