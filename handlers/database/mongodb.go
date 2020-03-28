package mongohandler

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createInstance() *mongo.Client {

	/*
		TODO : the mongo uri should be replaced by configuration parameter
	*/
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass%20Community&ssl=false"))

	if err != nil {
		log.Fatal(err)
	}

	return client
}
