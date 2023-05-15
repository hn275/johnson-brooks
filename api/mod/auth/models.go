package auth

import (
	"context"
	"jb/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Credentials struct {
	ID       primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Session  Session            `json:"-" bson:"session"`
}

type Session struct {
	SessionID      string `bson:"sessionID"`
	FailedAttempts uint8  `bson:"failedAttempts"`
	Locked         bool   `bson:"locked"`
}

type AuthDatabase struct {
	*database.Database
}

func newDb() *AuthDatabase {
	return &AuthDatabase{database.New()}
}

func (db *AuthDatabase) addSessionToUser(username, sessionID string) error {
	filter := bson.D{{Key: "username", Value: username}}
	update := bson.D{{Key: "sessionID", Value: sessionID}}
	ops := bson.D{{Key: "$set", Value: update}}

	ctx, cancel := context.WithTimeout(context.Background(), database.Timeout)
	defer cancel()

	_, err := db.Collection(database.Admin).UpdateOne(ctx, filter, ops)
	return err

}

func (db *AuthDatabase) findUser(cred *Credentials) *mongo.SingleResult {
	col := db.Collection(database.Admin)

	ctx, cancel := context.WithTimeout(context.Background(), database.Timeout)
	defer cancel()

	filter := bson.D{{Key: "username", Value: cred.Username}}

	return col.FindOne(ctx, filter)
}
