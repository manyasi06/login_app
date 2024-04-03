package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var URI = "mongodb://michael:secret@localhost:27017/"

type DbService struct {
	Client *mongo.Client
}

var DBClient *DbService

func InitDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}
	DBClient = &DbService{
		client,
	}
}
