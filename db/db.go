package db

import (
	"bytes"
	"context"
	"crwal/model"
	"crwal/util"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"
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

func IsHasNewsByUrl(url, title string) string {
	coll := client.Database("dy").Collection("news")
	var result model.UpdateHas
	err := coll.FindOne(context.TODO(), bson.D{{"url", url}, {"title", title}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return ""
	}
	if err != nil {
		panic(err)
	}

	return result.ID.Hex()
}

func IsHasIndexByUrl(url string, types string, title string) string {
	coll := client.Database("dy").Collection("index_list")
	var result model.IndexHas
	err := coll.FindOne(context.TODO(), bson.D{{"url", url}, {"type", types}, {"title", title}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return ""
	}
	if err != nil {
		panic(err)
	}

	return result.ID.Hex()
}

func IsDownOk(url string) int {
	coll := client.Database("dy").Collection("list")
	var result model.Default
	err := coll.FindOne(context.TODO(), bson.D{{"url", url}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return 3
	}
	if err != nil {
		panic(err)
	}

	return result.DownStatus
}

func GetDyInfo(url string) model.Default {

	coll := client.Database("dy").Collection("list")
	var result model.Default
	err := coll.FindOne(context.TODO(), bson.D{{"url", url}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		ss := model.Default{}
		ss.DownStatus = 8
		return ss
	}
	if err != nil {
		panic(err)
	}

	return result
}

func IsHasUpdateByUrl(url, title string) string {
	coll := client.Database("dy").Collection("update")
	var result model.UpdateHas
	err := coll.FindOne(context.TODO(), bson.D{{"url", url}, {"title", title}}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		return ""
	}
	if err != nil {
		panic(err)
	}

	return result.ID.Hex()
}

func IsHasTopicByUrl(url string) string {
	coll := client.Database("dy").Collection("topic")
	var result model.UpdateHas
	err := coll.FindOne(context.TODO(), bson.D{{"url", url}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return ""
	}
	if err != nil {
		panic(err)
	}

	return result.ID.Hex()
}

func SaveUpdate(update *model.Update) string {
	coll := client.Database("dy").Collection("update")
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{"url", update.Url}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		result, err := coll.InsertOne(context.TODO(), update)
		if err != nil {
			panic(err)
		}
		fmt.Println("新增", update.Title, time.Now().Format("2006-01-02 15:04:05"))
		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
			return oid.String()
		} else {
			return ""
		}
	} else {

		upS := model.UpUpdateStruct{}
		upS.Title = update.Title
		upS.UpdatedTime = time.Now()
		upS.ProductionDate = update.ProductionDate
		upS.Date = update.Date
		filter := bson.D{{"url", update.Url}}
		updateS := bson.D{{"$set", upS}}
		_, err := coll.UpdateMany(context.TODO(), filter, updateS)
		if err != nil {
			panic(err)
		}
		fmt.Println("更新下载信息", update.Title, time.Now().Format("2006-01-02 15:04:05"))

		fmt.Println("已存在", update.Title, time.Now().Format("2006-01-02 15:04:05"))
		return ""
	}
}

func SaveNews(news *model.NewsStruct) string {
	coll := client.Database("dy").Collection("news")
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{"url", news.Url}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		result, err := coll.InsertOne(context.TODO(), news)
		if err != nil {
			panic(err)
		}
		fmt.Println("新增", news.Title, time.Now().Format("2006-01-02 15:04:05"))
		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
			return oid.String()
		} else {
			return ""
		}
	} else {

		upS := model.UpNewsStruct{}
		upS.Title = news.Title
		upS.UpdatedTime = time.Now()
		upS.ProductionDate = news.ProductionDate
		upS.Date = news.Date
		filter := bson.D{{"url", news.Url}}
		update := bson.D{{"$set", upS}}
		_, err := coll.UpdateMany(context.TODO(), filter, update)
		if err != nil {
			panic(err)
		}
		fmt.Println("更新下载信息", news.Title, time.Now().Format("2006-01-02 15:04:05"))

		fmt.Println("已存在", news.Title, time.Now().Format("2006-01-02 15:04:05"))
		return ""
	}
}

func SaveTopic(topic *model.TopicStruct) string {
	coll := client.Database("dy").Collection("topic")
	var result model.DefaultTopicStruct
	err := coll.FindOne(context.TODO(), bson.D{{"url", topic.Url}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		result, err := coll.InsertOne(context.TODO(), topic)
		if err != nil {
			panic(err)
		}
		fmt.Println("新增", topic.Title, time.Now().Format("2006-01-02 15:04:05"))
		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
			return oid.Hex()
		} else {
			return ""
		}
	} else {
		fmt.Println("已存在", topic.Title, time.Now().Format("2006-01-02 15:04:05"))
		return result.ID.Hex()
	}
}

func SaveTopicList(topic_list *model.TopicListStruct) string {
	coll := client.Database("dy").Collection("topic_list")
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{"url", topic_list.Url}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		result, err := coll.InsertOne(context.TODO(), topic_list)
		if err != nil {
			panic(err)
		}
		fmt.Println("新增", topic_list.Title, time.Now().Format("2006-01-02 15:04:05"))
		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
			return oid.Hex()
		} else {
			return ""
		}
	} else {
		fmt.Println("已存在", topic_list.Title, time.Now().Format("2006-01-02 15:04:05"))
		return ""
	}
}

