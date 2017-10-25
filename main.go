package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	fs := http.FileServer(http.Dir("frontend/assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", rootHandler)

	fmt.Println("starting server on port 8085")
	http.ListenAndServe(":8085", nil)
}

var rootHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	holder := dataHolder{}
	filepath.Walk("./example", holder.walk)

	jsonData, _ := json.Marshal(holder)
	holder.JSON = string(jsonData)

	t, _ := template.ParseFiles("frontend/index.html")
	t.Execute(w, holder)
})

const fileExtension = ".js"

type dataHolder struct {
	Data []JSFile `json:"data,omitempty"`
	JSON string   `json:"JSON,omitempty"`
}

func (holder *dataHolder) walk(path string, info os.FileInfo, err error) error {
	if strings.HasSuffix(path, fileExtension) {
		jsf := JSFile{
			Path: path,
		}
		fmt.Println("parsing ... ", path)
		jsf.parse(path)
		holder.Data = append(holder.Data, jsf)
	}
	return nil
}
