package piscine

func BasicAtoi(s string) int {
	R := 0
	for i := 0; i <= strlen1(s)-1; i++ {
		R = (R * 10) + int(s[i]-'0')
	}
	return R
}

func strlen1(s string) int {
	len := 0

	for range s {
		len++
	}
	return len
}
