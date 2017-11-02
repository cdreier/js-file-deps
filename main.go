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
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/force", holder.force)
	http.HandleFunc("/noverlap", holder.noverlap)

	fmt.Println("starting server on port 8085")
	http.ListenAndServe(":8085", nil)
}

func (holder *dataHolder) runTemplate(w http.ResponseWriter, html string) {
	holder.Data = make([]JSFile, 0)
	filepath.Walk(holder.rootDir, holder.walk)

	box := packr.NewBox("./frontend")
	t, _ := template.New(html).Parse(box.String(html + ".html"))
	t.Execute(w, holder)
}

func (holder *dataHolder) force(w http.ResponseWriter, r *http.Request) {
	holder.runTemplate(w, "force")
}

func (holder *dataHolder) noverlap(w http.ResponseWriter, r *http.Request) {
	holder.runTemplate(w, "noverlap")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	box := packr.NewBox("./frontend")
	t, _ := template.New("index").Parse(box.String("index.html"))
	t.Execute(w, nil)
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
