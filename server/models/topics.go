package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Topics struct {
	TopicID   primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string              `json:"name" bson:"name"`
	CreatedAt primitive.Timestamp `json:"createdAt" bson:"createdAt"`
}
