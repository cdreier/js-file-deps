package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// JSFile is the container for parsed js files
type JSFile struct {
	path    string
	exports []jsExport
	imports []jsImport
}

type jsImport struct {
	fromPath      string
	name          string
	defaultImport bool
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
		for scanner.Scan() {
			line := scanner.Text()
			jsf.handleOneLineImportStatement(line)
		}
	}
}

const importKeyword = "import "

func isOneLineImportStatement(line string) bool {
	return strings.HasPrefix(line, importKeyword) && strings.Contains(line, " from ")
}

func (jsf *JSFile) handleOneLineImportStatement(line string) {
	if isOneLineImportStatement(line) {
		tmp := line[len(importKeyword):len(line)]
		parts := strings.Split(tmp, " from ")
		isDefaultImport := !strings.HasPrefix(parts[0], "{")
		importString := strings.SplitN(parts[0], " ", 2)

		if isDefaultImport {
			i := jsImport{
				defaultImport: isDefaultImport,
				fromPath:      parts[1],
			}
			defaultImportName := strings.Trim(importString[0], " ,")
			i.name = defaultImportName
			jsf.imports = append(jsf.imports, i)
			fmt.Println("added", i)
		} else {
			jsf.handleNamedImports(strings.Trim(parts[0], " ,"), parts[1])
		}

		if len(importString) == 2 {
			jsf.handleNamedImports(importString[1], parts[1])
		}

	}
}

func (jsf *JSFile) handleNamedImports(imports string, path string) {
	imports = strings.Replace(imports, "{", "", 1)
	imports = strings.Replace(imports, "}", "", 1)
	namedImports := strings.Split(imports, ",")

	for s := range namedImports {
		i := jsImport{
			defaultImport: false,
			fromPath:      path,
			name:          strings.Trim(namedImports[s], " ,"),
		}

		jsf.imports = append(jsf.imports, i)
		fmt.Println("added", i)
	}

}
