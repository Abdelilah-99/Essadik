package piscine

func SplitWhiteSpaces(s string) []string {
	var a []string
	var b string
	for i, c := range s {
		if c != ' ' && c != '\t' && c != '\n' {
			b += string(c)
		}
		if ((c == ' ' || c == '\t' || c == '\n') && b != "") || i == len(s)-1 {
			a = append(a, b)
			b = ""
		}
	}
	return a
}
