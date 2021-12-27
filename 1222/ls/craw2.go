package main

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func DzsScrape() {
	// Request the HTML page.
	url := "http://www.tingshuge.com/book/30937.html"

	books := &Book{}
	books.UpdateTime = time.Now()
	book_id := IsBookOk(url)

	if book_id == "" {
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

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("抓去简介", url)
		doc.Find("meta").Each(func(i int, s *goquery.Selection) {
			op, _ := s.Attr("property")
			res_title, _ := s.Attr("content")
			if op == "og:novel:author" {
				author, _ := iconv.ConvertString(res_title, "gbk", "utf-8")
				fmt.Println("作者：", author)
				author = strings.TrimSpace(author)
				books.Author = author
			}

			if op == "og:novel:status" {
				status, _ := iconv.ConvertString(res_title, "gbk", "utf-8")
				fmt.Println("状态：", status)
				status = strings.TrimSpace(status)

				books.StatusString = status

			}

			if op == "og:novel:category" {
				category, _ := iconv.ConvertString(res_title, "gbk", "utf-8")
				fmt.Println("分类：", category)
				category = strings.TrimSpace(category)

				books.Tags = []string{category}

			}
			if op == "og:url" {
				source_url, _ := iconv.ConvertString(res_title, "gbk", "utf-8")
				fmt.Println("地址：", source_url)
				books.Url = source_url

			}

			if op == "og:novel:book_name" {
				book_name, _ := iconv.ConvertString(res_title, "gbk", "utf-8")
				fmt.Println("书名：", book_name)
				book_name = strings.TrimSpace(book_name)

				books.Name = book_name

			}

			if op == "og:novel:update_time" {
				update_time, _ := iconv.ConvertString(res_title, "gbk", "utf-8")
				fmt.Println("最后更新日期：", update_time)
				update_time = strings.TrimSpace(update_time)
				books.LastUpdated = update_time

			}
			if op == "og:description" {
				description, _ := iconv.ConvertString(res_title, "gbk", "utf-8")
				fmt.Println("描述：", description)
				description = strings.TrimSpace(description)

				books.Description = description
			}

		})

		SaveBook(books)
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	res.Header.Add("Connection", "close")
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan Chapter, 5)
	// Find the review items
	doc.Find(".listmain dd").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find("a").Text()
		href, _ := s.Find("a").Attr("href")
		re_title, _ := iconv.ConvertString(title, "gbk", "utf-8")
		fmt.Println("标题：", re_title)
		re_title = strings.TrimSpace(re_title)
		chapters := &Chapter{}
		chapters.Url = "http://www.tingshuge.com" + href

		if IsChapterOk(chapters.Url) != "" {
			fmt.Println(re_title, "已抓去")
			return
		}
		fmt.Println(re_title, "抓取中...")

		chapters.BookId = book_id
		reg1 := regexp.MustCompile(`([^\d]*)([0-9]*)`)
		if reg1 == nil { //解释失败，返回nil
			fmt.Println("regexp err")
			return
		}

		//根据规则提取关键信息
		result1 := reg1.FindAllStringSubmatch(re_title, -1)
		fmt.Println(result1)
		chapters.Number = result1[0][2]
		chapters.Title = "第" + chapters.Number + "章"

		fmt.Println(chapters)
		if href != "" {
			fmt.Println("进入队列 result1 = ", result1)
			fmt.Println("进入队列 href = ", href)
			ch <- *chapters
		}
	})
	chapters := &Chapter{}
	chapters.ClickCount = 100
	ch <- *chapters
	close(ch)

	go func() {
		for {
			ss := <-ch
			chapters := &ss
			if chapters.ClickCount == 100 {
				fmt.Println("队列退出： ", chapters)
				break
			}
			contents := getContent(chapters)
			if contents != "" {
				chapters.Content = contents
				chapters.UpdateTime = time.Now()
				fmt.Println(chapters)
				fmt.Println("消费队列： ", chapters)
				SaveChapter(chapters)
			}
		}
	}()

}

func getContent(chapters *Chapter) string {
	res, err := http.Get(chapters.Url)
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
	re_contents, _ := iconv.ConvertString(contents, "gbk", "utf-8")
	contentsAll := strings.ReplaceAll(re_contents, "聽", "")
	contentsAllNew := strings.ReplaceAll(contentsAll, "\n\n", "\n")
	return contentsAllNew
}

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

func main() {
	DzsScrape()
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
