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
	ss := GetContentNewAll(dy)
	t.Log(ss)

}
