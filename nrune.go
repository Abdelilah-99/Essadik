package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}

	R := []rune(s)
	ln := 0

	for i := range s {
		ln = i + 1
	}
	if ln == 0 || n > ln {
		return 0
	}
	return R[n-1]
}
