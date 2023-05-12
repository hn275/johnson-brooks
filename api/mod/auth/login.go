package auth

import (
	"context"
	"encoding/json"
	"jb/database"
	"jb/lib"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"github.com/teris-io/shortid"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	result := struct{
		Username string
		Password string
	}{}
	err := findUser(&cred).Decode(&result)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		msg := "No Account Found"
		lib.NewErr(msg).HandleErr(w)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(cred.Password))

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		msg := "Invalid password"
		lib.NewErr(msg).HandleErr(w)
		return
	}

	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	session, err := sid.Generate()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := "Error generating SID (session)"
		lib.NewErr(msg).HandleErr(w)
		return
	}
	
	addSessionToUser(session)

	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(session))
}

func findUser(cred *Credentials) (*mongo.SingleResult) {
	db := database.New()
	defer db.Close()

	col := db.Collection(database.Admin)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{{"Username", cred.Username}}

	return col.FindOne(ctx, filter)
}

func addSessionToUser(session string) (*mongo.InsertOneResult, error) {
	//TODO: Add session to user
}
