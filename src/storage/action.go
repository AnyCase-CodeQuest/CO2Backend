package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

import log "github.com/sirupsen/logrus"

func Create(event Event) error {
	InitClient()
	_, err := GetCollection().InsertOne(context.TODO(), event)
	if err != nil {
		log.Fatalln("Error on inserting new sensor value", err)
	}
	log.Debugf("Received event and stored to DB: %+v", event)
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

func GetLatestBy(id string) (Event, error) {
	var event Event
	filter := bson.M{"deviceid": id}
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

func GetSensorList() SensorList {
	var list SensorList
	filter := bson.M{}
	ctx, _ := context.WithTimeout(context.TODO(), 30*time.Second)
	var findOptions options.FindOptions
	findOptions.SetSort(bson.M{"_id": -1})
	cursor, err := GetCollection().Find(ctx, filter, &findOptions)

	if err != nil {
		return list
	}
	for cursor.Next(ctx) {
		var result Sensor
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		list.data = append(list.data, result)
	}
	defer cursor.Close(ctx)
	return list
}
