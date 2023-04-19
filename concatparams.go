package piscine

func ConcatParams(args []string) string {
	concat := 0
	amad := ""
	for range args {
		concat++
	}
	for i := range args {
		if args[i] != args[concat-1] {
			amad += args[i] + "\n"
		}
	}
	amad += args[concat-1]
	return amad
}
