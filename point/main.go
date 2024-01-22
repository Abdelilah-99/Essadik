package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func Pnt(n int) {
	a := '0'
	for i := 0; i < n/10; i++ {
		a++
	}
	z01.PrintRune(a)
	a = '0'

	for i := 0; i < n%10; i++ {
		a++
	}
	z01.PrintRune(a)
}

func main() {
	points := &point{}

	setPoint(points)
	s := "x = i, y = j"
	for _, char := range s {
		if char == 'i' {
			Pnt(points.x)
		} else if char == 'j' {
			Pnt(points.y)
		} else {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}
