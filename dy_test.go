package main

import (
	"testing"
)

func TestGetContent(t *testing.T) {
	dy := &Dy{}
	dy.Url = "https://www.domp4.cc/html/7yt48T44444T.html"
	//dy.Url = "https://www.domp4.cc//detail/12589.html"
	ss := GetContent(dy)
	t.Log(ss)
}

func TestGetContentNew(t *testing.T) {
	dy := &Dy{}
	dy.Url = "https://www.domp4.cc//html/7yt48T44444T.html"
	ss := GetDwonUrlAndDoubanUrl(dy)
	t.Log(ss)
}

func TestGetContentNewAll(t *testing.T) {
	dy := &Dy{}
	dy.Url = "https://www.domp4.cc//html/ReoDhDBBBBBD.html"
	dy.Url = "https://www.domp4.cc//detail/12589.html"
	ss := GetContentNewAll(dy)
	t.Log(ss.Director)
	t.Log(ss.Stars)

}

func TestClickPage(t *testing.T) {
	url := "https://pkg.go.dev/time"
	url = "https://www.domp4.cc/list/99-1.html"
	ClickPages(url)
}
