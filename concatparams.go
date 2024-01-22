package piscine

func ConcatParams(args []string) string {
	var a string
	for i := range args {
		a += args[i]
		if i != len(args)-1 {
			a += "\n"
		}
	}
	return a
}
