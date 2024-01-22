package piscine

func IsLower(s string) bool {
	for _, a := range s {
		if a >= 'a' && a <= 'z' {
			continue
		} else {
			return false
		}
	}
	return true
}
