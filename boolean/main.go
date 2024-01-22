package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if (nbr % 2) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	Even_str := "I have an even number of arguments"
	Odd_str := "I have an odd number of arguments"
	if len(os.Args) > 1 {
		arg_length := len(os.Args[1:])
		if isEven(arg_length) == true {
			printStr(Even_str)
		} else {
			printStr(Odd_str)
		}
	}
}
