package piscine

func Index(s string, toFind string) int {
	rayane := 0
	sb := 0
	if toFind == "" {
		return 0
	}
	for range s {
		rayane++
	}
	for range toFind {
		sb++
	}
	for i := 0; i < rayane; i++ {
		for j := 0; j < sb; j++ {
			if i+j > rayane-1 {
				return -1
			} else if s[i+j] == toFind[j] && j == sb-1 {
				return i
			}
		}
	}
	return -1
}
