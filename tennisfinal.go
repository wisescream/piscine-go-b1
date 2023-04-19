package main

import "fmt"

func affichage_score(p int) string {
	if p == 0 {
		return "0"
	} else if p == 1 {
		return "15"
	} else if p == 2 {
		return "30"
	} else if p == 3 {
		return "40"
	} else {
		return "A"
	}
}

func simulation_jeu() int {
	score1 := 0
	score2 := 0                             
	for !(score1 == 3 && score2 < 3) || !(score2 == 3 && score1 < 3) {
		fmt.Print("Joueur gagnant le point : ")
		var winner int
		fmt.Scan(&winner)
		if winner == 1 {
			score1++
		} else {
			score2++
		}
		if score1 == 4 && score2 == 4 {
			score1 = 3
			score2 = 3
		}
		if score1 == 5 || score2 == 5 {
			break
		}
		if score1 == 4 && score2 < 3 || score2 == 4 && score1 < 3 {
			break
		}
		fmt.Println("Score : ", affichage_score(score1), "-", affichage_score(score2))
	}

	if score1 > score2 {
		return 1
	} else {
		return 2
	}
}

func simulation_tiebreak() int {
	score1 := 0
	score2 := 0
	for score1 < 7 || (score1 == 7 && score2 == 6) {
		fmt.Print("Joueur gagnant le point : ")
		var winner int
		fmt.Scan(&winner)
		if winner == 1 {
			score1++
		} else {
			score2++
		}
		fmt.Println("Score : ", score1, "-", score2)
	}
	if score1 > score2 {
		return 1
	} else {
		return 2
	}
}

func simulation_set() int {
	score1 := 0
	score2 := 0
	for score1 < 6 && score2 < 6 || (score1 == 6 && score2 == 6) {
		var winner int
		if score1 == 6 && score2 == 6 {
			winner = simulation_tiebreak()
		} else {
			winner = simulation_jeu()
		}
		if winner == 1 {
			score1++
		} else {
			score2++
		}
		fmt.Println("Score : ", score1, "-", score2)
	}
	if score1 > score2 {
		return 1
	} else {
		return 2
	}
}

func simulation_match() {
	score1 := 0
	score2 := 0
	for score1 < 2 && score2 < 2 {
		var winner int
		winner = simulation_set()
		if winner == 1 {
			score1++
		} else {
			score2++
		}
		fmt.Println("Score : ", score1, "-", score2)
	}
	if score1 > score2 {
		fmt.Println("Le joueur 1 remporte le match.")
	} else {
		fmt.Println("Le joueur 2 remporte le match.")
	}
}

func main() {
	simulation_match()
}
