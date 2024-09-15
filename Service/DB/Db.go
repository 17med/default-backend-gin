package DB

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var clinett *mongo.Client

func Init() {
	
	
	var uri string="mongodb://localhost:27017/test"
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	
	client, err := mongo.Connect(context.TODO(), opts)
	clinett=client
	if err != nil {
		panic(err)
	}
	defer func() {
		
	}()
	
	var result bson.M
	if err := client.Database("test").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")
}
func GetClient() *mongo.Client{
	return clinett
}
	
