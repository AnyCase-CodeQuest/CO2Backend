package storage

import (
	"context"
	"github.com/AnyCase-Company-LTD/CO2Backend/src/values"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	auth2 "go.mongodb.org/mongo-driver/x/mongo/driver/auth"
	"os"
)

var Client *mongo.Client

func InitClient() {
	if isConnectionActive() != true {
		log.Debug("made new connection")

		auth := options.Credential{
			Password:      os.Getenv(values.DbPassword),
			Username:      os.Getenv(values.DbUsername),
			AuthSource:    os.Getenv(values.DbName),
			AuthMechanism: auth2.SCRAMSHA1,
		}
		clientOptions := options.Client().ApplyURI(os.Getenv(values.DbConnection)).SetAuth(auth)
		var err error
		log.Debug("db connection try to connect")
		//Client, err = mongo.NewClient(clientOptions)
		//if err != nil {
		//	log.Errorf("db connection cant establish: %s", err.Error())
		//	log.Fatal(err)
		//}
		Client, err = mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Errorf("db connection connect error: %e", err.Error())
			log.Fatal(err)
		}
	}
}

func isConnectionActive() bool {
	if Client == nil {
		log.Debug("db connection not exist")
		return false
	}
	err := Client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Debug("db connection is die")
		return false
	} else {
		log.Debug("db connection is active")
		return true
	}
}

func GetDataCollection() *mongo.Collection {
	InitClient()
	return Client.Database(os.Getenv(values.DbName)).Collection(os.Getenv("COLLECTION_DATA"))
}

func GetSensorCollection() *mongo.Collection {
	InitClient()
	return Client.Database(os.Getenv(values.DbName)).Collection(os.Getenv("COLLECTION_SENSOR"))
}
