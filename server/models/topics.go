package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Topics struct {
	TopicID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name"`
	CreatedAt primitive.DateTime `json:"createdAt"`
	UpdatedAt primitive.DateTime `json:"updatedAt"`
}
