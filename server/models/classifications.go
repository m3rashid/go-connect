package models

import "gopkg.in/mgo.v2/bson"

type Classifications struct {
	ClassificationID bson.ObjectId       `json:"_id,omitempty" bson:"_id,omitempty"`
	IdentityAttack   bson.Decimal128     `json:"identity_attack" bson:"identity_attack"`
	Insult           bson.Decimal128     `json:"insult" bson:"insult"`
	Obscene          bson.Decimal128     `json:"obscene" bson:"obscene"`
	SevereToxicity   bson.Decimal128     `json:"severe_toxicity" bson:"severe_toxicity"`
	SexualExplicit   bson.Decimal128     `json:"sexual_explicit" bson:"sexual_explicit"`
	Threat           bson.Decimal128     `json:"threat" bson:"threat"`
	Toxicity         bson.Decimal128     `json:"toxicity" bson:"toxicity"`
	PostID           bson.ObjectId       `json:"postID" bson:"postID"`
	CreatedAt        bson.MongoTimestamp `json:"createdAt" bson:"createdAt"`
}
