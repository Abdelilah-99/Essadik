package piscine

func IsPrintable(str string) bool {
	for _, p_able := range str {
		if p_able >= ' ' && p_able <= '~' {
			continue
		} else {
			return false
		}
	}
	return true
}
