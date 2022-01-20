package main

import (
	"crwal/db"
	"fmt"
)

func main() {
	file_path := "https://img.domp4.cc/vod/0/61e786484ebb2.jpg"
	file_path = "https://demibaguette.com/wp-content/uploads/2019/01/erik-witsoe-647316-unsplash-1.jpg"

	id := db.SaveImage(file_path)
	fmt.Println(id)
}
