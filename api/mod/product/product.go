package product

import (
	"context"
	"encoding/json"
	"jb/database"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	var product []database.Product
	if err := getAllProducts(&product); err != nil {
		panic(err)
	}

	for i := range product {
		if err := product[i].Serializer(); err != nil {
			panic(err)
		}
	}

	log.Println("lasjdflkjasdf")
	for _, p := range product {
		log.Println(p.ID)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&product); err != nil {
		panic(err)
	}
}

func getAllProducts(products *[]database.Product) error {
	db := database.New()
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), database.Timeout)
	defer cancel()

	trx := db.Collection(database.Hangboards)
	cur, err := trx.Find(ctx, bson.D{})
	if err != nil {
		return err
	}

	return cur.All(ctx, products)
}
