package piscine

func BasicAtoi(s string) int {
	nombre := 0
	n := 0
	a_s := []rune(s)
	for _, word := range a_s {
		for i := '0'; i < word; i++ {
			n++
		}
		nombre = nombre*10 + n
		n = 0
	}
	return nombre
}
