package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func Create(event Event) error {
	InitClient()
	_, err := GetCollection().InsertOne(context.TODO(), event)
	if err != nil {
		log.Fatalln("Error on inserting new sensor value", err)
	}
	log.Println(fmt.Sprintf("%+v", event))
	return nil
}

func GetLatest() (Event, error) {
	var event Event
	filter := bson.M{}
	var findOneOptions options.FindOneOptions
	findOneOptions.SetSort(bson.M{"_id": -1})
	documentReturned := GetCollection().FindOne(context.TODO(), filter, &findOneOptions)

	documentReturned.Decode(&event)

	if documentReturned.Err() != nil {
		log.Println(documentReturned.Err())
		return event, &errorStorage{"No lines to show"}
	} else {
		return event, nil
	}
}
