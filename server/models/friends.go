package models

import "gopkg.in/mgo.v2/bson"

type Friends struct {
	FriendshipID bson.ObjectId       `json:"_id,omitempty" bson:"_id,omitempty"`
	FromID       bson.ObjectId       `json:"fromID,omitempty" bson:"fromID,omitempty"`
	ToID         bson.ObjectId       `json:"toID,omitempty" bson:"toID,omitempty"`
	Status       string              `json:"status" bson:"status"`
	CreatedAt    bson.MongoTimestamp `json:"createdAt" bson:"createdAt"`
}
