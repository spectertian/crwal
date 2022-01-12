package main

import (
	"crwal/model"
	"testing"
)

func TestGetContentNewAll(t *testing.T) {

	kk := []string{
		//"https://www.domp4.cc/html/tmtO6gOOOOOg.html",
		//"https://www.domp4.cc/html/X1os0SAAAAAS.html",
		//"https://www.domp4.cc/html/S4wW3N77777N.html",
		"https://www.domp4.cc/html/nCTaMXKKKKKX.html",
	}

	for _, v := range kk {
		dy := &model.Dy{}
		dy.Url = v
		ss := GetContentNewAll(dy)
		//t.Log(ss.Director)
		//t.Log(ss.Stars)
		// t.Log(ss.Title)
		// t.Log(ss.Alias)
		//t.Log(ss)
		t.Log(ss.Type)
	}

}
