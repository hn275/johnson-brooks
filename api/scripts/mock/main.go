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

	fmt.Printf("\tdata")
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
