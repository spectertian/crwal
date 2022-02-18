package db

import (
	"context"
	"crwal/model"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

func SaveAndUpdateFj(vod *model.WJVod) string {
	coll := client.Database("dy").Collection("vod_fj_list")
	var result model.VodIndexHas

	err := coll.FindOne(context.TODO(), bson.D{{"vod_id", vod.VodId}}).Decode(&result)
	if err == mongo.ErrNoDocuments {

		vod.UpdatedTime = time.Now()
		vod.CreatedTime = time.Now()
		results, err := coll.InsertOne(context.TODO(), vod)
		if err != nil {
			panic(err)
		}
		fmt.Println("新增wj信息", vod.VodName, time.Now().Format("2006-01-02 15:04:05"))

		if oid, ok := results.InsertedID.(primitive.ObjectID); ok {
			return oid.Hex()
		} else {
			return ""
		}
	} else {
		vod.UpdatedTime = time.Now()
		vod.CreatedTime = result.CreatedTime

		if result.VodDoubanId > 0 {
			vod.VodDoubanId = result.VodDoubanId
		}
		filter := bson.D{{"vod_id", vod.VodId}}
		update := bson.D{{"$set", vod}}
		_, err := coll.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			panic(err)
		}
		fmt.Println("更新wj信息", vod.VodName, time.Now().Format("2006-01-02 15:04:05"))
	}

	return result.ID.Hex()
}

func SaveVodImageFjById(id, pic_path string, tag string) {
	img_id := SaveVodImage(pic_path, tag)
	UpdateVodFjImagePic(id, img_id)
}

func UpdateVodFjImagePic(id string, img_id string) {
	coll := client.Database("dy").Collection("vod_fj_list")
	id_obj, _ := primitive.ObjectIDFromHex(id)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := coll.UpdateOne(
		ctx,
		bson.M{"_id": id_obj},
		bson.D{
			{"$set", bson.D{{"img_url", img_id}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(result)
	//fmt.Println(err)
	fmt.Printf("Updated pic %v Documents!\n", result.ModifiedCount)
}
