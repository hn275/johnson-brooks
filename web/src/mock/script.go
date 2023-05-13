package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type MockData struct {
	ID          string  `json:"id"`
	Src         string  `json:"src"`
	Title       string  `json:"title"`
	Material    string  `json:"material"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Inventory   uint16  `json:"inventory"`
}

func main() {
	imgFile, err := os.ReadFile("./assets/monorail.jpg")
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer

	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	defer encoder.Close()

	encoder.Write(imgFile)

	var data []MockData
	for i := 0; i < 10; i++ {
		data = append(data, MockData{
			ID:          fmt.Sprintf("%d", i),
			Title:       "Mono rail",
			Src:         buf.String(),
			Material:    "Plastic",
			Description: "Dolor soluta consequatur reprehenderit accusantium voluptatibus Quis ullam corporis repellendus consectetur expedita? Saepe maxime minus dolor qui numquam. Deserunt laudantium libero aperiam nobis obcaecati eligendi. Debitis in cumque delectus dicta",
			Price:       99.99,
			Inventory:   12,
		})
	}

	dataJson, err := json.Marshal(&data)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("./data.json", dataJson, 0666); err != nil {
		log.Fatal(err)
	}

	log.Println("Done")
}
