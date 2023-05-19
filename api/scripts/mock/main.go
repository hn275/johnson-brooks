package main

import (
	"context"
	"fmt"
	"jb/database"
	"jb/lib/img"
	"jb/mod/auth"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

var data []database.Product = []database.Product{
	{
		Title:       "hangboard",
		Material:    "plastic",
		Description: "Lorem fugit similique nesciunt soluta architecto Amet ullam quaerat velit?",
		Price:       123,
		Variants: []database.ProductVariant{
			{
				ThumbnailData: getImage("classic").Bytes(),
				Color:         "#fefefe",
				Inventory:     12,
			},
		},
	},
	{
		Title:       "Monorail2",
		Material:    "plastic",
		Description: "Lorem fugit similique nesciunt soluta architecto Amet ullam quaerat velit?",
		Price:       45,
		Variants: []database.ProductVariant{
			{
				ThumbnailData: getImage("monorail1").Bytes(),
				Color:         "#484848",
				Inventory:     12,
			},
			{
				ThumbnailData: getImage("monorail2").Bytes(),
				Color:         "#000000",
				Inventory:     0,
			},
		},
	},
	{
		Title:       "Monorail2",
		Material:    "plastic",
		Description: "Lorem fugit similique nesciunt soluta architecto Amet ullam quaerat velit?",
		Price:       45,
		Variants: []database.ProductVariant{
			{
				ThumbnailData: getImage("oak").Bytes(),
				Color:         "#484848",
				Inventory:     12,
			},
			{
				ThumbnailData: getImage("hangboard").Bytes(),
				Color:         "#000000",
				Inventory:     0,
			},
			{
				ThumbnailData: getImage("oak").Bytes(),
				Color:         "#ffffff",
				Inventory:     0,
			},
		},
	},
}

func main() {
	db := database.New()
	defer db.Close()

	fmt.Print("Dropping existing data")
	if err := db.Drop(context.TODO()); err != nil {
		log.Fatal(err)
	}

	fmt.Print("\tOK\n\n")

	fmt.Println("Mock data:")

	fmt.Printf("\tadmin")
	createAdmin(db)

	for _, v := range data {
		if _, err := db.Collection(database.Hangboards).InsertOne(context.TODO(), &v); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("\tOK\n")
}

func getImage(i string) *img.Img {
	f, err := os.ReadFile("./scripts/mock/" + i + ".jpg")
	if err != nil {
		log.Fatal(err)
	}

	image, err := img.FromBytes(f)
	if err != nil {
		log.Fatal(err)
	}
	return image
}

func createMonoRails(db *database.Database, src map[string]string) {
	variants := []database.ProductVariant{}

	for k, v := range src {
		f, err := os.ReadFile("./scripts/mock/" + v)
		if err != nil {
			log.Fatal(err)
		}
		image, err := img.FromBytes(f)
		if err != nil {
			log.Fatal(err)
		}
		variant := database.ProductVariant{
			Thumbnail:     image.Base64(),
			ThumbnailData: image.Bytes(),
			Color:         "#ffffff",
			Inventory:     uint16(len(k)),
		}
		variants = append(variants, variant)
	}

	product := database.Product{
		Variants:    variants,
		Title:       "Monorail",
		Material:    "plastic",
		Description: "Lorem fugit similique nesciunt soluta architecto Amet ullam quaerat velit?",
		Price:       123,
	}

	trx := db.Collection(database.Hangboards)
	if _, err := trx.InsertOne(context.TODO(), &product); err != nil {
		log.Fatal(err)
	}
}

func createHangboard(db *database.Database, img string) {
	imgSrc := fmt.Sprintf("./scripts/mock/%s.jpg", img)
	f, err := os.ReadFile(imgSrc)
	if err != nil {
		log.Fatal(err)
	}

	prod := buildProduct(f, img)
	trx := db.Collection(database.Hangboards)
	_, err = trx.InsertOne(context.TODO(), &prod)
	if err != nil {
		log.Fatal(err)
	}
}

func buildProduct(src []byte, title string) database.Product {
	image, err := img.FromBytes(src)

	if err != nil {
		log.Fatal(err)
	}

	if err := image.Resize(img.Thumbnail); err != nil {
		log.Fatal(err)
	}

	return database.Product{
		Variants: []database.ProductVariant{
			{
				Thumbnail:     image.Base64(),
				ThumbnailData: image.Bytes(),
				Color:         "#1e1e1e",
				Inventory:     434,
			},
			{
				Thumbnail:     image.Base64(),
				ThumbnailData: image.Bytes(),
				Color:         "#ffffff",
				Inventory:     434,
			},
		},
		Title:       title,
		Price:       234,
		Material:    "test",
		Description: "Amet voluptates ipsum ea natus suscipit! Rerum unde quam dolores?",
	}
}

func createAdmin(db *database.Database) {
	cred := auth.Credentials{Username: "foo", Password: "bar"}

	hashed, err := bcrypt.GenerateFromPassword([]byte(cred.Password), 10)
	if err != nil {
		log.Fatal(err)
	}

	plain := cred.Password
	cred.Password = string(hashed)

	result, err := db.Collection(database.Admin).InsertOne(context.TODO(), &cred)
	if err != nil {
		panic(err)
	}

	fmt.Print("\t\tOK - added user")
	fmt.Printf(`
			username: %s
			password: %s
			id:       %v
		`, cred.Username, plain, result.InsertedID)
}
