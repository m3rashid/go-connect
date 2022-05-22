package avatar

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection
var ctx = context.TODO()

type Avatar struct {
	avatarID primitive.ObjectID `bson:"_id"`
	userID   primitive.ObjectID `bson:"userID"`

	sex          string `bson:"sex"`          // ENUM('man', 'woman') NOT NULL,
	faceColor    string `bson:"faceColor"`    // VARCHAR(191) NOT NULL,
	earSize      string `bson:"earSize"`      // ENUM('small', 'big') NOT NULL,
	hairColor    string `bson:"hairColor"`    // VARCHAR(191) NOT NULL,
	hairStyle    string `bson:"hairStyle"`    // ENUM('normal', 'thick', 'mohawk', 'womanLong', 'womanShort') NOT NULL,
	hatColor     string `bson:"hatColor"`     // VARCHAR(191) NOT NULL,
	hatStyle     string `bson:"hatStyle"`     // ENUM('none', 'beanie', 'turban') NOT NULL,
	glassesStyle string `bson:"glassesStyle"` // ENUM('none', 'round', 'square') NOT NULL,
	noseStyle    string `bson:"noseStyle"`    // ENUM('short', 'long', 'round') NOT NULL,
	mouthStyle   string `bson:"mouthStyle"`   // ENUM('laugh', 'smile', 'peace') NOT NULL,
	shirtStyle   string `bson:"shirtStyle"`   // ENUM('hoody', 'short', 'polo') NOT NULL,
	shirtColor   string `bson:"shirtColor"`   // VARCHAR(191) NOT NULL,
	bgColor      string `bson:"bgColor"`      // VARCHAR(191) NOT NULL,
	isGradient   string `bson:"isGradient"`   // BOOLEAN NOT NULL,

	createdAt primitive.DateTime `bson:"createdAt"`
	updatedAt primitive.DateTime `bson:"updatedAt"`
}
