package main

import (
	"crwal/model"
	"crwal/util"
	"fmt"
	"testing"
)

func TestGetContentBDAll(t *testing.T) {

	kk := []string{
		"https://www.bdzy.tv/index.php/vod/detail/id/41343.html",
	}

	for _, v := range kk {
		dy := &model.BZYStruct{}
		dy.Url = v
		ss := util.GetContentBDAll(dy)
		//t.Log(ss.Director)
		//t.Log(ss.Stars)
		// t.Log(ss.Title)
		// t.Log(ss.Alias)
		//t.Log(ss)
		t.Log(ss.Type)
	}

}

func TestSetPageCounts(t *testing.T) {
	url := "https://api.wujinapi.com/api.php/provide/vod/at/json?ac=detail"
	GetJsonInfo(url)

}

func TestGetDetailByUrl(t *testing.T) {

	list := []int{
		26928226,
		//"https://api.wujinapi.com/api.php/provide/vod/at/json?ac=detail",
		//"https://movie.douban.com/subject/26698862/?from=showing",
		//"https://movie.douban.com/subject/26698862/?from=showing",

	}
	for _, url := range list {
		ss := util.GetDoubanDetailByUrl(url)
		fmt.Println(ss)
	}

}

func TestSaveLocalWiki(t *testing.T) {

	list := []int{
		26928226,
		35698677,
	}
	for _, id := range list {
		SaveLocalWiki(id)
	}

}

func TestGetJsonInfo(t *testing.T) {

	list := []string{
		//"https://api.wujinapi.com/api.php/provide/vod/at/json?ac=detail",
		"https://api.wujinapi.com/api.php/provide/vod/at/json?ac=detail&ids=561",
		"https://api.wujinapi.com/api.php/provide/vod/at/json?ac=detail&ids=28426",
	}
	for _, url := range list {
		GetJsonInfo(url)
	}
}

func TestStruct(*testing.T) {

	type As struct {
		a string
		b string
	}

}
