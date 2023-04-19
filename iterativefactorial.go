package piscine

func IterativeFactorial(nb int) int {
	if nb >= 21 {
		return 0
	} else if nb < 0 {
		return 0
	}
	res := 1
	for i := 0; i < nb; i++ {
		res *= 1 + i
	}
	return res
}
