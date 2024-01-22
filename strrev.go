package piscine

func StrRev(s string) string {
	lenrev := ""

	for i := strlen(s) - 1; i >= 0; i-- {
		lenrev += string(s[i])
	}
	return lenrev
}

func strlen(s string) int {
	len := 0

	for range s {
		len++
	}
	return len
}
