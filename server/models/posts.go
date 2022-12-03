package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Posts struct {
	PostID         primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty,"`
	Title          string              `json:"title" bson:"title"`
	Description    string              `json:"description" bson:"description"`
	TopicID        primitive.ObjectID  `json:"topicID,omitempty" bson:"topicID,omitempty"`
	UserID         primitive.ObjectID  `json:"userID,omitempty" bson:"userID,omitempty"`
	Likes          int32               `json:"likes" bson:"likes"`
	PostReputation primitive.ObjectID  `json:"postReputation,omitempty" bson:"postReputation,omitempty"`
	CommentsCount  int32               `json:"commentsCount" bson:"commentsCount"`
	CreatedAt      primitive.Timestamp `json:"createdAt" bson:"createdAt"`
}
