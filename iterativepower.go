package piscine

func IterativePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	if power < 0 {
		return 0
	} else {
		resultat := 1
		for a := 1; a <= power; a++ {
			resultat *= nb
		}
		return resultat
	}
}
