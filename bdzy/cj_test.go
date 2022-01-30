package main

import (
	"crwal/model"
	"crwal/util"
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
