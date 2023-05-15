package auth

import (
	"context"
	"errors"
	"jb/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrAuthFailed = errors.New("auth failed")
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

type authDatabase struct {
	*database.Database
}

func newDb() *authDatabase {
	return &authDatabase{database.New()}
}

func (db *authDatabase) addSessionToUser(id primitive.ObjectID, sessionID string) error {
	filter := bson.D{{Key: "_id", Value: id}}
	value := bson.D{{Key: "session.sessionID", Value: sessionID}}
	ops := bson.D{{Key: "$set", Value: value}}

	ctx, cancel := context.WithTimeout(context.Background(), database.Timeout)
	defer cancel()

	_, err := db.Collection(database.Admin).UpdateOne(ctx, filter, ops)
	return err

}

func (db *authDatabase) findUser(user *Credentials) *mongo.SingleResult {
	col := db.Collection(database.Admin)

	ctx, cancel := context.WithTimeout(context.Background(), database.Timeout)
	defer cancel()

	filter := bson.D{{Key: "username", Value: user.Username}}

	return col.FindOne(ctx, filter)
}

func (db *authDatabase) authFailed(id primitive.ObjectID) error {
	col := db.Collection(database.Admin)

	ctx, cancel := context.WithTimeout(context.Background(), database.Timeout)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: id}}
	value := bson.D{{Key: "session.failedAttempts", Value: 1}}
	ops := bson.D{{Key: "$inc", Value: value}}

	err := col.FindOneAndUpdate(ctx, filter, ops).Err()
	if err != nil {
		return err
	}

	return nil
}

func (db *authDatabase) lockUser(id primitive.ObjectID) error {
	col := db.Collection(database.Admin)

	ctx, cancel := context.WithTimeout(context.Background(), database.Timeout)
	defer cancel()

	value := bson.D{
		{Key: "session.locked", Value: true},
		{Key: "session.sessionID", Value: ""},
	}
	update := bson.D{{Key: "$set", Value: value}}

	_, err := col.UpdateByID(ctx, id, update)
	return err
}
