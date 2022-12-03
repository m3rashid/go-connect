package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bookmarks struct {
	BookmarkID primitive.ObjectID  `json:"_id,omitempty" bson:"_id"`
	Post       primitive.ObjectID  `json:"post,omitempty" bson:"post"`
	User       primitive.ObjectID  `json:"user" bson:"user"`
	CreatedAt  primitive.Timestamp `json:"createdAt" bson:"createdAt"`
}
