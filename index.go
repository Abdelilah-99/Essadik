package piscine

func Index(s string, toFind string) int {
	len := 0
	sub_len := 0
	for i := range s {
		len = i + 1
	}

	for i := range toFind {
		sub_len = i + 1
	}
	t := 0

	for i := 0; i < len; i++ {
		j := 0
		t = i
		for j < sub_len {
			if t < len && s[t] == toFind[j] {
				j++
				t++
			} else {
				break
			}
		}
		if j == sub_len {
			return i
		}
	}
	return -1
}
