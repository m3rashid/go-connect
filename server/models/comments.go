package models

import "gopkg.in/mgo.v2/bson"

type Comments struct {
	CommentID bson.ObjectId       `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID    bson.ObjectId       `json:"userID,omitempty" bson:"userID,omitempty"`
	PostID    bson.ObjectId       `json:"postID,omitempty" bson:"postID,omitempty"`
	Text      string              `json:"text" bson:"text"`
	CreatedAt bson.MongoTimestamp `json:"createdAt" bson:"createdAt"`
	UpdatedAt bson.MongoTimestamp `json:"updatedAt" bson:"updatedAt"`
}
