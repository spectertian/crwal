package mongo_orm

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var MClient mongo.Client

type Default struct {
	ID primitive.ObjectID `bson:"_id" json:"id,omitempty"`
}

type DownStruct struct {
	Title string
	Url   string
	Type  string
}
type Dy struct {
	Url               string
	CId               string `bson:"c_id"`
	RId               string `bson:"r_id"`
	Title             string
	Alias             []string
	LongTitle         string `bson:"long_title"`
	Pic               string
	Director          []string
	Stars             []string
	Introduction      string
	DownStatus        int
	LastUpdated       string `bson:"last_updated"`
	UpdatedDate       string `bson:"updated_date"`
	Source            string
	UpdateTime        time.Time `bson:"update_time"`
	ProductionDate    string    `bson:"production_date"`
	PageDate          string    `bson:"page_date"`
	Rating            string
	DoubanUrl         string `bson:"douban_url"`
	DoubanId          string `bson:"douban_id"`
	Tags              []string
	Type              []string
	Year              string
	Area              string
	RunTime           string `bson:"run_time"`
	Language          string
	DownUrl           []DownStruct `bson:"down_url"`
	ProductionCompany string       `bson:"production_company"`
	Status            string
	ClickCount        int `bson:"click_count"`
	DownCount         int `bson:"down_count"`
}

func MInit() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/#environment-variable")
	}
	MClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)

	}
	defer func() {
		if err := MClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

}

func IsDyListOk(url string) string {
	coll := MClient.Database("dy").Collection("list")
	var result Default
	err := coll.FindOne(context.TODO(), bson.D{{"url", url}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return ""
	}
	if err != nil {
		panic(err)
	}

	return result.ID.Hex()
}

func SaveDy(dy *Dy) string {

	coll := MClient.Database("dy").Collection("list")
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{"url", dy.Url}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the name %s\n", dy.LongTitle)

		result, err := coll.InsertOne(context.TODO(), dy)

		if err != nil {
			panic(err)
		}
		fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
			return oid.String()
		} else {
			return ""
		}
	}
	if err != nil {
		panic(err)
	}
	return ""
}
