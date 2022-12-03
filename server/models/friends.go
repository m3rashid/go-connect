package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Friends struct {
	FriendshipID primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	FromID       primitive.ObjectID  `json:"fromID,omitempty" bson:"fromID,omitempty"`
	ToID         primitive.ObjectID  `json:"toID,omitempty" bson:"toID,omitempty"`
	Status       string              `json:"status" bson:"status"`
	CreatedAt    primitive.Timestamp `json:"createdAt" bson:"createdAt"`
}
