package posts

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection
var ctx = context.TODO()

type Post struct {
	postID         primitive.ObjectID `bson:"_id,omitempty"`
	title          string             `bson:"title"`
	description    string             `bson:"description"`
	topicID        primitive.ObjectID `bson:"topicID"`
	userID         primitive.ObjectID `bson:"userID"`
	likes          int32              `bson:"likes"`
	postReputation primitive.ObjectID `bson:"postReputation"`
	commentsCount  int32              `bson:"commentsCount"`
	createdAt      primitive.DateTime `bson:"createdAt"`
	updatedAt      primitive.DateTime `bson:"updatedAt"`
}
