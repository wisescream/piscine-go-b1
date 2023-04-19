package piscine

func AppendRange(min, max int) []int {
	var amad []int
	if max > min {
		for i := min - 1; i < max-1; i++ {
			amad = append(amad, i+1)
		}
		return amad
	} else {
		return amad
	}
}
