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

var domin = "https://www.domp4.cc"
var wg sync.WaitGroup

func GetFetchUrl(url string) {
	chans := make(chan model.IndexListStruct, 2000)
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
	types := ""

	doc.Find(".index_update .list-group").Each(func(i int, s *goquery.Selection) {
		types = strings.TrimSpace(s.Find("#heading-text").Text())
		s.Find("ul li").Each(func(k int, ks *goquery.Selection) {
			index_list := model.IndexListStruct{}
			index_list.CreatedTime = time.Now()
			index_list.Sort = k
			index_list.Type = types
			hrefs, _ := ks.Find("a").Attr("href")
			if hrefs == "" {
				return
			}
			index_list.Title = strings.TrimSpace(ks.Find(" a").Text())
			index_list.Date = strings.TrimSpace(ks.Find("b").Text())
			index_list.Url = domin + strings.TrimSpace(hrefs)
			Regexp := regexp.MustCompile(`([^/]*?)\.html`)
			params := Regexp.FindStringSubmatch(index_list.Url)
			index_list.CId = params[1]
			if index_list.Url == "" {
				fmt.Println("不存在url", index_list)
				return
			}
			index_list_id := db.IsHasIndexByUrl(index_list.Url, index_list.Type)
			if index_list_id != "" {
				fmt.Println("已保存数据", index_list)
				return
			}

			fmt.Println(index_list)
			chans <- index_list
		})
	})

	doc.Find(".index_hot .list-group ").Each(func(i int, s *goquery.Selection) {
		types = strings.TrimSpace(s.Find("#heading-text").Text())
		s.Find("ul li").Each(func(k int, ks *goquery.Selection) {
			index_list := model.IndexListStruct{}
			index_list.CreatedTime = time.Now()
			index_list.Sort = k
			index_list.Type = types
			hrefs, _ := ks.Find("a").Attr("href")
			if hrefs == "" {
				return
			}
			index_list.Title = strings.TrimSpace(ks.Find(" a").Text())
			index_list.Date = strings.TrimSpace(ks.Find("b").Text())
			index_list.Url = domin + strings.TrimSpace(hrefs)
			Regexp := regexp.MustCompile(`([^/]*?)\.html`)
			params := Regexp.FindStringSubmatch(index_list.Url)
			index_list.CId = params[1]
			if index_list.Url == "" {
				fmt.Println("不存在url", index_list)
				return
			}
			index_list_id := db.IsHasIndexByUrl(index_list.Url, index_list.Type)
			if index_list_id != "" {
				fmt.Println("已保存数据", index_list)
				return
			}

			fmt.Println(index_list)
			chans <- index_list
		})
	})

	doc.Find(".index_today ul").Find("li").Each(func(i int, s *goquery.Selection) {
		index_list := model.IndexListStruct{}
		index_list.CreatedTime = time.Now()
		index_list.Type = "today_recommend"
		index_list.Sort = i
		hrefs, _ := s.Find("a").Attr("href")
		index_list.Title = strings.TrimSpace(s.Find(" a").Text())
		index_list.Date = strings.TrimSpace(s.Find("b").Text())
		index_list.Url = domin + strings.TrimSpace(hrefs)
		Regexp := regexp.MustCompile(`([^/]*?)\.html`)
		params := Regexp.FindStringSubmatch(index_list.Url)
		index_list.CId = params[1]
		if index_list.Url == "" {
			fmt.Println("今日推荐 不存在url", index_list)
			return
		}
		index_list_id := db.IsHasIndexByUrl(index_list.Url, index_list.Type)
		fmt.Println(index_list.Url)
		if index_list_id != "" {
			fmt.Println("已保存数据", index_list)
			return
		}

		fmt.Println("今日推荐", index_list)
		chans <- index_list
	})

	close(chans)
	wg.Add(6)
	for i := 1; i <= 6; i++ {
		CrwalInfo(chans, &wg)
	}
	wg.Wait()

}

func CrwalInfo(chans chan model.IndexListStruct, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case index_list, ok := <-chans:
			if ok {
				fmt.Println(index_list)
				info := db.GetDyInfo(index_list.Url)
				if info.LongTitle != index_list.Title {
					dy := model.Dy{}
					dy.Url = index_list.Url
					dy.LongTitle = index_list.Title
					dy.CId = index_list.CId
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

					index_list.InfoId = info_id
					db.SaveIndexList(&index_list)
				} else {
					index_list.InfoId = info.ID.Hex()
					db.SaveIndexList(&index_list)
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
	url := "https://www.domp4.cc"
	GetFetchUrl(url)
	ends := time.Now().Unix()

	fmt.Println("抓取结束", time.Now())
	fmt.Println("耗时", time.Now(), starts, ends, ends-starts)
}
