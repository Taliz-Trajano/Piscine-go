package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	for _, b := range name {
		if b != 46 && b != 47 {
			z01.PrintRune(b)
		}
	}
	z01.PrintRune('\n')
}
