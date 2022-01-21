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

type ChanTopStruct struct {
	Top  model.TopicStruct
	List model.TopicListStruct
}

func GetFetchUrl(crawl_url string) {
	chans := make(chan ChanTopStruct, 2000)
	i := 1
	skip := 0
	for {
		url := fmt.Sprintf(crawl_url, i)

		if db.IsHasTopicByUrl(url) != "" {
			i++
			fmt.Println("已抓取过跳过", url)
			continue
		}
		topic := model.TopicStruct{}
		topic.NId = i
		topic.Url = url
		topic.CreatedTime = time.Now()
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

		topic.Title = strings.TrimSpace(doc.Find(".special h1").Text())
		if skip > 1 {
			fmt.Println("已分析完成")
			break
		}
		if topic.Title == "" {
			skip++
			i++
			fmt.Println("内容为空跳过")
			continue
		}
		topic.FilmNum = strings.TrimSpace(doc.Find(".tips .film_nums").Text())

		Regexp := regexp.MustCompile(`(\d*)部`)
		params := Regexp.FindStringSubmatch(topic.FilmNum)
		topic.FilmNum = params[1]
		topic.Date = strings.TrimSpace(strings.Split(doc.Find(".tips .update_time").Text(), "：")[1])
		topic.Content = strings.TrimSpace(doc.Find(".tips .special_content").Text())

		doc.Find("#list_all ul li").Each(func(k int, s *goquery.Selection) {
			hrefs, _ := s.Find("h2 a").Attr("href")
			topic_list := model.TopicListStruct{}
			topic_list.NId = i
			topic_list.CreatedTime = time.Now()
			topic_list.Title = strings.TrimSpace(s.Find("h2 a").Text())
			topic_list.Url = domin + strings.TrimSpace(hrefs)
			Regexp := regexp.MustCompile(`([^/]*?)\.html`)
			params := Regexp.FindStringSubmatch(topic_list.Url)
			topic_list.CId = params[1]

			if topic_list.Url == "" {
				fmt.Println("不存在url", topic_list)
				return
			}

			chans <- ChanTopStruct{topic, topic_list}
		})
		i++
	}

	close(chans)

	wg.Add(6)
	for i := 1; i <= 6; i++ {
		CrwalInfo(chans, &wg)
	}
	wg.Wait()
}

func CrwalInfo(chans chan ChanTopStruct, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case top, ok := <-chans:
			if ok {
				topic := top.Top
				topic_list := top.List
				topic_id := db.SaveTopic(&topic)
				info := db.GetDyInfo(topic.Url)
				if info.LongTitle != topic_list.Title {
					dy := model.Dy{}
					dy.Url = topic_list.Url
					dy.LongTitle = topic_list.Title
					dy.CId = topic_list.CId
					dy_info := util.GetContentNewAll(&dy)
					if dy_info.Status == 0 {
						fmt.Println("重新抓取一次", topic_list)
						dy_info = util.GetContentNewAll(&dy)
					}
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

					topic_list.InfoId = info_id
					topic_list.Director = dy_info.Director
					topic_list.Stars = dy_info.Stars
					topic_list.Area = dy_info.Area
					topic_list.Rating = dy_info.Rating
					topic_list.TopicId = topic_id
					topic_list.Pic = dy_info.Pic
					topic_list.Introduction = dy_info.Introduction
					db.SaveTopicList(&topic_list)
				} else {
					topic_list.InfoId = info.ID.Hex()
					topic_list.Director = info.Director
					topic_list.Stars = info.Stars
					topic_list.Area = info.Area
					topic_list.Rating = info.Rating
					topic_list.TopicId = topic_id
					topic_list.Pic = info.Pic
					topic_list.Introduction = info.Introduction
					db.SaveTopicList(&topic_list)
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
	url := "https://www.domp4.cc/special/%v.html"
	GetFetchUrl(url)
	ends := time.Now().Unix()

	fmt.Println("抓取结束", time.Now())
	fmt.Println("耗时", time.Now(), starts, ends, ends-starts)
}
