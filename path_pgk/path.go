package main

import (
	"fmt"
	"path"
)

func main() {
	var (
		dir  string
		file string
	)
	dir, file = path.Split("temp/kk.txt")
	fmt.Println("Dir:", dir)
	fmt.Println("file:", file)
}
