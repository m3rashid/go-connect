package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Classifications struct {
	ClassificationID primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	IdentityAttack   primitive.Decimal128 `json:"identity_attack" bson:"identity_attack"`
	Insult           primitive.Decimal128 `json:"insult" bson:"insult"`
	Obscene          primitive.Decimal128 `json:"obscene" bson:"obscene"`
	SevereToxicity   primitive.Decimal128 `json:"severe_toxicity" bson:"severe_toxicity"`
	SexualExplicit   primitive.Decimal128 `json:"sexual_explicit" bson:"sexual_explicit"`
	Threat           primitive.Decimal128 `json:"threat" bson:"threat"`
	Toxicity         primitive.Decimal128 `json:"toxicity" bson:"toxicity"`
	PostID           primitive.ObjectID   `json:"postID" bson:"postID"`
	CreatedAt        primitive.Timestamp  `json:"createdAt" bson:"createdAt"`
}
