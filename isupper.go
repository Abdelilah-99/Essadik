package piscine

func IsUpper(s string) bool {
	for _, a := range s {
		if a >= 'A' && a <= 'Z' {
			continue
		} else {
			return false
		}
	}
	return true
}
