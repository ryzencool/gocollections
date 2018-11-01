package collections

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// MongoCon is mongo connection
func MongoCon() {

	client, err := mongo.NewClient("mongodb://localhost:27017")
	if err != nil {
		log.Printf("err is : %v", err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Printf("connect err is : %v", err)
	}
	collection := client.Database("gotest").Collection("test1")
	res, err := collection.InsertOne(context.Background(), map[string]string{"name": "zmy"})
	if err !=nil {
		log.Printf("insert err is: %v", err)
	}
	log.Printf("insert id is : %v", res.InsertedID)
}
