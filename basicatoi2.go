package piscine

func BasicAtoi2(s string) int {
	x := []rune(s)
	n := len(s)
	numero := 0
	for i := 0; i < n; i++ {
		if x[i] < '0' || x[i] > '9' {
			return 0
		} else {
			numero *= 10
			numero += int(x[i]) - '0'
		}
	}
	return numero
}
