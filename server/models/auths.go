package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Auths struct {
	UserID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName  string             `json:"userName"`
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
	Email     string             `json:"email"`
	PhNumber  string             `json:"phNumber"`
	Gender    string             `json:"gender"`
	Password  string             `json:"password"`
	AvatarID  primitive.ObjectID `json:"avatarID,omitempty"`
	Dob       primitive.DateTime `json:"dob"`
	CreatedAt primitive.DateTime `json:"createdAt"`
	UpdatedAt primitive.DateTime `json:"updatedAt"`
}
