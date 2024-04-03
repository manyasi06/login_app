package main

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var URI = "mongodb://michael:secret@localhost:27017/"

type migrationFunc func(currClient *mongo.Client) error

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your databases. You successfully connected to MongoDB!")

	migrates := []struct {
		migrationName  string
		collectionName string
		migration      migrationFunc
		rollback       migrationFunc
	}{
		{
			migrationName:  "Creating Migration Collection",
			collectionName: "migration",
			migration:      createUserCollection(client),
			rollback: func(currClient *mongo.Client) error {
				return nil
			},
		},
	}

	for _, m := range migrates {
		check, err := checkIfMigrationExistsAlready(client, m.collectionName, m.migrationName)
		if err != nil {
			log.Fatal(err)
		}

		if !check {
			err := createMigrationCollection(client, m.migrationName)
			if err != nil {
				log.Fatal(err)
			}
			err = m.migration(client)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func createMigrationCollection(client *mongo.Client, migrateName string) error {
	collection := client.Database("reddit").Collection("migration")
	one, err := collection.InsertOne(context.Background(), bson.D{{"name", migrateName}})
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("Successfully used %s", one.InsertedID))
	return nil
}

func checkIfMigrationExistsAlready(client *mongo.Client, collectionName string, migration string) (bool, error) {
	var result bson.D
	err := client.Database("reddit").Collection(collectionName).FindOne(context.Background(), bson.D{{"name", migration}}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func createUserCollection(client *mongo.Client) migrationFunc {
	return migrationFunc(func(currClient *mongo.Client) error {
		if client == nil {
			return errors.New("Client is not initiliazed")
		}

		db := client.Database("reddit")

		const name = "usercollection"
		collection := db.Collection(name)
		if collection == nil {
			err := db.CreateCollection(context.TODO(), name)
			if err != nil {
				return err
			}
		}

		return nil
	})

}
