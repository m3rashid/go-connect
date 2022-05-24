package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Avatars struct {
	AvatarID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID       primitive.ObjectID `json:"userID,omitempty" bson:"userID"`
	Sex          string             `json:"sex"`          // ENUM('man', 'woman') NOT NULL,
	FaceColor    string             `json:"faceColor"`    // VARCHAR(191) NOT NULL,
	EarSize      string             `json:"earSize"`      // ENUM('small', 'big') NOT NULL,
	HairColor    string             `json:"hairColor"`    // VARCHAR(191) NOT NULL,
	HairStyle    string             `json:"hairStyle"`    // ENUM('normal', 'thick', 'mohawk', 'womanLong', 'womanShort') NOT NULL,
	HatColor     string             `json:"hatColor"`     // VARCHAR(191) NOT NULL,
	HatStyle     string             `json:"hatStyle"`     // ENUM('none', 'beanie', 'turban') NOT NULL,
	GlassesStyle string             `json:"glassesStyle"` // ENUM('none', 'round', 'square') NOT NULL,
	NoseStyle    string             `json:"noseStyle"`    // ENUM('short', 'long', 'round') NOT NULL,
	MouthStyle   string             `json:"mouthStyle"`   // ENUM('laugh', 'smile', 'peace') NOT NULL,
	ShirtStyle   string             `json:"shirtStyle"`   // ENUM('hoody', 'short', 'polo') NOT NULL,
	ShirtColor   string             `json:"shirtColor"`   // VARCHAR(191) NOT NULL,
	BgColor      string             `json:"bgColor"`      // VARCHAR(191) NOT NULL,
	IsGradient   string             `json:"isGradient"`   // BOOLEAN NOT NULL,
	CreatedAt    primitive.DateTime `json:"createdAt"`
	UpdatedAt    primitive.DateTime `json:"updatedAt"`
}
