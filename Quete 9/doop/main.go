package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 4 {
		os.Args = os.Args[1:]
		a, errA := strconv.Atoi(os.Args[0])
		b, errB := strconv.Atoi(os.Args[2])
		if errA != nil {
			a = 0
		}
		if errB != nil {
			b = 0
		}

		switch os.Args[1] {
		case "+":
			fmt.Println(a + b)
		}
	}
}
