package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args[1:]
	length := 0
	for i := range arguments {
		length = i + 1
	}
	if length > 2 {
		fmt.Print("Too many arguments")
		fmt.Print("\n")
	} else if length == 0 {
		fmt.Print("File name missing")
		fmt.Print("\n")
	} else if arguments[0] == "quest8.txt" {
		content, err := ioutil.ReadFile(arguments[0])
		if err != nil {
			fmt.Print(err.Error())
			fmt.Print("\n")
			return
		}
		fmt.Print(string(content))
	}
}
