package auth

import (
	"context"
	"jb/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Credentials struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type AuthDatabase struct {
	*database.Database
}

func newDb() (*AuthDatabase) {
	return &AuthDatabase{database.New()}
}

func (db *AuthDatabase) addSessionToUser(username, sessionID string) (error) {
	filter := bson.D{{"username", username}}
	update := bson.D{{"$set", bson.D{{"sessionID", sessionID}}}}
	_, err := db.Collection(database.Admin).UpdateOne(context.Background(), filter, update)
	return err

}

func (db *AuthDatabase) findUser(cred *Credentials) (*mongo.SingleResult) {
	col := db.Collection(database.Admin)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{{Key: "username", Value: cred.Username}}

	return col.FindOne(ctx, filter)

}