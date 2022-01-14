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

func IsHasNewsByUrl(url string) string {
	coll := client.Database("dy").Collection("news")
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

func IsHasIndexByUrl(url string, types string) string {
	coll := client.Database("dy").Collection("index_list")
	var result model.IndexHas
	err := coll.FindOne(context.TODO(), bson.D{{"url", url}, {"type", types}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return ""
	}
	if err != nil {
		panic(err)
	}

	return result.ID.Hex()
}

func IsDownOk(url string) int {
	coll := client.Database("dy").Collection("lists")
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

	coll := client.Database("dy").Collection("lists")
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

func IsHasUpdateByUrl(url string) string {
	coll := client.Database("dy").Collection("update")
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
		fmt.Println("已存在", update.Title, time.Now().Format("2006-01-02 15:04:05"))
		return ""
	}
}

func SaveNews(update *model.NewsStruct) string {
	coll := client.Database("dy").Collection("news")
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
		fmt.Println("已存在", update.Title, time.Now().Format("2006-01-02 15:04:05"))
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
		return ""
	}
}

func SaveDy(dy *model.Dy) string {
	coll := client.Database("dy").Collection("lists")
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
		return result.ID.Hex()
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
