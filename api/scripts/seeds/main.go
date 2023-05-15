package main

import (
	"context"
	"jb/database"
	"jb/mod/auth"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	createAdmin()
}

func createAdmin() {
	cred := auth.Credentials{Username: "foo", Password: "bar"}
	db := database.New()
	defer db.Close()

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

	log.Printf("Added user\nusername: %s\npassword: %s\nid: %v]", cred.Username, plain, result.InsertedID)
}
