package util

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var MClient *mongo.Client

func MInit() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	fmt.Println(uri)
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/#environment-variable")
	}
	MClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)

	}
	// 检查连接
	err = MClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func GetMClient() *mongo.Client {
	if MClient == nil {
		MInit()
	}
	return MClient
}
