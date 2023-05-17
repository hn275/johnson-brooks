package database

import (
	"errors"
	"jb/lib/img"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PRODUCT
type Product struct {
	ObjectID      primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	ID            string             `json:"id" bson:"-"`
	Title         string             `json:"title" bson:"title"`
	Thumbnail     string             `json:"thumbnail" bson:"-"` // base64 encoding of bytearray
	ThumbnailData []byte             `json:"-" bson:"thumbnail"`
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
	Variant       string `json:"variant" bson:"variant"`
	Thumbnail     string `json:"thumbnail" bson:"-"` // base64 encoding of bytearray
	ThumbnailData []byte `json:"-" bson:"thumbnail"`
	Color         string `json:"color" bson:"color"`
	Inventory     uint16 `json:"inventory" bson:"inventory"`
}

func (p *ClimbingProduct) Serializer() error {
	if err := p.Product.Serializer(); err != nil {
		return err
	}

	for _, variant := range p.Variants {
		variant.Serializer()
	}

	return nil
}

func (p *ClimbingVariant) Serializer() error {
	// get Thumbnail base64
	i, err := img.FromBytes(p.ThumbnailData)
	if err != nil {
		return err
	}
	p.Thumbnail = i.Base64()
	return nil

}

func (p *Product) Serializer() error {
	// get Thumbnail base64
	i, err := img.FromBytes(p.ThumbnailData)
	if err != nil {
		return err
	}
	p.Thumbnail = i.Base64()

	// get id
	if p.ObjectID == primitive.NilObjectID {
		return errors.New("nil objectID")
	}
	p.ID = p.ObjectID.Hex()

	return nil
}
