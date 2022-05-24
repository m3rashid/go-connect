package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Likes struct {
	LikeID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"userID,omitempty" bson:"userID,omitempty"`
	PostID    primitive.ObjectID `json:"postID,omitempty" bson:"postID,omitempty"`
	CreatedAt primitive.DateTime `json:"createdAt"`
	UpdatedAt primitive.DateTime `json:"updatedAt"`
}
