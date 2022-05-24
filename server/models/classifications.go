package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Classifications struct {
	ClassificationID primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	IdentityAttack   primitive.Decimal128 `json:"identity_attack"`
	Insult           primitive.Decimal128 `json:"insult"`
	Obscene          primitive.Decimal128 `json:"obscene"`
	SevereToxicity   primitive.Decimal128 `json:"severe_toxicity"`
	SexualExplicit   primitive.Decimal128 `json:"sexual_explicit"`
	Threat           primitive.Decimal128 `json:"threat"`
	Toxicity         primitive.Decimal128 `json:"toxicity"`
	PostID           primitive.ObjectID   `json:"postID"`
	CreatedAt        primitive.DateTime   `json:"createdAt"`
	UpdatedAt        primitive.DateTime   `json:"updatedAt"`
}
