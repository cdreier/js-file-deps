package main

import (
	"crypto/sha1"
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
	holder := dataHolder{
		rootDir: "./example",
	}
	filepath.Walk(holder.rootDir, holder.walk)

	t, _ := template.ParseFiles("frontend/index.html")
	t.Execute(w, holder)
})

const fileExtension = ".js"

type dataHolder struct {
	Data    []JSFile `json:"data,omitempty"`
	rootDir string
}

func buildHash(in string) string {
	sha := sha1.New()
	sha.Write([]byte(in))
	return fmt.Sprintf("%x\n", sha.Sum(nil))
}

func (holder *dataHolder) walk(path string, info os.FileInfo, err error) error {
	if strings.HasSuffix(path, fileExtension) {
		jsf := JSFile{
			Path: path,
			Hash: buildHash(path),
		}
		fmt.Println("parsing ... ", path)
		jsf.parse(path)
		holder.Data = append(holder.Data, jsf)
	}
	return nil
}
