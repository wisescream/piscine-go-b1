package piscine

func Join(strs []string, sep string) string {
	resultat := ""
	rayane := 0
	for i := range strs {
		rayane = i + 1
	}
	for i, str := range strs {
		if i == rayane-1 {
			resultat += str
		} else {
			resultat += str + sep
		}
	}
	return resultat
}
