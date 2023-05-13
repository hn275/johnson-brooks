package database

import "go.mongodb.org/mongo-driver/bson/primitive"

// NOTE: `Thumbnail` is a base64 encoded string of bytearray of an image
type Product struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Thumbnail   string             `json:"thumbnail" bson:"thumbnail"`
	Description string             `json:"description" bson:"description"`
	Material    string             `json:"material" bson:"material"`
	Price       float32            `json:"price" bson:"price"`
	Inventory   uint16             `json:"inventory" bson:"inventory"`
}

// CLIMBING
type ClimbingProduct struct {
	Product  `bson:",inline"`
	Variants []ClimbingVariant `json:"variants" bson:"variants"`
}

type ClimbingVariant struct {
	Thumbnail string `json:"thumbnail" bson:"thumbnail"`
	Color     string `json:"color" bson:"color"`
}

// BOARDS
type BoardProduct struct {
	Product `bson:",inline"`
}
