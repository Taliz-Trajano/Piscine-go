package main

import (
	"github.com/01-edu/z01"
)

func main() {

	for ch := 'a'; ch <= 'z'; ch++ {
		z01.PrintRune("%c", ch)
	}
	z01.PrintRune('\n')
}
