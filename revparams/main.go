package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	index := len(arg) - 1
	for index >= 0 {
		for _, str := range arg[index] {
			z01.PrintRune(str)
		}
		index--
		z01.PrintRune('\n')
	}
}
