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
	path_sour := "/www/craw.domp4.cc/.env"
	if err := godotenv.Load(path_sour); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	fmt.Println(uri)
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/#environment-variable")
	}
	client_s, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	MClient = client_s
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
