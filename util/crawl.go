package util

import (
	"context"
	"crwal/model"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

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
	chromeCtx, cancels := chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
	defer cancels()
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

func GetContentNewAll(dy *model.Dy) model.Dy {
	htmlContent, _ := GetHttpHtmlContent(dy.Url, "#download1", "document.querySelector(\"body\")")
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
		return *dy
	}
	dy.Status = 1

	//fmt.Printf(dy.Url)
	dy.DownStatus = 1
	dy.UpdatedTime = time.Now()
	dy.CreatedTime = time.Now()
	dy.Type = []string{strings.TrimSpace(doc.Find(".post-meta span").Eq(0).Find("a").Text())}
	dy.ProductionDate = strings.TrimSpace(doc.Find(".pubtime").Text())
	dy.Pic, _ = doc.Find(".pic img").Attr("src")
	dy.Title = strings.TrimSpace(doc.Find(".text p").Eq(0).Find("span").Text())
	dy.LongTitle = strings.TrimSpace(doc.Find(".article-header h1").Text())
	dy.Type = []string{strings.TrimSpace(doc.Find(".breadcrumb").Find("li").Eq(1).Text())}
	if dy.Type[0] == "电视剧" {
		match, _ := regexp.MatchString(`全\d*集$`, dy.LongTitle)
		if match {
			dy.DownStatus = 1
		} else {
			match2, _ := regexp.MatchString(`更新至\d*集$`, dy.LongTitle)
			if match2 {
				dy.DownStatus = 0
			}
		}
	}

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

	doc.Find(".director").NextUntilMatcher(goquery.Single(".text .tag ")).Each(func(i int, s *goquery.Selection) {
		t_em := strings.TrimSpace(s.Find("em").Text())
		t_em = strings.ReplaceAll(t_em, "：", "")
		if t_em == "地区" {
			dy.Area = strings.TrimSpace(s.Find("span").Text())
		}
		if t_em == "年份" {
			dy.Year = strings.TrimSpace(s.Find("span").Text())
		}
		if t_em == "语言" {
			dy.Language = strings.TrimSpace(s.Find("span").Text())
		}

		if t_em == "更新" {
			dy.UpdatedDate = strings.TrimSpace(s.Find("span").Text())
		}

		if t_em == "时长" {
			dy.RunTime = strings.TrimSpace(s.Find("span").Text())
		}
	})

	tags := []string{}
	doc.Find(".text .tag a").Each(func(i int, s *goquery.Selection) {
		tags = append(tags, strings.TrimSpace(s.Text()))
	})
	dy.Tags = tags

	dy.Introduction = strings.TrimSpace(doc.Find(".article-related").Find("p").Eq(0).Text())

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

func GetContentBDAll(dy *model.BZYStruct) model.BZYStruct {
	res, err := http.Get(dy.Url)
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

	dy.DownStatus = 1
	dy.UpdatedTime = time.Now()
	dy.CreatedTime = time.Now()
	dy.ProductionDate = strings.TrimSpace(doc.Find(".pubtime").Text())
	dy.Pic, _ = doc.Find(".stui-content__thumb .img-responsive").Attr("src")
	dy.Title, _ = doc.Find(".stui-content__thumb .img-responsive").Attr("alt")
	dy.Rating = strings.TrimSpace(doc.Find(".stui-content__detail h1 small").Text())
	dy.Type = []string{strings.TrimSpace(doc.Find(".breadcrumb").Find("li").Eq(1).Text())}

	dy.Introduction = strings.TrimSpace(doc.Find(".stui-content__desc").Text())

	doc.Find(".stui-content__detail p").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			dy.Alias = []string{strings.TrimSpace(s.AfterSelection(s.Find("span")).Text())}
			fmt.Println("别名:", dy.Alias)
			return
		}

		if i == 1 {
			dy.Area = strings.TrimSpace(s.AfterSelection(s.Find("span")).Text())
			fmt.Println("地区:", dy.Area)
			return
		}

		//if i == 2 {
		//	dys := strings.TrimSpace(s.AfterSelection(s.Find("span")).Text())
		//	fmt.Println("别名:", dys)
		//}

		if i == 3 {
			stars := strings.TrimSpace(s.AfterSelection(s.Find("span")).Text())
			dy.Stars = strings.Split(stars, " ")
			fmt.Println("演员:", dy.Stars)
			return
		}

		if i == 4 {
			directors := strings.TrimSpace(s.AfterSelection(s.Find("span")).Text())
			dy.Director = strings.Split(directors, " ")
			fmt.Println("导演:", dy.Director)
		}

		if i == 5 {

			fmt.Println()
			all_html, _ := s.Html()
			//fmt.Println(all_html)

			reg := regexp.MustCompile(`：</span>([^<span]*?)`)
			if reg == nil {
				fmt.Println("MustCompile err")
				return
			}
			//提取关键信息
			result := reg.FindAllStringSubmatch(all_html, -1)

			fmt.Println(result)
			os.Exit(1)
			cs, _ := s.Find("span").Eq(6).Html()
			ss := s.Find("span").Eq(6).NextUntilMatcher(goquery.Single(""))
			fmt.Println(ss.Text())
			fmt.Println(s.Html())
			os.Exit(1)

			fmt.Println("cs", cs)
			tags := []string{}
			s.Find("span").Eq(2).Find("a").Each(func(i int, s *goquery.Selection) {

				tags = append(tags, strings.TrimSpace(s.Text()))
			})
			dy.Tags = tags
			fmt.Println("扩展:", dy.Tags)

			years, _ := s.Find("span").Eq(6).Html()
			fmt.Println("years:", years)
			os.Exit(1)

		}
	})

	if dy.TPTitle == "电视剧" {
		match, _ := regexp.MatchString(`全\d*集$`, dy.LongTitle)
		if match {
			dy.DownStatus = 1
		} else {
			match2, _ := regexp.MatchString(`更新至\d*集$`, dy.LongTitle)
			if match2 {
				dy.DownStatus = 0
			}
		}
	}

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

	tags := []string{}
	doc.Find(".text .tag a").Each(func(i int, s *goquery.Selection) {
		tags = append(tags, strings.TrimSpace(s.Text()))
	})
	dy.Tags = tags

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

