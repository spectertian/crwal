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
	url := "https://www.tiankongzy.com/"
	GetInfo(url)

}

func TestGetDetailByUrl(t *testing.T) {

	list := []string{
		"https://www.tiankongzy.com/index.php/vod/detail/id/61698.html",
		"https://www.tiankongzy.com/index.php/vod/detail/id/61655.html",
	}
	for _, url := range list {
		ss := GetDetailByUrl(url, 0)
		fmt.Println(ss)
	}

}

func TestGetInfo(t *testing.T) {

	list := []string{
		"https://www.tiankongzy.com/",
	}
	for _, url := range list {
		GetInfo(url)
	}

}
