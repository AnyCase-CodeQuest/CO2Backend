package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

var Client *mongo.Client

func InitClient() {
	if isConnectionActive() != true {
		log.Println("made new connection")
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
		auth := options.Credential{Password: "kube", Username: "root"}
		clientOptions.SetAuth(auth)
		var err error
		Client, err = mongo.NewClient(clientOptions)
		if err != nil {
			log.Fatal(err)
		}
		err = Client.Connect(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
	}
}

func isConnectionActive() bool {
	if Client == nil {
		return false
	}
	err := Client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return false
	} else {
		return true
	}
}

func GetCollection() *mongo.Collection {
	InitClient()
	return Client.Database("co2").Collection("values")
}
