package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bookmarks struct {
	BookmarkID primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Post       primitive.ObjectID `json:"post,omitempty"`
	User       primitive.ObjectID `json:"user"`
	CreatedAt  primitive.DateTime `json:"createdAt"`
	UpdatedAt  primitive.DateTime `json:"updatedAt"`
}
