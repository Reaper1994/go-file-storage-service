package parser

import (
	"context"
	"encoding/xml"
	"fmt"
	"os"
	"time"

	"github.com/Reaper1994/go-file-storage-service/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type XMLData struct {
	Items []Item `xml:"item"`
}

type Item struct {
	ID    int    `xml:"id"`
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

func ParseAndSaveXML(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("unable to open file: %v", err)
	}
	defer file.Close()

	var data XMLData
	err = xml.NewDecoder(file).Decode(&data)
	if err != nil {
		return fmt.Errorf("unable to parse XML: %v", err)
	}

	collection := db.Client.Database("fileimporter").Collection("items")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	for _, item := range data.Items {
		_, err := collection.InsertOne(ctx, bson.M{
			"id":    item.ID,
			"name":  item.Name,
			"value": item.Value,
		})
		if err != nil {
			return fmt.Errorf("unable to insert data: %v", err)
		}
	}

	return nil
}
