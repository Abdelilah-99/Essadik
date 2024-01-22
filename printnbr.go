package piscine

import (
	"github.com/01-edu/z01"
)

func Prinfunc(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n == -9223372036854775808 {
		z01.PrintRune('-')
		z01.PrintRune('9')
		n = 223372036854775808
	}
	if n < 0 {
		z01.PrintRune('-')
		n = n * (-1)
	}
	r(n)
}

func r(n int) {
	if n >= 10 {
		r(n / 10)
		dg := n % 10
		z01.PrintRune(rune(dg) + '0')
	}
	if n < 10 {
		dg := n % 10
		z01.PrintRune(rune(dg) + '0')
	}
}
