package main

func AlphaCount(s string) int {
	// ru := []rune(s)

	alpha_count := 0

	for _, i := range s /*ru*/ {
		if i >= 'a' && i <= 'z' || i >= 'A' && i <= 'Z' {
			alpha_count++
		}
	}
	return alpha_count
}
