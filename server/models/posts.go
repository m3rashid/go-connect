package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Posts struct {
	PostID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty,"`
	Title          string             `json:"title"`
	Description    string             `json:"description"`
	TopicID        primitive.ObjectID `json:"topicID,omitempty" bson:"topicID,omitempty"`
	UserID         primitive.ObjectID `json:"userID,omitempty" bson:"userID,omitempty"`
	Likes          int32              `json:"likes"`
	PostReputation primitive.ObjectID `json:"postReputation,omitempty" bson:"postReputation,omitempty"`
	CommentsCount  int32              `json:"commentsCount"`
	CreatedAt      primitive.DateTime `json:"createdAt"`
	UpdatedAt      primitive.DateTime `json:"updatedAt"`
}
