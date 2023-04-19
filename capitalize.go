package piscine

func IsAlphab(r rune) bool {
	if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
		return true
	}
	return false
}

func Capitalize(s string) string {
	rayane := []rune(s)
	for i, ascii := range rayane {
		if IsAlphab(ascii) {
			if i == 0 || IsAlphab(rayane[i-1]) == false {
				if rayane[i] >= 'a' && rayane[i] <= 'z' {
					rayane[i] = ascii - 32
				}
			} else {
				if rayane[i] >= 'A' && rayane[i] <= 'Z' {
					rayane[i] = ascii + 32
				}
			}
		}
	}
	return string(rayane)
}
