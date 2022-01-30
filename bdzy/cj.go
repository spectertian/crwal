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

var domin = "https://www.bdzy.tv"
var wg sync.WaitGroup

func GetFetchListUrl(info *BDZYStruct, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 1

	fmt.Println(info)
	craw_url := "https://www.bdzy.tv/index.php/vod/type/id/" + info.TId + "/page/%v.html"

	for {
		url := fmt.Sprintf(craw_url, i)

		bd := model.BZYStruct{}
		bd.TId = info.TId
		bd.TTitle = info.TTitle
		bd.TPId = info.TPId
		bd.TPTitle = info.TPTitle
		bd.UpdatedTime = time.Now()
		bd.CreatedTime = time.Now()
		bd.Status = 1
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
		doc.Find(".stui-vodlist").Find("li").Each(func(i int, s *goquery.Selection) {
			if s.HasClass("hidden-xs") {
				return
			}

			hrefs, _ := s.Find("h3 a").Attr("href")
			bd.LongTitle = strings.TrimSpace(s.Find("h3 a").Text())
			titles, _ := s.Find("h3 .title a").Attr("title")
			bd.Title = strings.TrimSpace(titles)

			bd.PageDate = strings.TrimSpace(s.Find(".time").Text())
			bd.Url = domin + strings.TrimSpace(hrefs)

			fmt.Println(hrefs)
			bd.DownStatus = 1
			Regexp := regexp.MustCompile(`([^/]*?)\.html`)
			params := Regexp.FindStringSubmatch(bd.Url)
			bd.CId = params[1]

			if bd.Url == "" {
				fmt.Println("不存在url", bd)
				return
			}

			bd_info := db.GetBDInfo(bd.Url)
			if bd_info.DownStatus == 1 {
				fmt.Println("已保存数据", bd.LongTitle)
				return
			}

			if bd_info.DownStatus == 0 {
				fmt.Println("开始抓取  更新数据", bd.Url, bd.LongTitle, time.Now().Format("2006-01-02 15:04:05"))
			} else {
				fmt.Println("开始抓取", bd.Url, bd.LongTitle, time.Now().Format("2006-01-02 15:04:05"))
			}

			CrwaInfo(&bd)
		})

		_, ok := doc.Find(".pagination li").Eq(5).Find("a").Attr("href")
		if ok == false {
			break
		}
		i++

	}
}

func CrwaInfo(dy *model.BZYStruct) {
	dy_info := util.GetContentBDAll(dy)
	dy_id := db.SaveBDDy(&dy_info)

	db.SaveImageById(dy_id, dy.Pic)
}

type BDZYStruct struct {
	TId     string `bson:"t_id"`
	TTitle  string `bson:"t_title"`
	TPId    string `bson:"t_p_id"`
	TPTitle string `bson:"t_p_title"`
}

func main() {
	fmt.Println("抓取开始", time.Now().Format("2006-01-02 15:04:05"))
	starts := time.Now().Unix()
	wg.Add(36)

	list := []BDZYStruct{
		//BDZYStruct{"20", "电影", "21", "动作片"},
		//BDZYStruct{"20", "电影", "22", "喜剧片"},
		//BDZYStruct{"20", "电影", "23", "爱情片"},
		//BDZYStruct{"20", "电影", "24", "科幻片"},
		//BDZYStruct{"20", "电影", "25", "恐怖片"},
		//BDZYStruct{"20", "电影", "26", "犯罪片"},
		//BDZYStruct{"20", "电影", "27", "战争片"},
		//BDZYStruct{"20", "电影", "28", "动画电影"},
		//BDZYStruct{"20", "电影", "29", "剧情片"},
		//BDZYStruct{"20", "电影", "30", "记录片"},
		//BDZYStruct{"31", "电视剧", "32", "国产剧"},
		//BDZYStruct{"31", "电视剧", "33", "香港剧"},
		//BDZYStruct{"31", "电视剧", "34", "日本剧"},
		//BDZYStruct{"31", "电视剧", "35", "欧美剧"},
		//BDZYStruct{"31", "电视剧", "57", "海外剧"},
		//BDZYStruct{"31", "电视剧", "58", "台湾剧"},
		//BDZYStruct{"31", "电视剧", "59", "韩国剧"},
		//BDZYStruct{"36", "综艺", "37", "大陆综艺"},
		//BDZYStruct{"36", "综艺", "38", "日韩综艺"},
		//BDZYStruct{"36", "综艺", "39", "港台综艺"},
		//BDZYStruct{"36", "综艺", "40", "欧美综艺"},
		//BDZYStruct{"41", "动漫", "42", "国产动漫"},
		//BDZYStruct{"41", "动漫", "43", "日韩动漫"},
		//BDZYStruct{"41", "动漫", "44", "港台动漫"},
		//BDZYStruct{"41", "动漫", "45", "欧美动漫"},
		//BDZYStruct{"46", "体育频道", "49", "WWE摔跤娱乐"},
		//BDZYStruct{"46", "体育频道", "51", "赛车"},
		//BDZYStruct{"46", "体育频道", "52", "UFC终极格斗"},
		//BDZYStruct{"46", "体育频道", "54", "篮球"},
		//BDZYStruct{"46", "体育频道", "55", "足球"},
		//BDZYStruct{"46", "体育频道", "63", "自行车"},
		//BDZYStruct{"46", "体育频道", "64", "棒球"},
		//BDZYStruct{"46", "体育频道", "65", "摩托车"},
		//BDZYStruct{"46", "体育频道", "66", "橄榄球"},
		//BDZYStruct{"46", "体育频道", "77", " 拳击"},
		BDZYStruct{"68", "短片视频", "69", " 歌曲MV"},
	}

	for _, v := range list {
		go GetFetchListUrl(&v, &wg)
	}

	wg.Wait()

	ends := time.Now().Unix()

	fmt.Println("抓取结束", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("耗时", time.Now().Format("2006-01-02 15:04:05"), starts, ends, ends-starts)
}
