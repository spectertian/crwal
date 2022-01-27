package main

import (
	"crwal/db"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	fmt.Println("begin import image", time.Now().Format("2006-01-02 15:04:05"))

	var path = flag.String("path", "", "image path")
	flag.Parse()
	fmt.Println("path:", *path)
	if *path == "" {
		fmt.Println("path 不能为空")
		os.Exit(1)
	}

	id := db.SaveImage(*path)
	fmt.Println("mongo gridfs id: ", id)

	fmt.Println("end import image", time.Now().Format("2006-01-02 15:04:05"))

}
