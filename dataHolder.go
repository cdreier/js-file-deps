package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/packr"
)

type dataHolder struct {
	Data     []JSFile `json:"data,omitempty"`
	rootDir  string
	excludes []string
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

func (holder *dataHolder) matchExcludes(path string) bool {
	if len(holder.excludes) == 0 {
		return false
	}
	for _, e := range holder.excludes {
		if strings.Contains(path, e) {
			return true
		}
	}
	return false
}

func (holder *dataHolder) walk(path string, info os.FileInfo, err error) error {
	if strings.HasSuffix(path, fileExtension) {
		smallPath := strings.Replace(path, holder.rootDir, "", 1)
		if holder.matchExcludes(smallPath) {
			return nil
		}
		jsf := JSFile{
			Path: smallPath,
			Hash: buildHash(smallPath),
		}
		fmt.Println("parsing ... ", path)
		jsf.parse(path)
		holder.Data = append(holder.Data, jsf)
	}
	return nil
}
