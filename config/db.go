package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/shraddha0602/golang-mongodb/controllers"
)

func Connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	//ping db
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Error while connecting db", err)
	}

	log.Println("Connected\n")

	//create db newMongo
	db := client.Database("books")

	//db instance
	controllers.BookCollection(db)
	return
}
