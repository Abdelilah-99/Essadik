package piscine

func IterativePower(nb int, power int) int {
	res := 1

	if power < 0 {
		return 0
	}
	for j := 0; j < power; j++ {
		res = res * nb
	}
	return res
}
