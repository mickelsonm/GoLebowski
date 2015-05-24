package main

import (
	"fmt"
	"io/ioutil"
	//"strings"
)

func main() {
	data, err := ioutil.ReadFile("script.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileContents := string(data)
	fmt.Println(fileContents)

	//lines := strings.Split(string(data), "\n")
	//for _, line := range lines {
	//	fmt.Println(line)
	//}
}
