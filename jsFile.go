package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// JSFile is the container for parsed js files
type JSFile struct {
	Path    string     `json:"path,omitempty"`
	Exports []jsExport `json:"exports,omitempty"`
	Imports []jsImport `json:"imports,omitempty"`
}

type jsImport struct {
	FromPath      string `json:"from_path,omitempty"`
	Name          string `json:"name,omitempty"`
	DefaultImport bool   `json:"default_import,omitempty"`
}

type jsExport struct {
	fromPath      string
	name          string
	defaultExport bool
}

func (jsf *JSFile) parse(path string) {
	handle, err := os.Open(path)
	if err != nil {
		log.Fatal("faild to open file" + path)
	} else {
		defer handle.Close()
		scanner := bufio.NewScanner(handle)
		// scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			token := scanner.Text()

			if strings.Contains(token, "from ") {
				path := strings.SplitAfter(token, "from ")
				i := jsImport{
					FromPath: path[1],
				}
				jsf.Imports = append(jsf.Imports, i)
			}

		}
	}
}
