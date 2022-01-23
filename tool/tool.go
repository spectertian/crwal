package main

import (
	"crwal/db"
	"flag"
	"fmt"
	"os"
)

func main() {
	var id = flag.String("id", "", "wiki id")
	fmt.Println("id:", *id)
	if *id == "" {
		fmt.Println("id 不能为空")
		os.Exit(1)
	}
	info := db.GetDyInfoById(*id)
	fmt.Println(info)
	if info.DownStatus == 8 {
		fmt.Println("不存在的的WIKI")
		os.Exit(1)
	}
	db.SaveImageById(info.ID.Hex(), info.Pic)

	fmt.Println("sucess")
}
