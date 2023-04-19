package piscine

func NRune(s string, n int) rune {
	i := 1
	rayane := []rune(s)
	for range s {
		if n > len(s) || n <= 0 {
			return 0
		} else {
			i = n - 1
		}
	}
	return rayane[i]
}
