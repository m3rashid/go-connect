package friends

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection
var ctx = context.TODO()

type Topic struct {
	friendshipID primitive.ObjectID `bson:"_id"`
	fromID       primitive.ObjectID `bson:"fromID"`
	toID         primitive.ObjectID `bson:"toID"`
	status       string             `bson:"status"`
	createdAt    primitive.DateTime `bson:"createdAt"`
	updatedAt    primitive.DateTime `bson:"updatedAt"`
}
