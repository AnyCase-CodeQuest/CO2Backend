package storage

import (
	"context"
	"fmt"
	"log"
)

func Create(event Event) error {
	InitClient()
	_, err := Client.Database("co2").Collection("values").InsertOne(context.TODO(), event)
	if err != nil {
		log.Fatalln("Error on inserting new Hero", err)
	}
	log.Println(fmt.Sprintf("%+v", event))
	return nil
}

func GetLatest(count int) ([]Event, error) {
	var result []Event
	return result, &errorStorage{"No lines to show"}
}
