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
	createHangboard(db)
}

func createHangboard(db *database.Database) {
	fmt.Printf("\nMocking hangboard")
	f, err := os.ReadFile("./scripts/mock/oak.jpg")
	if err != nil {
		log.Fatal(err)
	}

	prod := buildProduct(f, "Hangboard")
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
		Thumbnail:     image.Base64(),
		ThumbnailData: image.Bytes(),
		Title:         title,
		Price:         69,
		Material:      "test",
		Inventory:     420,
		Description:   "Amet voluptates ipsum ea natus suscipit! Rerum unde quam dolores?",
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
