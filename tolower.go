package piscine

func ToLower(s string) string {
	amad := []rune(s)
	for i := 0; i < len(s); i++ {
		if amad[i] >= 65 && amad[i] <= 90 {
			amad[i] += 32
		}
	}
	return string(amad)
}
