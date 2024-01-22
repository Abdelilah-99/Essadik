package piscine

func IsAlpha(s string) bool {
	if len(s) == 0 {
		return true
	}
	for _, b := range s {
		if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9') {
			continue
		} else {
			return false
		}
	}
	return true
}
