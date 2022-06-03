package models

import "gopkg.in/mgo.v2/bson"

type Posts struct {
	PostID         bson.ObjectId       `json:"_id,omitempty" bson:"_id,omitempty,"`
	Title          string              `json:"title" bson:"title"`
	Description    string              `json:"description" bson:"description"`
	TopicID        bson.ObjectId       `json:"topicID,omitempty" bson:"topicID,omitempty"`
	UserID         bson.ObjectId       `json:"userID,omitempty" bson:"userID,omitempty"`
	Likes          int32               `json:"likes" bson:"likes"`
	PostReputation bson.ObjectId       `json:"postReputation,omitempty" bson:"postReputation,omitempty"`
	CommentsCount  int32               `json:"commentsCount" bson:"commentsCount"`
	CreatedAt      bson.MongoTimestamp `json:"createdAt" bson:"createdAt"`
}
