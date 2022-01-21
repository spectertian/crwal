package main

import (
	"context"
	"crwal/db"
	"crwal/model"
	"crwal/util"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

func main() {

	fmt.Println("begin", time.Now().Format("2006-01-02 15:04:05"))

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
		dy := model.Dy{}
		dy.Url = elem.Url
		all := util.GetContentNewAll(&dy)

		re_save := model.UpDyStruct{}
		re_save.UpdatedTime = time.Now()
		re_save.Language = all.Language
		re_save.LongTitle = all.LongTitle
		re_save.UpdatedDate = all.UpdatedDate
		re_save.Year = all.Year
		re_save.Area = all.Area
		re_save.RunTime = all.RunTime
		re_save.RunTime = all.RunTime

		fmt.Println(elem.Title, dy.LongTitle, elem.Url)
		db.UpdateDy(elem.ID.Hex(), &re_save)
		//os.Exit(1)
	}

	fmt.Println("end", time.Now().Format("2006-01-02 15:04:05"))

}
