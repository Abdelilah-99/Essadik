package main

import "github.com/01-edu/z01"

func main() {
	for a := 'a'; a <= 'z'; a++ {
		z01.PrintRune(rune(a))
	}
	z01.PrintRune('\n')
}
