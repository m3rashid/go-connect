package chat

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection
var ctx = context.TODO()

type Topic struct {
	chatID    primitive.ObjectID `bson:"_id"`
	userID    primitive.ObjectID `bson:"userID"`
	userName  string             `bson:"userName"`
	message   string             `bson:"message"`
	createdAt primitive.DateTime `bson:"createdAt"`
}
