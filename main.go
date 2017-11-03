package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gobuffalo/packr"
)

func main() {

	excludes := flag.String("excludes", "", "comma seperated strings")
	flag.Parse()
	rootDir := flag.Args()
	checkRootDir(rootDir[0])

	excludesList := make([]string, 0)
	if len(*excludes) > 0 {
		excludesList = strings.Split(*excludes, ",")
	}

	holder := dataHolder{
		rootDir:  string(rootDir[0]),
		excludes: excludesList,
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

func checkRootDir(dir string) {
	_, err := os.Stat(dir)
	if err != nil {
		log.Fatal("argument needs to be valid directory")
	}
	if os.IsNotExist(err) {
		log.Fatal("argument needs to be valid directory")
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	box := packr.NewBox("./frontend")
	t, _ := template.New("index").Parse(box.String("index.html"))
	t.Execute(w, nil)
}

const fileExtension = ".js"

func buildHash(in string) string {
	sha := sha1.New()
	sha.Write([]byte(in))
	return fmt.Sprintf("%x\n", sha.Sum(nil))
}
