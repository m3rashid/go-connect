package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comments struct {
	CommentID primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID    primitive.ObjectID  `json:"userID,omitempty" bson:"userID,omitempty"`
	PostID    primitive.ObjectID  `json:"postID,omitempty" bson:"postID,omitempty"`
	Text      string              `json:"text" bson:"text"`
	CreatedAt primitive.Timestamp `json:"createdAt" bson:"createdAt"`
}
