package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const URI = "mongodb+srv://j:rootroot@cluster0.rj0tg.mongodb.net/test"
const DATABASE = "twitter_mobileapp"

var clientURI = options.Client().ApplyURI(URI)

var CM = connectToDB()

func errorFatal(err error) {
	log.Fatal(err.Error())
}

func connectToDB() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), clientURI)
	if err != nil {
		errorFatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		errorFatal(err)
	}

	return client
}

func CheckConnection() int {
	err := CM.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}

	return 1
}
