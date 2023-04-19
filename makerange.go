package piscine

func MakeRange(min, max int) []int {
	if max > min {
		amad := make([]int, max-min)
		a := 0
		for i := min; i < max; i++ {
			amad[a] = i
			a++
		}
		return amad
	} else {
		var amad []int = nil
		return amad
	}
}
