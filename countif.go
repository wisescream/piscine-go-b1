package piscine

func CountIf(f func(string) bool, tab []string) int {
	if len(tab) == 0 {
		return 0
	}
	amad := 0
	for _, s := range tab {
		if f(s) == true {
			amad++
		}
	}
	return amad
}
