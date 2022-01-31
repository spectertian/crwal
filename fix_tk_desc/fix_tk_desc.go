package main

import (
	"context"
	"crwal/db"
	"crwal/model"
	"crwal/util"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	//file_path := "https://img.domp4.cc/vod/0/61e786484ebb2.jpg"
	//file_path = "https://demibaguette.com/wp-content/uploads/2019/01/erik-witsoe-647316-unsplash-1.jpg"

	var client = util.GetMClient()

	coll := client.Database("dy").Collection("tk_list")
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
		all := GetIntroductionByUrl(elem.Url, 0)
		fmt.Println(elem.ID.Hex(), all.Title, all.Introduction)
		tk_up := model.TKUpdateIntroductionStruct{all.Introduction}
		db.UpdateTkDy(elem.ID.Hex(), &tk_up)
	}

}

func GetIntroductionByUrl(url string, c_count int) model.TKStruct {
forStart:
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	res.Close = true
	res.Header.Add("Connection", "close")
	defer res.Body.Close()
	if res.StatusCode != 200 {
		c_count = c_count + 1
		time.Sleep(time.Second * 1)
		fmt.Println("抓取次数：", c_count, "----", url)
		goto forStart
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	tk := model.TKStruct{}
	tk.Title = strings.TrimSpace(doc.Find(".vodh h2").Text())
	tk.Introduction = strings.TrimSpace(doc.Find(".vodplayinfo").Eq(0).Text())

	return tk
}
