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

func GetFetchUrl(url string) {
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

	chans := make(chan model.Update, 200)

	fmt.Println("开始:", url, time.Now().Format("2006-01-02 15:04:05"))
	doc.Find("#vod .list-group-item").Each(func(i int, s *goquery.Selection) {
		hrefs, _ := s.Find("a").Attr("href")
		update := model.Update{}
		update.Type = "电影"
		update.CreatedTime = time.Now()
		update.UpdatedTime = time.Now()
		update.Date = strings.TrimSpace(s.Find("b").Text())
		update.Title = strings.TrimSpace(s.Find("a").Text())
		update.Url = domin + strings.TrimSpace(hrefs)
		Regexp := regexp.MustCompile(`([^/]*?)\.html`)
		params := Regexp.FindStringSubmatch(update.Url)
		update.CId = params[1]

		if update.Url == "" {
			fmt.Println("不存在url", update)
			return
		}

		hasId := db.IsHasUpdateByUrl(update.Url, update.Title)
		if hasId != "" {
			fmt.Println("已保存数据", update.Title)
			return
		}

		chans <- update
		fmt.Println("开始抓取", update.Url, update.Title, time.Now().Format("2006-01-02 15:04:05"))
	})

	doc.Find("#tv .list-group-item").Each(func(i int, s *goquery.Selection) {
		update := model.Update{}
		update.Type = "电视剧"
		update.CreatedTime = time.Now()
		update.UpdatedTime = time.Now()
		hrefs, _ := s.Find("a").Attr("href")
		update.Date = strings.TrimSpace(s.Find("b").Text())
		update.Title = strings.TrimSpace(s.Find("a").Text())

		update.Url = domin + strings.TrimSpace(hrefs)
		Regexp := regexp.MustCompile(`([^/]*?)\.html`)
		params := Regexp.FindStringSubmatch(update.Url)
		update.CId = params[1]

		if update.Url == "" {
			fmt.Println("不存在url", update)
			return
		}

		hasId := db.IsHasUpdateByUrl(update.Url, update.Title)
		if hasId != "" {
			fmt.Println("已保存数据", update.Title)
			return
		}

		chans <- update
		fmt.Println("开始抓取", update.Url, update.Title, time.Now().Format("2006-01-02 15:04:05"))
	})

	close(chans)

	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go CrwalInfo(chans, &wg)
	}
	wg.Wait()
}

func CrwalInfo(chans chan model.Update, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case update, ok := <-chans:
			if ok {
				info := db.GetDyInfo(update.Url)
				if info.LongTitle != update.Title {
					dy := model.Dy{}
					dy.Url = update.Url
					dy.LongTitle = update.Title
					dy.CId = update.CId
					dy_info := util.GetContentNewAll(&dy)
					info_id := db.SaveDy(&dy_info)

					db.SaveImageById(info_id, dy.Pic)

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

					update.InfoId = info_id
					update.ProductionDate = dy.ProductionDate
					db.SaveUpdate(&update)
				} else {
					update.InfoId = info.ID.Hex()
					update.ProductionDate = info.ProductionDate
					db.SaveUpdate(&update)
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
	url := "https://www.domp4.cc/custom/update.html"
	GetFetchUrl(url)
	ends := time.Now().Unix()

	fmt.Println("抓取结束", time.Now())
	fmt.Println("耗时", time.Now(), starts, ends, ends-starts)
}
