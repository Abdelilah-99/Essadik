package piscine

func ToUpper(s string) string {
	var result string
	for _, t := range s {
		if t >= 'a' && t <= 'z' {
			result += string(t - 32)
		} else {
			result += string(t)
		}
	}
	return result
}
