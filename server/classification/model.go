package classification

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection
var ctx = context.TODO()

type Topic struct {
	classificationID primitive.ObjectID   `bson:"_id"`
	identityAttack   primitive.Decimal128 `bson:"identity_attack"`
	insult           primitive.Decimal128 `bson:"insult"`
	obscene          primitive.Decimal128 `bson:"obscene"`
	severeToxicity   primitive.Decimal128 `bson:"severe_toxicity"`
	sexualExplicit   primitive.Decimal128 `bson:"sexual_explicit"`
	threat           primitive.Decimal128 `bson:"threat"`
	toxicity         primitive.Decimal128 `bson:"toxicity"`
	postID           primitive.ObjectID   `bson:"postID"`
	createdAt        primitive.DateTime   `bson:"createdAt"`
	updatedAt        primitive.DateTime   `bson:"updatedAt"`
}
