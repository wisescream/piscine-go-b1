package piscine

func Map(f func(int) bool, a []int) []bool {
	var b []bool
	for i := range a {
		b = append(b, f(a[i]))
	}
	return b
}
