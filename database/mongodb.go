package mongodb

import (
	"MessageSecure/models"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("goMessageSecure").Collection(collectionName)

	return collection
}

func AddDataToCollection(client *mongo.Client, msg models.Message) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	connectionErr := client.Connect(ctx)

	if connectionErr != nil {
		log.Fatal(connectionErr)
	}

	collection := GetCollection(client, "messages")
	docs := []interface{}{
		bson.D{{"id", msg.Id}, {"message", msg.Message}, {"created_at", time.Now()}},
	}

	res, insertErr := collection.InsertMany(ctx, docs)
	if insertErr != nil {
		log.Fatal(insertErr)
	}

	disconnectErr := client.Disconnect(ctx)
	if disconnectErr != nil {
		log.Fatal(disconnectErr)
	}

	fmt.Println("Inserted documents: ", res.InsertedIDs)

	return nil
}

func FindDataFromCollection(client *mongo.Client, id string) bool {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	connectionErr := client.Connect(ctx)

	if connectionErr != nil {
		log.Fatal(connectionErr)
	}

	collection := GetCollection(client, "messages")
	msg := collection.FindOne(ctx, bson.D{{"id", id}})

	return msg != nil
}

func GetDataFromCollectionAndDelete(client *mongo.Client, id string) models.Message {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	connectionErr := client.Connect(ctx)

	if connectionErr != nil {
		log.Fatal(connectionErr)
	}

	collection := GetCollection(client, "messages")
	msg := collection.FindOne(ctx, bson.D{{"id", id}})

	// Get Message
	message := models.Message{}
	decodeErr := msg.Decode(&message)
	if decodeErr != nil {
		log.Print(decodeErr)
	}

	// Delete Entry
	collection.FindOneAndDelete(ctx, bson.D{{"id", id}})

	disconnectErr := client.Disconnect(ctx)
	if disconnectErr != nil {
		log.Fatal(disconnectErr)
	}

	return message
}
