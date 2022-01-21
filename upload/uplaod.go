package main

import (
	"context"
	"crwal/db"
	"crwal/model"
	"crwal/util"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func main() {
	//file_path := "https://img.domp4.cc/vod/0/61e786484ebb2.jpg"
	//file_path = "https://demibaguette.com/wp-content/uploads/2019/01/erik-witsoe-647316-unsplash-1.jpg"

	var client = util.GetMClient()

	coll := client.Database("dy").Collection("list")
	cur, err := coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Println(err)
	}

	for cur.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem = model.FDyStruct{}
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(elem)
		fmt.Println(elem.Pic)
		pic_path := elem.Pic
		db.SaveImageById(elem.ID.Hex(), pic_path)
	}

}
