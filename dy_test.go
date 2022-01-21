package main

import (
	"crwal/db"
	"crwal/model"
	"crwal/util"
	"fmt"
	"regexp"
	"sync"
	"testing"
)

func TestMin(t *testing.T) {
	main()
}

func TestGetContentNew(t *testing.T) {
	dy := &model.Dy{}
	dy.Url = "https://www.domp4.cc//html/7yt48T44444T.html"
	ss := util.GetDwonUrlAndDoubanUrl(dy)
	t.Log(ss)
}

func TestGetContentNewAll(t *testing.T) {

	kk := []string{
		//"https://www.domp4.cc/html/tmtO6gOOOOOg.html",
		//"https://www.domp4.cc/html/X1os0SAAAAAS.html",
		//"https://www.domp4.cc/html/S4wW3N77777N.html",
		//"https://www.domp4.cc/html/qcrEUI55555I.html",
		//"https://www.domp4.cc/html/qcrEUI55555I.html",
		"https://www.domp4.cc/detail/15871.html",
	}

	for _, v := range kk {
		dy := &model.Dy{}
		dy.Url = v
		ss := util.GetContentNewAll(dy)
		//t.Log(ss.Director)
		//t.Log(ss.Stars)
		// t.Log(ss.Title)
		// t.Log(ss.Alias)
		t.Log(ss)
		t.Log(ss.Introduction)
		t.Log(ss.ProductionDate)
		t.Log(ss.RunTime)
		t.Log(ss.Language)
		t.Log(ss.Area)
		t.Log(ss.UpdatedDate)
		t.Log("########################################")
	}

}

func TestGetFetchUrl(t *testing.T) {

	list := "https: //www.domp4.cc/list/1-%v.html"
	var wg *sync.WaitGroup
	GetFetchUrl(list, wg)
}

func TestCrwaInfo(t *testing.T) {

	list := []model.Dy{}

	dy := &model.Dy{}
	dy.Url = "https://www.domp4.cc/detail/16459.html"
	list = append(list, *dy)

	dy2 := &model.Dy{}
	dy2.Url = "https://www.domp4.cc/html/SJZIpWVVVVVW.html"
	list = append(list, *dy2)

	for _, v := range list {
		CrwaInfo(&v)
	}

}

func TestZz(t *testing.T) {

	list := []string{
		"2005悬疑美剧《天鹅人》更新至13集",
		"2020生活美剧《联系》第一季更新至全8集",
		"2020搞笑美剧《忙碌黛布拉三连》第一季全集",
	}
	for _, v := range list {

		match, err := regexp.MatchString(`全\d*集$`, v)

		fmt.Println(match, err, v)

		match2, err2 := regexp.MatchString(`更新至\d*集$`, v)

		fmt.Println(match2, err2, v)
	}
}

func TestGetDyInfo(t *testing.T) {

	list := []string{"https://www.domp4.cc//detail/6242.html", "https://www.domp4.cc//html/fNPU26333336.html"}

	for _, v := range list {
		ss := db.GetDyInfo(v)
		fmt.Println(v, ss)
	}

}