func GetDoubanDetailByUrl(wiki_id int) model.Wiki {
	m_url := "https://movie.douban.com/subject/%v/?from=showing"
	url := fmt.Sprintf(m_url, wiki_id)
	fmt.Println("豆瓣地址", url)
	c_count := 0
forStart:
	client := &http.Client{}
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	//增加header选项
	reqest.Header.Add("Referer", "https://movie.douban.com/subject/35241052/?tag=%E7%83%AD%E9%97%A8&from=gaia")
	reqest.Header.Add("Host", "https://search.douban.com/movie/subject_search")
	reqest.Header.Add("Origin", "https://movie.douban.com")
	reqest.Header.Add("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"98\", \"Google Chrome\";v=\"98\"")
	reqest.Header.Add("sec-ch-ua-mobile", "?0")
	reqest.Header.Add("sec-ch-ua-platform", "macOS")
	reqest.Header.Add("Sec-Fetch-Dest", "empty")
	reqest.Header.Add("Sec-Fetch-Mode", "cors")
	reqest.Header.Add("Sec-Fetch-Site", "same-site")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.80 Safari/537.36")
	reqest.Header.Add("Pragma", "no-cache")
	reqest.Header.Add("Cookie", "no-cache")

	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(reqest)
	defer response.Body.Close()

	if response.StatusCode != 200 {
		c_count = c_count + 1
		time.Sleep(time.Second * 1)
		fmt.Println("抓取次数：", c_count, "----", url)
		if c_count > 3 {
			wiki := model.Wiki{}
			wiki.WikiId = 0
			return wiki
			//log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
		}
		response.Body.Close()
		time.Sleep(time.Second * 2)
		goto forStart
		//log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	wiki := model.Wiki{}
	wiki.UpdatedTime = time.Now()
	wiki.CreatedTime = time.Now()
	wiki.WikiId = wiki_id
	wiki.Title = strings.TrimSpace(doc.Find("#content h1 span").Eq(0).Text())

	if wiki.Title == "" {
		wiki := model.Wiki{}
		wiki.WikiId = 0
		return wiki
	}

	year := strings.TrimSpace(doc.Find("#content h1 span").Eq(1).Text())
	re := regexp.MustCompile("[0-9]+")
	newY := re.FindAllString(year, -1)
	wiki.Year, _ = strconv.Atoi(newY[0])
	wiki.Rating, _ = strconv.ParseFloat(strings.TrimSpace(doc.Find(".rating_self .rating_num").Text()), 64)
	info_s := strings.Replace(doc.Find("#info").Text(), " ", "", -1)
	res := strings.Split(info_s, "\n")
	for _, re_str := range res {
		if re_str != "" {
			re_str_last := strings.Split(re_str, ":")
			switch re_str_last[0] {
			case "导演":
				wiki.Director = strings.Split(re_str_last[1], "/")
				break
			case "编剧":
				wiki.Writes = strings.Split(re_str_last[1], "/")
				break
			case "主演":
				wiki.Stars = strings.Split(re_str_last[1], "/")
				break
			case "类型":
				wiki.Tags = strings.Split(re_str_last[1], "/")
				break
			case "制片国家/地区":
				wiki.Area = strings.Split(re_str_last[1], "/")
				break
			case "语言":
				wiki.Language = strings.Split(re_str_last[1], "/")
				break
			case "首播":
				wiki.FirstPlayDate = re_str_last[1]
				break
			case "集数":
				wiki.Episodes = re_str_last[1]
				break
			case "单集片长":
				wiki.EpisodesTime = re_str_last[1]
				break
			case "片长":
				wiki.RunTime = re_str_last[1]
				break
			case "又名":
				wiki.Alias = strings.Split(re_str_last[1], "/")
				break
			case "IMDb":
				wiki.IMDb = re_str_last[1]
				break
			default:
				break
			}
		}
	}

	introduction, _ := doc.Find("#link-report span[property='v:summary']").Html()
	wiki.Introduction = strings.TrimSpace(introduction)
	post_image, _ := doc.Find(".nbgnbg img").Attr("src")
	post_image = strings.Replace(post_image, "/s_ratio_poster/", "/l/", -1)
	wiki.PostImage = strings.TrimSpace(post_image)
	return wiki
}
func GetDoubanHtmlDetailByUrl(wiki_id int) model.Wiki {
	m_url := "https://movie.douban.com/subject/%v/?from=showing"
	url := fmt.Sprintf(m_url, wiki_id)
	fmt.Println("豆瓣地址", url)

	htmlContent, err := GetHttpHtmlContent(m_url, "html", "document.querySelector(\"body\")")

	fmt.Println("xxxx", htmlContent)
	fmt.Println("err", err)
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		wiki := model.Wiki{}
		wiki.WikiId = 0
		return wiki
	}

	wiki := model.Wiki{}
	wiki.UpdatedTime = time.Now()
	wiki.CreatedTime = time.Now()
	wiki.WikiId = wiki_id
	wiki.Title = strings.TrimSpace(doc.Find("#content h1 span").Eq(0).Text())

	if wiki.Title == "" {
		wiki := model.Wiki{}
		wiki.WikiId = 0
		return wiki
	}

	year := strings.TrimSpace(doc.Find("#content h1 span").Eq(1).Text())
	re := regexp.MustCompile("[0-9]+")
	newY := re.FindAllString(year, -1)
	wiki.Year, _ = strconv.Atoi(newY[0])
	wiki.Rating, _ = strconv.ParseFloat(strings.TrimSpace(doc.Find(".rating_self .rating_num").Text()), 64)
	info_s := strings.Replace(doc.Find("#info").Text(), " ", "", -1)
	res := strings.Split(info_s, "\n")
	for _, re_str := range res {
		if re_str != "" {
			re_str_last := strings.Split(re_str, ":")
			switch re_str_last[0] {
			case "导演":
				wiki.Director = strings.Split(re_str_last[1], "/")
				break
			case "编剧":
				wiki.Writes = strings.Split(re_str_last[1], "/")
				break
			case "主演":
				wiki.Stars = strings.Split(re_str_last[1], "/")
				break
			case "类型":
				wiki.Tags = strings.Split(re_str_last[1], "/")
				break
			case "制片国家/地区":
				wiki.Area = strings.Split(re_str_last[1], "/")
				break
			case "语言":
				wiki.Language = strings.Split(re_str_last[1], "/")
				break
			case "首播":
				wiki.FirstPlayDate = re_str_last[1]
				break
			case "集数":
				wiki.Episodes = re_str_last[1]
				break
			case "单集片长":
				wiki.EpisodesTime = re_str_last[1]
				break
			case "片长":
				wiki.RunTime = re_str_last[1]
				break
			case "又名":
				wiki.Alias = strings.Split(re_str_last[1], "/")
				break
			case "IMDb":
				wiki.IMDb = re_str_last[1]
				break
			default:
				break
			}
		}
	}

	introduction, _ := doc.Find("#link-report span[property='v:summary']").Html()
	wiki.Introduction = strings.TrimSpace(introduction)
	post_image, _ := doc.Find(".nbgnbg img").Attr("src")
	post_image = strings.Replace(post_image, "/s_ratio_poster/", "/l/", -1)
	wiki.PostImage = strings.TrimSpace(post_image)
	return wiki
}
