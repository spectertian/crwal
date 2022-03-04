package main

import (
	"crwal/db"
	"crwal/model"
	"crwal/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup
var PageCount = 0

func GetFetchListUrl(craw_url string) {

	TChannel := make(chan string, PageCount)
	i := 1
	for {
		url := fmt.Sprintf(craw_url, i)
		TChannel <- url
		fmt.Println(url)
		i++
		if PageCount < i {
			break
		}
	}

	close(TChannel)

	wg.Add(8)

	for i := 1; i <= 8; i++ {

		go CrawlInfo(&TChannel, &wg)

	}
	wg.Wait()
}

func CrawlInfo(TChannel *chan string, wg *sync.WaitGroup) {

	defer wg.Done()
	for {
		select {
		case url, ok := <-*TChannel:
			if ok {
				GetJsonInfo(url)
			} else {
				fmt.Println("退出", time.Now().Format("2006-01-02 15:04:05"))
				goto forEnd
			}
		case <-time.After(time.Millisecond * 650):
			fmt.Println("0.5 >>>>>")
		}

	}
forEnd:
	return

}

func GetJsonInfo(url string) {
	c_count := 0
forStart:
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	res.Close = true
	res.Header.Add("Connection", "close")
	defer res.Body.Close()
	if res.StatusCode != 200 {
		c_count = c_count + 1
		time.Sleep(time.Second * 1)
		fmt.Println("抓取次数：", c_count, "----", url)
		if c_count > 3 {
			res.Body.Close()
			fmt.Println("wj抓取失败")
			return
		}
		res.Body.Close()
		goto forStart
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	body, _ := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var m model.JsonResult

	//fmt.Println(string(body))
	if err := json.Unmarshal(body, &m); err != nil {
		fmt.Println("err", err)
		log.Fatal(m)
	}

	for _, info := range m.List {
		fmt.Println(url, info.VodId, info.VodDoubanId, info.VodName)
		info.Director = strings.Split(info.VodDirector, ",")
		info.Class = strings.Split(info.VodClass, ",")
		info.Tags = strings.Split(info.VodTag, ",")
		info.Alias = strings.Split(info.VodSub, ",")
		info.Stars = strings.Split(info.VodActor, ",")

		vod_pay_title := strings.Split(info.VodPlayFrom, "$$$")
		vod_pay_list := strings.Split(info.VodPlayUrl, "$$$")
		fmt.Println(info)
		for k, v := range vod_pay_list {
			v_s := model.VodStruct{}
			v_s.Title = vod_pay_title[k]
			play_list := strings.Split(v, "#")
			for _, u := range play_list {
				if u == "" {
					continue
				}
				res1 := strings.Contains(u, "$")
				if res1 {
					p_list := strings.Split(u, "$")
					v_s.List = append(v_s.List, model.VodPlayStruct{p_list[0], p_list[1]})
				} else {
					v_s.List = append(v_s.List, model.VodPlayStruct{"_", u})
				}
			}
			info.Play = append(info.Play, v_s)
		}

		vod_id := db.SaveAndUpdateWj(&info)
		if vod_id == "" {
			fmt.Println("vod_id为空", info, vod_id)
			return
		}
		db.SaveVodImageById(vod_id, info.VodPic, "wj_")
		//if info.VodDoubanId > 0 {
		//	SaveLocalWiki(info.VodDoubanId)
		//}
	}
	return
}

func SaveLocalWiki(id int) {
	if id == 0 {
		return
	}
	wiki_id := db.IsHasWiki(id)
	if wiki_id == "" {
		wikis := util.GetDoubanDetailByUrl(id)
		if wikis.WikiId == 0 {
			fmt.Println("wiki：", id, "抓取失败")
			return
		}
		insert_id := db.SaveWiki(&wikis)
		if insert_id == "" {
			fmt.Println("wiki insert_id为空")
			return
		}
		db.SaveWikiImageById(insert_id, wikis.PostImage, "wiki_")
	} else {
		fmt.Println("已存在wiki", id)
	}
}

func main() {
	fmt.Println("抓取开始", time.Now().Format("2006-01-02 15:04:05"))
	url := "https://api.wujinapi.com/api.php/provide/vod/at/json?ac=detail&pg=%v"
	starts := time.Now().Unix()
	SetPageCounts()
	fmt.Println("总页数：", PageCount)

	if PageCount < 1 {
		fmt.Println("采集的数据太少")
		os.Exit(1)

	}

	GetFetchListUrl(url)
	ends := time.Now().Unix()

	fmt.Println("抓取结束", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("耗时", starts, ends, ends-starts)
}

func SetPageCounts() {
	url := "https://api.wujinapi.com/api.php/provide/vod/at/json?ac=detail"
	fmt.Println(url)
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

	body, _ := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var m model.JsonResult

	//fmt.Println(string(body))
	if err := json.Unmarshal(body, &m); err != nil {
		fmt.Println("count_err", err)
		log.Fatal(m)
	}
	PageCount = m.PageCount
}
