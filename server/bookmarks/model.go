package bookmarks

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection
var ctx = context.TODO()

type Bookmark struct {
	bookmarkID primitive.ObjectID `bson:"_id"`
	post       primitive.ObjectID `bson:"post"`
	user       primitive.ObjectID `bson:"user"`
	createdAt  primitive.DateTime `bson:"createdAt"`
	updatedAt  primitive.DateTime `bson:"updatedAt"`
}
