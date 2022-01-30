package main

import (
	"crwal/db"
	"crwal/model"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var Domin = "https://www.tiankongzy.com"
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
				GetInfo(url)
			} else {
				fmt.Println("退出", time.Now().Format("2006-01-02 15:04:05"))
				goto forEnd
			}
		}
	}
forEnd:
	return

}

func GetInfo(url string) {
	if url == "" {
		return
	}
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

	fmt.Println("url", url)
	doc.Find(".xing_vb ul").Each(func(i int, s *goquery.Selection) {
		if s.Find("span").HasClass("tt") {
			href, _ := s.Find("a").Attr("href")
			hrefs := strings.TrimSpace(href)
			urls := Domin + hrefs
			pageDate := strings.TrimSpace(s.Find("span[class=xing_vb7]").Text())
			has := db.IsHasTKCrawl(urls, pageDate)
			fmt.Println("GetDetailByUrl", urls)
			if has == "" {
				saves := GetDetailByUrl(urls)
				fmt.Println("info", saves)
				SaveInfo(&saves)
			} else {
				fmt.Println("已存在跳过")
			}
		} else {
			return
		}
	})
}

func SaveInfo(dy *model.TKStruct) {
	if dy == nil {

		fmt.Println("dy为空", dy)
		return
	}
	dy_id := db.SaveTkDy(dy)

	if dy_id == "" {
		fmt.Println("dy_id为空", dy, dy_id)
		return
	}
	db.SaveTKImageById(dy_id, dy.Pic)
}

func GetDetailByUrl(url string) model.TKStruct {
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

	tk := model.TKStruct{}
	tk.UpdatedTime = time.Now()
	tk.CreatedTime = time.Now()
	tk.Url = url
	tk.Title = strings.TrimSpace(doc.Find(".vodh h2").Text())
	tk.EmTitle = strings.TrimSpace(doc.Find(".vodh span").Text())
	tk.Pic, _ = doc.Find(".vodImg img").Attr("src")

	tk.Rating = strings.TrimSpace(doc.Find(".vodh label").Text())
	tk.Rating = strings.Split(tk.Rating, ":")[1]

	tk.Alias = strings.Split(strings.TrimSpace(doc.Find(".vodinfobox li").Eq(0).Find("span").Text()), ",")
	tk.Director = strings.Split(strings.TrimSpace(doc.Find(".vodinfobox li").Eq(1).Find("span").Text()), ",")
	tk.Stars = strings.Split(strings.TrimSpace(doc.Find(".vodinfobox li").Eq(2).Find("span").Text()), ",")

	tk.Tags = strings.Split(strings.TrimSpace(doc.Find(".vodinfobox li").Eq(3).Find("span").Text()), ",")
	tk.Type = tk.Tags

	tk.Area = strings.TrimSpace(doc.Find(".vodinfobox li").Eq(4).Find("span").Text())
	tk.Language = strings.TrimSpace(doc.Find(".vodinfobox li").Eq(5).Find("span").Text())
	tk.Year = strings.TrimSpace(doc.Find(".vodinfobox li").Eq(6).Find("span").Text())
	tk.PageDate = strings.TrimSpace(doc.Find(".vodinfobox li").Eq(7).Find("span").Text())
	tk.DoubanId = strings.TrimSpace(doc.Find(".vodinfobox li").Eq(8).Find("span").Text())
	tk.Introduction = strings.TrimSpace(doc.Find("vodplayinfo").Eq(0).Text())

	doc.Find(".vodplayinfo").Eq(1).Find("ul").Each(func(i int, s *goquery.Selection) {
		plays := model.TKLStruct{}
		plays.Title = strings.TrimSpace(strings.Split(doc.Find(".vodplayinfo h3").Eq(i).Text(), "：")[1])

		s.Find("li").Each(func(i int, selection *goquery.Selection) {
			re_string := strings.Split(selection.Text(), "$")
			plays.List = append(plays.List, model.TPlayStruct{re_string[0], re_string[1]})
		})

		tk.Play = append(tk.Play, plays)
	})

	return tk
}

func main() {
	fmt.Println("抓取开始", time.Now().Format("2006-01-02 15:04:05"))
	url := "https://www.tiankongzy.com/index.php/index/index/page/%v.html"
	starts := time.Now().Unix()
	fmt.Println(url)
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
	url := "https://www.tiankongzy.com/"
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

	href, _ := doc.Find(".stui-page").Find(" li").Eq(9).Find("a").Attr("href")

	Regexp := regexp.MustCompile(`([^/]*?)\.html`)
	params := Regexp.FindStringSubmatch(href)
	i, err := strconv.Atoi(params[1])
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	PageCount = i
}
