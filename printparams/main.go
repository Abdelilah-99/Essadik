package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	for _, str := range arg {
		for _, c := range str {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}
