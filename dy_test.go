package main

import (
	"crwal/model"
	"sync"
	"testing"
)

func TestMin(t *testing.T) {
	main()
}
func TestGetContent(t *testing.T) {
	dy := &model.Dy{}
	dy.Url = "https://www.domp4.cc/html/7yt48T44444T.html"
	//dy.Url = "https://www.domp4.cc//detail/12589.html"
	ss := GetContent(dy)
	t.Log(ss)
}

func TestGetContentNew(t *testing.T) {
	dy := &model.Dy{}
	dy.Url = "https://www.domp4.cc//html/7yt48T44444T.html"
	ss := GetDwonUrlAndDoubanUrl(dy)
	t.Log(ss)
}

func TestGetContentNewAll(t *testing.T) {

	kk := []string{
		"https://www.domp4.cc/html/tmtO6gOOOOOg.html",
		"https://www.domp4.cc//html/X1os0SAAAAAS.html",
	}

	for _, v := range kk {
		dy := &model.Dy{}
		dy.Url = v
		ss := GetContentNewAll(dy)
		//t.Log(ss.Director)
		//t.Log(ss.Stars)
		t.Log(ss.Title)
		t.Log(ss.Alias)
	}

}

func TestGetFetchUrl(t *testing.T) {
	//list := []string{
	//	"https: //www.domp4.cc/list/1-%v.html",
	//	"https: //www.domp4.cc/list/2-%v.html",
	//	"https: //www.domp4.cc/list/3-%v.html",
	//	"https: //www.domp4.cc/list/4-%v.html",
	//	"https: //www.domp4.cc/list/5-%v.html",
	//	"https: //www.domp4.cc/list/6-%v.html",
	//	"https: //www.domp4.cc/list/7-%v.html",
	//	"https: //www.domp4.cc/list/8-%v.html",
	//	"https: //www.domp4.cc/list/9-%v.html",
	//	"https: //www.domp4.cc/list/10-%v.html",
	//}
	list := "https: //www.domp4.cc/list/1-%v.html"
	var wg *sync.WaitGroup
	GetFetchUrl(list, wg)

}
