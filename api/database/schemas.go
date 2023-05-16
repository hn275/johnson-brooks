package database

import "go.mongodb.org/mongo-driver/bson/primitive"

// PRODUCT
type Product struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title         string             `json:"title" bson:"title"`
	Thumbnail     string             `json:"thumbnail" bson:"-"` // base64 encoding of bytearray
	ThumbnailData []byte             `json:"-" bson:"thumbnailBin"`
	Description   string             `json:"description" bson:"description"`
	Material      string             `json:"material" bson:"material"`
	Price         float32            `json:"price" bson:"price"`
	Inventory     uint16             `json:"inventory" bson:"inventory"`
}

type ClimbingProduct struct {
	Product  `bson:",inline"`
	Variants []ClimbingVariant `json:"variants" bson:"variants"`
}

type ClimbingVariant struct {
	Variant   string `json:"variant" bson:"variant"`
	Thumbnail string `json:"thumbnail" bson:"thumbnail"`
	Color     string `json:"color" bson:"color"`
	Inventory uint16 `json:"inventory" bson:"inventory"`
}

type BoardProduct struct {
	Product `bson:",inline"`
}
