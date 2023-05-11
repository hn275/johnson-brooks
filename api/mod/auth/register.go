package auth

import (
	"context"
	"encoding/json"
	"jb/database"
	"jb/lib"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var cred Credentials
	if err := json.NewDecoder(r.Body).Decode(&cred); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.StdErr(err)
		return
	}
	defer r.Body.Close()

	if err := serializeUser(&cred); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.StdErr(err)
		return
	}

	result, err := createUser(&cred)
	if err != nil {
		switch err.(type) {
		case mongo.WriteError:
			w.WriteHeader(http.StatusBadRequest)
			msg := "User exists, try a different username."
			lib.NewErr(msg).HandleErr(w)
			return

		default:
			w.WriteHeader(http.StatusInternalServerError)
			lib.StdErr(err)
			return
		}
	}

	insertedID := result.InsertedID.(primitive.ObjectID).String()

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(insertedID))
}

func createUser(cred *Credentials) (*mongo.InsertOneResult, error) {
	// TODO: Set index so no 2 same usernames can exist
	db := database.New()
	defer db.Close()

	col := db.Collection(database.Admin)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return col.InsertOne(ctx, cred)
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
