package ping

import (
	"context"
	"jb/database"
	"log"
	"net/http"
	"time"
)

type Ping struct {
	Hello string `bson:"hello"`
}

func Controller() http.Handler {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = database.New(ctx)
	db := database.New(ctx)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := Ping{"world"}

		res, err := db.Collection(database.Boards).InsertOne(context.TODO(), &buf)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusCreated)
		log.Println(res.InsertedID)
	})
}
