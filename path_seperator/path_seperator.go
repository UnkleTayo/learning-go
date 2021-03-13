package path_seperator

import (
	"fmt"
	"path"
)

func Path() {
	var dir, file string

	dir, file = path.Split("css/main.css")

	fmt.Println("dir :", dir)
	fmt.Println("file :", file)
}
