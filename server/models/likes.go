package models

import "gopkg.in/mgo.v2/bson"

type Likes struct {
	LikeID    bson.ObjectId       `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID    bson.ObjectId       `json:"userID,omitempty" bson:"userID,omitempty"`
	PostID    bson.ObjectId       `json:"postID,omitempty" bson:"postID,omitempty"`
	CreatedAt bson.MongoTimestamp `json:"createdAt" bson:"createdAt"`
}
