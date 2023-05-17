package database

import (
	"errors"
	"jb/lib/img"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PRODUCT
type Product struct {
	ObjectID    primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	ID          string             `json:"id" bson:"-"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Material    string             `json:"material" bson:"material"`
	Price       float32            `json:"price" bson:"price"`
	Variants    []ProductVariant   `json:"variants" bson:"variants"`
}

type ProductVariant struct {
	Variant       string `json:"variant" bson:"variant"`
	Thumbnail     string `json:"thumbnail" bson:"-"` // base64 encoding of bytearray
	ThumbnailData []byte `json:"-" bson:"thumbnail"`
	Color         string `json:"color" bson:"color"`
	Inventory     uint16 `json:"inventory" bson:"inventory"`
}

func (p *Product) Serializer() error {
	// serializer id
	p.ID = p.ObjectID.Hex()

	// get Thumbnail base64 for all variants
	for i := range p.Variants {
		image, err := img.FromBytes(p.Variants[i].ThumbnailData)
		if err != nil {
			return err
		}

		p.Variants[i].Thumbnail = image.Base64()
	}

	// get id
	if p.ObjectID == primitive.NilObjectID {
		return errors.New("nil objectID")
	}
	p.ID = p.ObjectID.Hex()

	return nil
}
