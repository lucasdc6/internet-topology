package main

import (
	"bufio"
	"fmt"
	"github.com/thoas/go-funk"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestFile(t *testing.T) {
	files, err := ioutil.ReadDir("../datasets")
	if err != nil {
		panic(err)
	}

	var routes [][]string
	for _, file := range files {
		path := fmt.Sprintf("../datasets/%s", file.Name())

		inFile, err := os.Open(path)
		if err != nil {
			panic(err)
		}

		scanner := bufio.NewScanner(inFile)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			line := strings.Split(scanner.Text(), "\t")
			if funk.Contains(line, "5692") && funk.Contains(line, "D") {
				routes = append(routes, line)
			}
		}
		inFile.Close()
	}
	fmt.Println(routes)
}
