package piscine

func TrimAtoi(s string) int {
	var ordre []int
	resultat := 0
	min := -1
	premiernombre := 0
	index := 0
	compteur := 0
	for _, rune := range s {
		if rune == '-' {
			min = index
		}
		if isDigit(rune) {
			if premiernombre == 0 {
				premiernombre = index
			}
			ordre = append(ordre, int(rune-'0'))
		}
		index++
	}

	for count := range ordre {
		compteur = count + 1
	}

	for i := 0; i < compteur; i++ {
		resultat = resultat*10 + ordre[i]
	}

	if min < premiernombre && min != -1 {
		resultat = resultat * -1
	}
	return resultat
}

func isDigit(digit rune) bool {
	for a := '0'; a <= '9'; a++ {
		if digit == a {
			return true
		}
	}
	return false
}
