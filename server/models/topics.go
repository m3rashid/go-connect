package models

import "gopkg.in/mgo.v2/bson"

type Topics struct {
	TopicID   bson.ObjectId       `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string              `json:"name" bson:"name"`
	CreatedAt bson.MongoTimestamp `json:"createdAt" bson:"createdAt"`
}
