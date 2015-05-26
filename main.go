package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("script.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//fileContents := string(data)
	//fmt.Println(fileContents)

	//dudeLines := make([]string, 0)
	//walterLines := make([]string, 0)
	//donnyLines := make([]string, 0)
	//theStrangerLines := make([]string, 0)

	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		fmt.Println("#", i, " ", line)
	}
}
