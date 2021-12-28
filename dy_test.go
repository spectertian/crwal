package main

import (
	mongo_orm "crwal/mongo.orm"
	"sync"
	"testing"
)

func TestGetContent(t *testing.T) {
	dy := &mongo_orm.Dy{}
	dy.Url = "https://www.domp4.cc/html/7yt48T44444T.html"
	//dy.Url = "https://www.domp4.cc//detail/12589.html"
	ss := GetContent(dy)
	t.Log(ss)
}

func TestGetContentNew(t *testing.T) {
	dy := &mongo_orm.Dy{}
	dy.Url = "https://www.domp4.cc//html/7yt48T44444T.html"
	ss := GetDwonUrlAndDoubanUrl(dy)
	t.Log(ss)
}

func TestGetContentNewAll(t *testing.T) {
	dy := &mongo_orm.Dy{}
	dy.Url = "https://www.domp4.cc//html/ReoDhDBBBBBD.html"
	dy.Url = "https://www.domp4.cc//detail/12589.html"
	ss := GetContentNewAll(dy)
	t.Log(ss.Director)
	t.Log(ss.Stars)

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
