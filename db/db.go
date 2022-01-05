package db

import (
	"context"
	"crwal/model"
	"crwal/util"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var client = util.GetMClient()

func IsDyListOk(url string) string {
	coll := client.Database("dy").Collection("lists")
	var result model.Default
	err := coll.FindOne(context.TODO(), bson.D{{"url", url}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return ""
	}
	if err != nil {
		panic(err)
	}

	return result.ID.Hex()
}

func SaveDy(dy *model.Dy) string {
	coll := client.Database("dy").Collection("lists")
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{"url", dy.Url}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		result, err := coll.InsertOne(context.TODO(), dy)
		if err != nil {
			panic(err)
		}
		fmt.Println("新增", dy.LongTitle, time.Now().Format("2006-01-02 15:04:05"))
		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
			return oid.String()
		} else {
			return ""
		}
	} else {
		fmt.Println("已存在", dy.LongTitle, time.Now().Format("2006-01-02 15:04:05"))
		return ""
	}

}

func SaveAndUpdateDownInfo(down_info *model.DownInfoStruct) string {
	coll := client.Database("dy").Collection("down_info")
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{"url", down_info.Url}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		result, err := coll.InsertOne(context.TODO(), down_info)
		if err != nil {
			panic(err)
		}
		fmt.Println("新增下载信息", down_info.Title, time.Now().Format("2006-01-02 15:04:05"))

		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
			return oid.String()
		} else {
			return ""
		}
	} else {
		down_info.UpdatedTime = time.Now()
		filter := bson.D{{"url", down_info.Url}}
		update := bson.D{{"$set", down_info}}
		_, err := coll.UpdateMany(context.TODO(), filter, update)
		if err != nil {
			panic(err)
		}
		fmt.Println("更新下载信息", down_info.Title, time.Now().Format("2006-01-02 15:04:05"))
		return ""
	}
}
