package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	filepath.Walk("./example", walk)
}

const fileExtension = ".js"

func walk(path string, info os.FileInfo, err error) error {
	if strings.HasSuffix(path, fileExtension) {
		jsf := JSFile{
			path: path,
		}
		fmt.Println("## start processing ", path)
		jsf.parse(path)
		fmt.Println("## done ")
	}
	return nil
}
