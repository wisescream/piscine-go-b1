package piscine

func Swap(a *int, b *int) {
	C := *a
	D := *b
	*a = D
	*b = C
}
