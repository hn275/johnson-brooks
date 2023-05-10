package database

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri string

const (
	Hangboards string = "hangboards"
	Monorails  string = "monorails"
	Boards     string = "boards"
)

type Database struct {
	*mongo.Database
}

func init() {
	uri = os.Getenv("MONGODB_URL")
	if uri == "" {
		panic("missing env [MONGODB_URL]")
	}
}

func New(ctx context.Context) *Database {
	stableAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(stableAPI)
	db, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	return &Database{db.Database("johnson-brooks")}
}

func (db *Database) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.Database.Client().Disconnect(ctx); err != nil {
		panic(err)
	}
}
