package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Auths struct {
	UserID    bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName  string        `json:"userName" bson:"userName"`
	FirstName string        `json:"firstName" bson:"firstName"`
	LastName  string        `json:"lastName" bson:"lastName"`
	Email     string        `json:"email" bson:"email"`
	PhNumber  string        `json:"phNumber" bson:"phNumber"`
	Gender    string        `json:"gender" bson:"gender"`
	Password  string        `json:"password" bson:"password"`
	AvatarID  bson.ObjectId `json:"avatarID,omitempty" bson:"avatarID"`
}
