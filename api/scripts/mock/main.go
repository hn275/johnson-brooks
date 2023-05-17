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

func main() {
	db := database.New()
	defer db.Close()

	fmt.Println("Dropping existing data")
	if err := db.Drop(context.TODO()); err != nil {
		log.Fatal(err)
	}

	createAdmin(db)
	createHangboard(db, "oak")
	createHangboard(db, "hangboard")
}

func createHangboard(db *database.Database, img string) {
	imgSrc := fmt.Sprintf("./scripts/mock/%s.jpg", img)
	f, err := os.ReadFile(imgSrc)
	if err != nil {
		log.Fatal(err)
	}

	prod := buildProduct(f, img)
	trx := db.Collection(database.Hangboards)
	i, err := trx.InsertOne(context.TODO(), &prod)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nAdded hangboard\ntitle: %v\nid: %v\n", prod.Title, i.InsertedID)
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
				Variant:       "main",
				Thumbnail:     image.Base64(),
				ThumbnailData: image.Bytes(),
				Color:         "#1e1e1e",
				Inventory:     420,
			},
			{
				Variant:       "white",
				Thumbnail:     image.Base64(),
				ThumbnailData: image.Bytes(),
				Color:         "#ffffff",
				Inventory:     420,
			},
		},
		Title:       title,
		Price:       69,
		Material:    "test",
		Description: "Amet voluptates ipsum ea natus suscipit! Rerum unde quam dolores?",
	}
}

func createAdmin(db *database.Database) {
	fmt.Printf("\nMocking user\n")
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

	fmt.Printf("Added user\nusername: %s\npassword: %s\nid: %v\n", cred.Username, plain, result.InsertedID)
}
