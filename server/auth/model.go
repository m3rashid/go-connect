package auth

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection
var ctx = context.TODO()

type Auth struct {
	userID    primitive.ObjectID `bson:"_id,omitempty"`
	userName  string             `bson:"userName"`
	firstName string             `bson:"firstName"`
	lastName  string             `bson:"lastName"`
	email     string             `bson:"email"`
	phNumber  string             `bson:"phNumber"`
	gender    string             `bson:"gender"`
	password  string             `bson:"password"`
	avatarID  primitive.ObjectID `bson:"avatarID,omitempty"`
	dob       primitive.DateTime `bson:"dob"`
	createdAt primitive.DateTime `bson:"createdAt"`
	updatedAt primitive.DateTime `bson:"updatedAt"`
}
