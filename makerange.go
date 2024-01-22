package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	var a []int
	a = make([]int, max-min)
	for i := 0; i < max-min; i++ {
		a[i] = i + min
	}
	return a
}
