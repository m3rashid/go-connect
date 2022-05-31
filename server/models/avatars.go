package models

import "gopkg.in/mgo.v2/bson"

type Avatars struct {
	AvatarID     bson.ObjectId       `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID       bson.ObjectId       `json:"userID,omitempty" bson:"userID"`
	Sex          string              `json:"sex" bson:"sex"`                   // ENUM('man', 'woman') NOT NULL,
	FaceColor    string              `json:"faceColor" bson:"faceColor"`       // VARCHAR(191) NOT NULL,
	EarSize      string              `json:"earSize" bson:"earSize"`           // ENUM('small', 'big') NOT NULL,
	HairColor    string              `json:"hairColor" bson:"hairColor"`       // VARCHAR(191) NOT NULL,
	HairStyle    string              `json:"hairStyle" bson:"hairStyle"`       // ENUM('normal', 'thick', 'mohawk', 'womanLong', 'womanShort') NOT NULL,
	HatColor     string              `json:"hatColor" bson:"hatColor"`         // VARCHAR(191) NOT NULL,
	HatStyle     string              `json:"hatStyle" bson:"hatStyle"`         // ENUM('none', 'beanie', 'turban') NOT NULL,
	GlassesStyle string              `json:"glassesStyle" bson:"glassesStyle"` // ENUM('none', 'round', 'square') NOT NULL,
	NoseStyle    string              `json:"noseStyle" bson:"noseStyle"`       // ENUM('short', 'long', 'round') NOT NULL,
	MouthStyle   string              `json:"mouthStyle" bson:"mouthStyle"`     // ENUM('laugh', 'smile', 'peace') NOT NULL,
	ShirtStyle   string              `json:"shirtStyle" bson:"shirtStyle"`     // ENUM('hoody', 'short', 'polo') NOT NULL,
	ShirtColor   string              `json:"shirtColor" bson:"shirtColor"`     // VARCHAR(191) NOT NULL,
	BgColor      string              `json:"bgColor" bson:"bgColor"`           // VARCHAR(191) NOT NULL,
	IsGradient   string              `json:"isGradient" bson:"isGradient"`     // BOOLEAN NOT NULL,
	CreatedAt    bson.MongoTimestamp `json:"createdAt" bson:"createdAt"`
	UpdatedAt    bson.MongoTimestamp `json:"updatedAt" bson:"updatedAt"`
}
