package main

import (
	"context"
	"crwal/db"
	"crwal/model"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

var domin = "https://www.domp4.cc/"
var wg sync.WaitGroup

func GetFetchUrl(crawl_url string, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 1
	for {
		url := fmt.Sprintf(crawl_url, i)
		fmt.Println(url)
		dy := &model.Dy{}
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

		fmt.Println("开始抓取详情数据：", url)
		doc.Find("#list_all ul li").Each(func(i int, s *goquery.Selection) {
			hrefs, _ := s.Find(".text_info h2 a").Attr("href")
			dy.LongTitle = strings.TrimSpace(s.Find(".text_info h2 a").Text())
			dy.PageDate = strings.TrimSpace(s.Find(".text_info .update_time").Text())
			dy.Url = domin + strings.TrimSpace(hrefs)
			dy.DownStatus = 0
			Regexp := regexp.MustCompile(`([^/]*?)\.html`)
			params := Regexp.FindStringSubmatch(dy.Url)

			dy.CId = params[1]
			if dy.Url == "" {
				fmt.Println("不存在url", dy)
				return
			}

			if db.IsDyListOk(dy.Url) != "" {
				fmt.Println("已保存数据", dy.LongTitle)
				return
			} else {
				fmt.Println("开始抓取", dy.LongTitle)
			}
			CrwaInfo(dy)
		})

		_, ok := doc.Find(".pagination li").Eq(5).Find("a").Attr("href")
		if ok == false {
			fmt.Println(crawl_url, "抓取完成")
			break
		}
		i++

	}
}
func GetHotList() {
	url := "https://www.domp4.cc/list/99-1.html"
	dy := &model.Dy{}
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

	doc.Find("#list_dy ul li").Each(func(i int, s *goquery.Selection) {
		hrefs, _ := s.Find("a").Attr("href")
		dy.LongTitle = strings.TrimSpace(s.Find("a").Text())
		dy.PageDate = strings.TrimSpace(s.Find("span").Text())
		dy.Url = domin + strings.TrimSpace(hrefs)
		dy.DownStatus = 0
		Regexp := regexp.MustCompile(`([^/]*?)\.html`)
		params := Regexp.FindStringSubmatch(dy.Url)

		dy.CId = params[1]
		if dy.Url == "" {
			fmt.Println("不存在url", dy)
			return
		}

		if db.IsDyListOk(dy.Url) != "" {
			fmt.Println("已保存数据", dy.LongTitle)
			return
		} else {
			fmt.Println("开始抓取", dy.LongTitle)
		}
		CrwaInfo(dy)
	})
	fmt.Println("执行完成", dy.LongTitle)
}

func CrwaInfo(dy *model.Dy) {
	dy_info := GetContentNewAll(dy)
	db.SaveDy(&dy_info)
}

func GetContentNewAll(dy *model.Dy) model.Dy {
	htmlContent, _ := GetHttpHtmlContent(dy.Url, "#download1", "document.querySelector(\"body\")")
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf(dy.Url)

	dy.Type = []string{strings.TrimSpace(doc.Find(".post-meta span").Eq(0).Find("a").Text())}
	dy.ProductionDate = strings.TrimSpace(doc.Find(".pubtime").Text())
	dy.Pic, _ = doc.Find(".pic img").Attr("src")
	dy.Title = strings.TrimSpace(doc.Find(".text p").Eq(0).Find("span").Text())

	em := doc.Find(".text p").Eq(1).Find("em").Text()
	if em == "别名：" {
		alias := strings.TrimSpace(doc.Find(".text p").Eq(1).Find("span").Text())
		re_alias := strings.Split(alias, "/")
		for k, v := range re_alias {
			re_alias[k] = strings.TrimSpace(v)
		}
		dy.Alias = re_alias
	}

	dy.Rating = strings.TrimSpace(doc.Find(".rating_num ").Text())
	dy.DoubanUrl, _ = doc.Find(".rating_num a").Attr("href")
	dy.DoubanId, _ = doc.Find(".rating_num").Attr("subject")

	star := []string{}
	doc.Find(".actor .attrs span").Each(func(i int, s *goquery.Selection) {
		star = append(star, strings.TrimSpace(s.Find("a").Text()))
	})
	dy.Stars = star

	dirct := []string{}
	doc.Find(".director .attrs span").Each(func(i int, s *goquery.Selection) {
		dirct = append(dirct, strings.TrimSpace(s.Find("a").Text()))
	})
	dy.Director = dirct

	dy.Area = strings.TrimSpace(doc.Find(".director").Next().Find("span").Text())
	dy.Year = strings.TrimSpace(doc.Find(".director").Next().Next().Find("span").Text())
	dy.Language = strings.TrimSpace(doc.Find(".director").Next().Next().Next().Find("span").Text())
	dy.RunTime = strings.TrimSpace(doc.Find(".director").Next().Next().Next().Next().Find("span").Text())

	tags := []string{}
	doc.Find(".text .tag a").Each(func(i int, s *goquery.Selection) {
		tags = append(tags, strings.TrimSpace(s.Text()))
	})
	dy.Tags = tags

	dy.Introduction = strings.TrimSpace(doc.Find(".article-related p").Text())

	down_Urls := []model.DownStruct{}
	doc.Find(".url-left").Each(func(i int, s *goquery.Selection) {
		t, _ := s.Find(".url-left a").Attr("title")
		h, _ := s.Find(".url-left a").Attr("href")
		reg, _ := regexp.Compile(`[^:]+`)
		down_Urls = append(down_Urls, model.DownStruct{t, h, reg.FindString(h)})
	})

	dy.DownUrl = down_Urls
	return *dy
}

// 下载地址单独保存
func GetContent(dy *model.Dy) model.Dy {
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
	//fmt.Printf(dy.Url)
	//contents :=
	dy.Type = []string{strings.TrimSpace(doc.Find(".post-meta a").Text())}
	dy.ProductionDate = strings.TrimSpace(doc.Find(".pubtime").Text())
	dy.Pic, _ = doc.Find(".pic img").Attr("src")
	dy.Title = strings.TrimSpace(doc.Find(".text span").Eq(0).Text())
	dy.Rating = strings.TrimSpace(doc.Find(".rating_num ").Text())
	dy.DoubanId, _ = doc.Find(".rating_num").Attr("subject")

	alias := strings.TrimSpace(doc.Find(".text span").Eq(1).Text())
	re_alias := strings.Split(alias, "/")
	for k, v := range re_alias {
		re_alias[k] = strings.TrimSpace(v)
	}
	dy.Alias = re_alias

	star := []string{}
	doc.Find(".actor .attrs span").Each(func(i int, s *goquery.Selection) {
		star = append(star, strings.TrimSpace(s.Find("a").Text()))
	})
	dy.Stars = star

	dirct := []string{}
	doc.Find(".director .attrs span").Each(func(i int, s *goquery.Selection) {
		dirct = append(dirct, strings.TrimSpace(s.Find("a").Text()))
	})
	dy.Director = dirct

	dy.Area = strings.TrimSpace(doc.Find(".director").Next().Find("span").Text())
	dy.Year = strings.TrimSpace(doc.Find(".director").Next().Next().Find("span").Text())
	dy.Language = strings.TrimSpace(doc.Find(".director").Next().Next().Next().Find("span").Text())
	dy.RunTime = strings.TrimSpace(doc.Find(".director").Next().Next().Next().Next().Find("span").Text())

	tags := []string{}
	doc.Find(".text .tag a").Each(func(i int, s *goquery.Selection) {
		tags = append(tags, strings.TrimSpace(s.Text()))
	})
	dy.Tags = tags

	dy.Introduction = strings.TrimSpace(doc.Find(".article-related p").Text())
	return *dy
}

func GetDwonUrlAndDoubanUrl(dy *model.Dy) model.Dy {
	htmlContent, _ := GetHttpHtmlContent(dy.Url, "#download1", "document.querySelector(\"body\")")
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf(dy.Url)
	down_Urls := []model.DownStruct{}
	doc.Find(".url-left").Each(func(i int, s *goquery.Selection) {
		t, _ := s.Find(".url-left a").Attr("title")
		h, _ := s.Find(".url-left a").Attr("href")
		reg, _ := regexp.Compile(`[^:]+`)
		down_Urls = append(down_Urls, model.DownStruct{t, h, reg.FindString(h)})
	})
	dy.DownUrl = down_Urls

	dy.DoubanUrl, _ = doc.Find(".rating_num a").Attr("href")

	return *dy
}

//获取网站上爬取的数据
func GetHttpHtmlContent(url string, selector string, sel interface{}) (string, error) {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true), // debug使用
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	//初始化参数，先传一个空的数据
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, _ := chromedp.NewExecAllocator(context.Background(), options...)

	// create context
	chromeCtx, cancel := chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
	// 执行一个空task, 用提前创建Chrome实例
	chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)

	//创建一个上下文，超时时间为40s
	timeoutCtx, cancel := context.WithTimeout(chromeCtx, 2*time.Minute)
	defer cancel()

	var htmlContent string
	err := chromedp.Run(timeoutCtx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(selector),
		chromedp.OuterHTML(sel, &htmlContent, chromedp.ByJSPath),
	)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	//log.Println(htmlContent)

	return htmlContent, nil
}

func main() {
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
	fmt.Println("抓取结束", time.Time{})
}
