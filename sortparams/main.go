package main

import (
	"os"

	"github.com/01-edu/z01"
)

func SortTable(table []string) {
	for a := 0; a < len(table)-1; a++ {
		for b := a + 1; b < len(table); b++ {
			if table[a] > table[b] {
				table[a], table[b] = table[b], table[a]
			}
		}
	}
}

func main() {
	args := os.Args[1:]
	SortTable(args)
	for _, ar := range args {
		for _, c := range ar {
			z01.PrintRune(rune(c))
		}
		z01.PrintRune('\n')
	}
}
