package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoCN, function for connect to the database */
var MongoCN = ConnectionBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://apps:Nels0mar@go-twittor-nofm.lc0y4.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/* ConnectionBD, function for connect to the database */
func ConnectionBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Success connection to database")
	return client
}

/* CheckConnection, check status to the database */
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return 1
}
