package main

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var domin = "https://www.domp4.cc/"

type Default struct {
	ID primitive.ObjectID `bson:"_id" json:"id,omitempty"`
}

func IsBookOk(url string) string {
	//panic(2)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)

	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("dzs").Collection("book")
	var result Default
	err = coll.FindOne(context.TODO(), bson.D{{"url", url}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return ""
	}
	if err != nil {
		panic(err)
	}

	return result.ID.Hex()
}

func IsChapterOk(url string) string {
	//panic(2)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)

	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("dzs").Collection("chapter")
	var result Default
	err = coll.FindOne(context.TODO(), bson.D{{"url", url}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return ""
	}
	if err != nil {
		panic(err)
	}

	return result.ID.Hex()
}

type Dy struct {
	CId               string
	RId               string
	Introduction      string
	Title             string
	Alias             string
	Director          []string
	Stars             []string
	LastUpdated       string
	UpdatedDate       string
	Source            string
	UpdateTime        time.Time
	ProductionDate    string
	PageDate          string
	Rating            string
	Tags              []string
	Type              []string
	Url               string
	Year              string
	Area              string
	RunTime           string
	Language          string
	DownUrl           []string
	ProductionCompany string
	Status            string
	ClickCount        int `bson:"click_count"`
	DownCount         int `bson:"down_count"`
}

func main() {

	url := "https://www.domp4.cc/list/99-1.html"
	dy := &Dy{}
	dy.UpdateTime = time.Now()

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	res.Close = true
	res.Header.Add("Connection", "close")
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan Dy, 5)

	doc.Find("#list_dy li").Each(func(i int, s *goquery.Selection) {
		hrefs, _ := s.Find("a").Attr("href")
		dy.Alias = strings.TrimSpace(s.Find("a").Text())
		dy.PageDate = strings.TrimSpace(s.Find("span").Text())
		dy.Url = domin + strings.TrimSpace(hrefs)
		if dy.Url == "" {
			fmt.Println("不存在url", dy)
			return
		}
		ch <- *dy

		fmt.Println("存在url", dy)

	})

	close(ch)

	go func() {
		for {
			select {

			case dy = <-ch:
				getContent(dy)
				break
			default:
				fmt.Printf("no communication\n")

			}

		}
	}()

	panic("aa")

}

func getContent(dy *Dy) Dy {
	res, err := http.Get(dy.Url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	contents := doc.Find("#content").Text()

	fmt.Println(contents)
	panic("cc")
	return *dy
}

func GetChapterInfo(ch chan Chapter, chapter *Chapter) {
	ch <- *chapter

}

type Chapter struct {
	Title      string
	Number     string
	Url        string
	Content    string
	UpdateTime time.Time
	BookId     string
	ClickCount int `bson:"click_count"`
}

func SaveChapter(chapters *Chapter) string {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)

	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("dzs").Collection("chapter")
	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"url", chapters.Url}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the name %s\n", chapters.Title)

		result, err := coll.InsertOne(context.TODO(), chapters)

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

type Book struct {
	Name         string
	Author       string
	Description  string
	LastUpdated  string
	Source       string
	UpdateTime   time.Time
	Tags         []string
	Url          string
	StatusString string
	Status       int `bson:"status"`
	CId          int `bson:"c_id"`
	ClickCount   int `bson:"click_count"`
}

func SaveBook(books *Book) string {

	//panic(2)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)

	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("dzs").Collection("book")
	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"name", books.Name}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the name %s\n", books.Name)

		result, err := coll.InsertOne(context.TODO(), books)

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