func SaveIndexList(index_list *model.IndexListStruct) string {
	coll := client.Database("dy").Collection("index_list")
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{"url", index_list.Url}, {"type", index_list.Type}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		result, err := coll.InsertOne(context.TODO(), index_list)
		if err != nil {
			panic(err)
		}
		fmt.Println("新增", index_list.Title, time.Now().Format("2006-01-02 15:04:05"))
		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
			return oid.Hex()
		} else {
			return ""
		}
	} else {
		fmt.Println("已存在", index_list.Title, time.Now().Format("2006-01-02 15:04:05"))

		upS := model.UpdateIndexListStruct{}
		upS.Title = index_list.Title
		upS.UpdatedTime = time.Now()
		upS.ProductionDate = index_list.ProductionDate
		upS.Date = index_list.Date
		filter := bson.D{{"url", index_list.Url}, {"type", index_list.Type}}
		update := bson.D{{"$set", upS}}
		_, err := coll.UpdateMany(context.TODO(), filter, update)
		if err != nil {
			panic(err)
		}
		fmt.Println("更新下载信息", index_list.Title, time.Now().Format("2006-01-02 15:04:05"))

		return ""
	}
}

func SaveDy(dy *model.Dy) string {
	coll := client.Database("dy").Collection("list")
	var result model.Default
	err := coll.FindOne(context.TODO(), bson.D{{"url", dy.Url}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		result, err := coll.InsertOne(context.TODO(), dy)
		if err != nil {
			panic(err)
		}
		fmt.Println("新增", dy.LongTitle, time.Now().Format("2006-01-02 15:04:05"))
		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
			return oid.Hex()
		} else {
			return ""
		}
	} else {
		fmt.Println("已存在", dy.LongTitle, time.Now().Format("2006-01-02 15:04:05"))
		dy.UpdatedTime = time.Now()

		upS := model.UpdateDyStruct{}
		upS.ProductionDate = dy.ProductionDate
		upS.LongTitle = dy.LongTitle
		upS.DownUrl = dy.DownUrl
		upS.DownStatus = dy.DownStatus
		upS.UpdatedTime = dy.UpdatedTime
		upS.Rating = dy.Rating
		upS.DoubanUrl = dy.DoubanUrl

		filter := bson.D{{"url", dy.Url}}
		update := bson.D{{"$set", upS}}
		_, err := coll.UpdateMany(context.TODO(), filter, update)
		if err != nil {
			panic(err)
		}
		fmt.Println("更新下载信息", dy.Title, time.Now().Format("2006-01-02 15:04:05"))
		return result.ID.Hex()
	}
}

func UpdateDy(id string, dy *model.UpDyStruct) {
	coll := client.Database("dy").Collection("list")
	id_obj, _ := primitive.ObjectIDFromHex(id)
	fmt.Println(bson.M{"_id": id_obj})
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := coll.UpdateOne(
		ctx,
		bson.M{"_id": id_obj},
		bson.D{
			{"$set", dy},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated pic %v Documents!\n", result.ModifiedCount)
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

func UploadFile(body *[]byte, filename, contentType string) string {
	bucket, err := gridfs.NewBucket(
		client.Database("image"),
	)

	if err != nil {
		panic(err)
	}

	uploadOpts := options.GridFSUpload().
		SetMetadata(bson.D{{"content-type", contentType}})

	fileID, err := bucket.UploadFromStream(
		filename,
		bytes.NewBuffer(*body),
		uploadOpts)
	if err != nil {
		log.Fatal(err)
	}

	return fileID.Hex()
}

func IsHasFile(fileName string) string {
	db := client.Database("image")
	fsFiles := db.Collection("fs.files")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var results model.IndexHas
	err := fsFiles.FindOne(ctx, bson.M{"filename": fileName}).Decode(&results)

	// you can print out the results
	if err == mongo.ErrNoDocuments {
		return ""
	}
	if err != nil {
		panic(err)
	}
	return results.ID.Hex()

}

func SaveImage(path_url string) string {
	if path_url == "" {
		return ""
	}

	file_name := path.Base(path_url)
	file_id := IsHasFile(file_name)
	if file_id != "" {
		fmt.Println("has_file:", file_id)
		return file_id
	}
	resp, _ := http.Get(path_url)
	body, _ := ioutil.ReadAll(resp.Body)
	contentType := http.DetectContentType(body)
	fmt.Println(contentType)
	insert_id := UploadFile(&body, file_name, contentType)
	fmt.Println("插入图片", insert_id)
	return insert_id
}

func SaveLocalImage(path_url string) string {
	if path_url == "" {
		return ""
	}

	file_name := path.Base(path_url)
	file_id := IsHasFile(file_name)
	if file_id != "" {
		fmt.Println("has_file:", file_id)
		return file_id
	}

	resp, err := os.Open("D:/gopath/src/golang_development_notes/example/log.txt")
	if err != nil {
		panic(err)
	}
	defer resp.Close()
	body, _ := ioutil.ReadAll(resp)
	contentType := http.DetectContentType(body)
	fmt.Println(contentType)
	insert_id := UploadFile(&body, file_name, contentType)
	fmt.Println("插入图片", insert_id)
	return insert_id
}

func UpdateImagePic(id string, img_id string) {
	coll := client.Database("dy").Collection("list")
	id_obj, _ := primitive.ObjectIDFromHex(id)
	fmt.Println(bson.M{"_id": id_obj})
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

func SaveImageById(id, pic_path string) {
	img_id := SaveImage(pic_path)
	UpdateImagePic(id, img_id)
}

func GetDyInfoById(id string) model.Default {
	coll := client.Database("dy").Collection("list")
	id_obj, _ := primitive.ObjectIDFromHex(id)

	var result model.Default
	err := coll.FindOne(context.TODO(), bson.M{"_id": id_obj}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		ss := model.Default{}
		ss.DownStatus = 8
		return ss
	}
	if err != nil {
		panic(err)
	}

	return result

}
