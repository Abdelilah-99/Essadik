package piscine

func IsNumeric(str string) bool {
	for _, a := range str {
		if a >= '0' && a <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
