package db

import (
	"context"
	"crwal/model"
	"crwal/util"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client = util.GetMClient()

func IsDyListOk(url string) string {
	coll := client.Database("dy").Collection("list")
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
	coll := client.Database("dy").Collection("list")
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{"url", dy.Url}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		result, err := coll.InsertOne(context.TODO(), dy)
		if err != nil {
			panic(err)
		}
		fmt.Printf("新增: %s\n", dy.LongTitle)
		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
			return oid.String()
		} else {
			return ""
		}
	} else {
		fmt.Printf("已存在 %s\n", dy.LongTitle)
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
		fmt.Printf("新增下载信息: %s\n", down_info.Title)
		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
			return oid.String()
		} else {
			return ""
		}
	} else {

		filter := bson.D{{"url", down_info.Url}}
		update := bson.D{{"$set", down_info}}
		_, err := coll.UpdateMany(context.TODO(), filter, update)
		if err != nil {
			panic(err)
		}

		fmt.Printf("更新下载信息 %s\n", down_info.Title)
		return ""

	}
}
