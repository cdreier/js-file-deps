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
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			token := scanner.Text()

			if token == "import" {
				parseImport(scanner)
			}

		}
	}
}

func parseImport(scanner *bufio.Scanner) {
	for scanner.Scan() {
		token := scanner.Text()
		if token == "from" {
			scanner.Scan()
			// path := scanner.Text()
			// fmt.Println(path)
			break
		}

		if strings.Contains(token, "{") {
			fmt.Println(token)
			// parseNamedImports(scanner)
		} else {
			fmt.Println(token)
		}

	}
}

func parseNamedImports(scanner *bufio.Scanner) {
	for scanner.Scan() {
		token := scanner.Text()
		if strings.Contains(token, "}") {
			break
		}

	}
}
