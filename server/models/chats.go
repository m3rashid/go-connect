package models

import "gopkg.in/mgo.v2/bson"

type Chats struct {
	ChatID    bson.ObjectId       `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID    bson.ObjectId       `json:"userID,omitempty" bson:"userID,omitempty"`
	UserName  string              `json:"userName" bson:"userName"`
	Message   string              `json:"message" bson:"message"`
	CreatedAt bson.MongoTimestamp `json:"createdAt" bson:"createdAt"`
}
