package main

import (
	"fmt"
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

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
		// Display the current score using SDL
		scoreString1 := affichage_score(score1)
		scoreString2 := affichage_score(score2)
		scoreText1, _ := font.RenderUTF8_Blended(scoreString1, white)
		scoreText2, _ := font.RenderUTF8_Blended(scoreString2, white)
		scoreText1.Blit(nil, screen, &sdl.Rect{X: 50, Y: 50, W: 100, H: 50})
		scoreText2.Blit(nil, screen, &sdl.Rect{X: 400, Y: 50, W: 100, H: 50})
		screen.Update()

		// Prompt the user for input using SDL
		event := sdl.WaitEvent()
		if event.Type == sdl.MOUSEBUTTONDOWN {
			x, y := event.GetMouseButton().X, event.GetMouseButton().Y
			if x >= 50 && x <= 150 && y >= 50 && y <= 100 {
				score1++
			} else if x >= 400 && x <= 500 && y >= 50 && y <= 100 {
				score2++
			}
		}
		if score1 == 4 && score2 == 4 {
			score1 = 3
			score2 = 3
		}
		if score1 == 5 || score2 == 5 {
			break
		}
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
		// Display the current score using SDL
		scoreString1 := strconv.Itoa(score1)
		scoreString2 := strconv.Itoa(score2)
		scoreText1, _ := font.RenderUTF8_Blended(scoreString1, white)
		scoreText2, _ := font.RenderUTF8_Blended(scoreString2, white)
		scoreText1.Blit(nil, screen, &sdl.Rect{X: 50, Y: 50, W: 100, H: 50})
		scoreText2.Blit(nil, screen, &sdl.Rect{X: 400, Y: 50, W: 100, H: 50})
		screen.Update()

		// Prompt the user for input using SDL
		event := sdl.WaitEvent()
		if event.Type == sdl.MOUSEBUTTONDOWN {
			x, y := event.GetMouseButton().X, event.GetMouseButton().Y
			if x >= 50 && x <= 150 && y >= 50 && y <= 100 {
				score1++
			} else if x >= 400 && x <= 500 && y >= 50 && y <= 100 {
				score2++
			}
		}
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

		// Display the current score using SDL
		scoreString1 := strconv.Itoa(score1)
		scoreString2 := strconv.Itoa(score2)
		scoreText1, _ := font.RenderUTF8_Blended(scoreString1, white)
		scoreText2, _ := font.RenderUTF8_Blended(scoreString2, white)
		scoreText1.Blit(nil, screen, &sdl.Rect{X: 50, Y: 50, W: 100, H: 50})
		scoreText2.Blit(nil, screen, &sdl.Rect{X: 400, Y: 50, W: 100, H: 50})
		screen.Update()
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

		// Display the current score using SDL
		scoreString1 := strconv.Itoa(score1)
		scoreString2 := strconv.Itoa(score2)
		scoreText1, _ := font.RenderUTF8_Blended(scoreString1, white)
		scoreText2, _ := font.RenderUTF8_Blended(scoreString2, white)
		scoreText1.Blit(nil, screen, &sdl.Rect{X: 50, Y: 50, W: 100, H: 50})
		scoreText2.Blit(nil, screen, &sdl.Rect{X: 400, Y: 50, W: 100, H: 50})
		screen.Update()
	}
	if score1 > score2 {
		// Display the winner using SDL
		winnerString := "Le joueur 1 remporte le match."
		winnerText, _ := font.RenderUTF8_Blended(winnerString, white)
		winnerText.Blit(nil, screen, &sdl.Rect{X: 50, Y: 100, W: 500, H: 50})
		screen.Update()
	} else {
		// Display the winner using SDL
		winnerString := "Le joueur 1 remporte le match."
		winnerText, _ := font.RenderUTF8_Blended(winnerString, white)
		winnerText.Blit(nil, screen, &sdl.Rect{X: 50, Y: 100, W: 500, H: 50})
		screen.Update()
	}
}
func main() {
	// Initialize SDL
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sdl.Quit()

	// Create a window and renderer
	window, err := sdl.CreateWindow("Tennis Simulation", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Destroy()
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer renderer.Destroy()

	// Initialize the font
	err = ttf.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ttf.Quit()
	font, err := ttf.OpenFont("font.ttf", 24)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer font.Close()

	// Initialize the screen
	screen, err := sdl.CreateRGBSurface(0, 800, 600, 32, 0, 0, 0, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer screen.Free()

	// Initialize the colors
	white := sdl.Color{R: 255, G: 255, B: 255, A: 255}

	// Start the simulation
	simulation_match()
}
