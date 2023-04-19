package piscine

func ToUpper(s string) string {
	amad := []rune(s)
	for i := 0; i < len(s); i++ {
		if amad[i] >= 97 && amad[i] <= 122 {
			amad[i] -= 32
		}
	}
	return string(amad)
}
