package main

import (
	"crwal/db"
	"crwal/model"
	"crwal/util"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var domin = "https://www.domp4.cc/"
var wg sync.WaitGroup

func GetFetchUrl(crawl_url string, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 1
	for {
		url := fmt.Sprintf(crawl_url, i)
		dy := model.Dy{}
		dy.UpdatedTime = time.Now()
		dy.CreatedTime = time.Now()
		dy.Status = 1
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

		fmt.Println("开始:", url, time.Now().Format("2006-01-02 15:04:05"))
		doc.Find("#list_all ul li").Each(func(i int, s *goquery.Selection) {
			hrefs, _ := s.Find(".text_info h2 a").Attr("href")
			dy.LongTitle = strings.TrimSpace(s.Find(".text_info h2 a").Text())
			dy.PageDate = strings.TrimSpace(s.Find(".text_info .update_time").Text())
			dy.Url = domin + strings.TrimSpace(hrefs)
			dy.DownStatus = 1
			Regexp := regexp.MustCompile(`([^/]*?)\.html`)
			params := Regexp.FindStringSubmatch(dy.Url)
			dy.CId = params[1]

			if dy.Url == "" {
				fmt.Println("不存在url", dy)
				return
			}

			dy_info := db.GetDyInfo(dy.Url)
			if dy_info.DownStatus == 1 {
				fmt.Println("已保存数据", dy.LongTitle)
				return
			}

			if dy_info.DownStatus == 0 {
				fmt.Println("开始抓取  更新数据", dy.Url, dy.LongTitle, time.Now().Format("2006-01-02 15:04:05"))
			} else {
				fmt.Println("开始抓取", dy.Url, dy.LongTitle, time.Now().Format("2006-01-02 15:04:05"))
			}

			CrwaInfo(&dy)
		})

		_, ok := doc.Find(".pagination li").Eq(5).Find("a").Attr("href")
		if ok == false {
			fmt.Println(crawl_url, "抓取完成")
			break
		}
		i++

	}
}

func CrwaInfo(dy *model.Dy) {
	dy_info := util.GetContentNewAll(dy)
	dy_id := db.SaveDy(&dy_info)

	db.SaveImageById(dy_id, dy.Pic)

	down_info := model.DownInfoStruct{}
	down_info.DownUrl = dy.DownUrl
	down_info.Url = dy.Url
	down_info.Title = dy.Title
	down_info.LongTitle = dy.LongTitle
	down_info.DownStatus = dy.DownStatus
	down_info.CId = dy.CId
	down_info.Type = dy.Type
	down_info.UpdatedTime = dy.UpdatedTime
	down_info.CreatedTime = dy.CreatedTime
	db.SaveAndUpdateDownInfo(&down_info)
}
func main() {
	fmt.Println("抓取开始", time.Now())
	starts := time.Now().Unix()
	wg.Add(10)
	list := []string{
		"https://www.domp4.cc/list/1-%v.html",
		"https://www.domp4.cc/list/2-%v.html",
		"https://www.domp4.cc/list/3-%v.html",
		"https://www.domp4.cc/list/4-%v.html",
		"https://www.domp4.cc/list/5-%v.html",
		"https://www.domp4.cc/list/6-%v.html",
		"https://www.domp4.cc/list/7-%v.html",
		"https://www.domp4.cc/list/8-%v.html",
		"https://www.domp4.cc/list/9-%v.html",
		"https://www.domp4.cc/list/10-%v.html",
	}

	for _, v := range list {
		go GetFetchUrl(v, &wg)
	}

	wg.Wait()

	ends := time.Now().Unix()

	fmt.Println("抓取结束", time.Now())
	fmt.Println("耗时", time.Now(), starts, ends, ends-starts)
}
