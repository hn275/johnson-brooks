package auth

import (
	"context"
	"encoding/json"
	"jb/database"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// @controller
func Register(w http.ResponseWriter, r *http.Request) {
	var cred Credentials
	if err := json.NewDecoder(r.Body).Decode(&cred); err != nil {
		panic(err)
	}
	defer r.Body.Close()

	if err := serializeUser(&cred); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	log.Println(cred)
	result, err := createUser(&cred)
	if err != nil {
		panic(err)
	}

	log.Println(result.InsertedID)
	w.WriteHeader(http.StatusCreated)
}

func createUser(cred *Credentials) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db := database.New(ctx)
	defer db.Close()

	col := db.Collection(database.Admin)
	c, can := context.WithTimeout(context.Background(), 5*time.Second)
	defer can()
	return col.InsertOne(c, cred)
}

func serializeUser(cred *Credentials) error {
	plain := []byte(cred.Password)
	cost := 10
	hashed, err := bcrypt.GenerateFromPassword(plain, cost)
	if err != nil {
		return err
	}

	cred.Password = string(hashed)
	return nil
}
