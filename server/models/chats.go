package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chats struct {
	ChatID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"userID,omitempty" bson:"userID,omitempty"`
	UserName  string             `json:"userName"`
	Message   string             `json:"message"`
	CreatedAt primitive.DateTime `json:"createdAt"`
}
