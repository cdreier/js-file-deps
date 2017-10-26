package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/packr"
)

func main() {

	rootDir := os.Args[1:]
	_, err := os.Stat(rootDir[0])
	if err != nil {
		log.Fatal("argument needs to be valid directory")
	}
	if os.IsNotExist(err) {
		log.Fatal("argument needs to be valid directory")
	}

	holder := dataHolder{
		rootDir: string(rootDir[0]),
	}

	box := packr.NewBox("./frontend/assets/")
	fs := http.FileServer(box)
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", holder.rootHandler)

	fmt.Println("starting server on port 8085")
	http.ListenAndServe(":8085", nil)
}

func (holder *dataHolder) rootHandler(w http.ResponseWriter, r *http.Request) {
	filepath.Walk(holder.rootDir, holder.walk)

	box := packr.NewBox("./frontend")
	t, _ := template.New("index").Parse(box.String("index.html"))
	t.Execute(w, holder)
}

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
