package main

import (
	"testing"
)

func TestGetContent(t *testing.T) {
	dy := &Dy{}
	dy.Url = "https://www.domp4.cc//html/7yt48T44444T.html"
	ss := GetContent(dy)
	t.Log(ss)
}

func TestGetContentNew(t *testing.T) {
	dy := &Dy{}
	dy.Url = "https://www.domp4.cc//html/7yt48T44444T.html"
	ss := GetContentNew(dy)
	t.Log(ss)
}
