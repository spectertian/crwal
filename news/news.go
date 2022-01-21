package main

import (
	"crwal/db"
	"crwal/model"
	"crwal/util"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

var domin = "https://www.domp4.cc/"
var wg sync.WaitGroup

func GetFetchUrl(crawl_url string) {
	chans := make(chan model.NewsStruct, 2000)
	i := 1
	for {
		url := fmt.Sprintf(crawl_url, i)
		news := model.NewsStruct{}
		news.CreatedTime = time.Now()
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
		doc.Find("#list_dy ul li").Each(func(i int, s *goquery.Selection) {
			hrefs, _ := s.Find("a").Attr("href")
			news.Title = strings.TrimSpace(s.Find(" a").Text())
			news.Date = strings.TrimSpace(s.Find("span").Text())
			news.Url = domin + strings.TrimSpace(hrefs)
			Regexp := regexp.MustCompile(`([^/]*?)\.html`)
			params := Regexp.FindStringSubmatch(news.Url)
			news.CId = params[1]

			if news.Url == "" {
				fmt.Println("不存在url", news)
				return
			}

			news_id := db.IsHasNewsByUrl(news.Url)
			if news_id != "" {
				fmt.Println("已保存数据", news.Title)
				return
			}
			chans <- news
		})
		_, ok := doc.Find(".pagination li").Eq(5).Find("a").Attr("href")
		fmt.Println("第", i, "页")
		if ok == false {
			fmt.Println(crawl_url, "页面url便利完成")
			break
		}
		i++
	}

	close(chans)

	wg.Add(6)
	for i := 1; i <= 6; i++ {
		CrwalInfo(chans, &wg)
	}
	wg.Wait()

}

func CrwalInfo(chans chan model.NewsStruct, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case news, ok := <-chans:
			if ok {

				fmt.Println(news)
				info := db.GetDyInfo(news.Url)
				if info.LongTitle != news.Title {
					dy := model.Dy{}
					dy.Url = news.Url
					dy.LongTitle = news.Title
					dy.CId = news.CId
					dy_info := util.GetContentNewAll(&dy)
					info_id := db.SaveDy(&dy_info)

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

					news.InfoId = info_id
					db.SaveNews(&news)
				} else {
					news.InfoId = info.ID.Hex()
					db.SaveNews(&news)
				}

			} else {
				fmt.Println("退出", time.Now().Format("2006-01-02 15:04:05"))
				goto forEnd
			}

		}
	}
forEnd:
	return
}

func main() {
	fmt.Println("抓取开始", time.Now())
	starts := time.Now().Unix()
	url := "https://www.domp4.cc/list/99-%v.html"
	GetFetchUrl(url)
	ends := time.Now().Unix()

	fmt.Println("抓取结束", time.Now())
	fmt.Println("耗时", time.Now(), starts, ends, ends-starts)
}
