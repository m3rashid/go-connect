package topics

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection
var ctx = context.TODO()

type Topic struct {
	topicID   primitive.ObjectID `bson:"_id"`
	name      string             `bson:"name"`
	createdAt primitive.DateTime `bson:"createdAt"`
	updatedAt primitive.DateTime `bson:"updatedAt"`
}
