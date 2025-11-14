package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type Cell struct {
	X, Y int
}

type Snake struct {
	Body      []Cell
	Direction rune
}

type Game struct {
	Snake *Snake
	Apple *Cell
	Score int
	Best  int
	Over  bool
}

const width = 30
const height = 15

const scoreFile = "score.txt"

func (g *Game) generateRandomApple() {
	for {
		X := rand.Intn(width - 1)
		Y := rand.Intn(height - 1)

		isOccuped := false

		for _, p := range g.Snake.Body {
			if p.X == X && p.Y == Y {
				isOccuped = true
				break
			}
		}

		if !isOccuped {
			g.Apple = &Cell{
				X: X,
				Y: Y,
			}
			return
		}
	}
}

func createNewGame() *Game {

	data, err := os.ReadFile(scoreFile)
	if os.IsNotExist(err) {
		file, err := os.Create(scoreFile)
		if err != nil {
			panic(err)
		}
		_, err = fmt.Fprint(file, "0")
		if err != nil {
			panic(err)
		}
		data = []byte{'0'}
	} else if err != nil {
		panic(err)
	}

	bestScoreStr := string(data)
	bestScore, err := strconv.Atoi(bestScoreStr)
	if err != nil {
		panic(err)
	}

	return &Game{
		Snake: &Snake{
			Body: []Cell{
				{X: width / 2, Y: height / 2},
			},
			Direction: 'd',
		},
		Score: 0,
		Best:  bestScore,
		Over:  false,
	}
}

func (g *Game) update() {
	head := g.Snake.Body[0]
	newHead := head

	switch g.Snake.Direction {
	case 'w':
		newHead.Y--
	case 'a':
		newHead.X--
	case 's':
		newHead.Y++
	case 'd':
		newHead.X++
	}

	if newHead.X < 0 || newHead.Y < 0 || newHead.X >= width || newHead.Y >= height {
		g.Over = true
		return
	}

	for _, p := range g.Snake.Body {
		if p.X == newHead.X && p.Y == newHead.Y {
			g.Over = true
			return
		}
	}

	if newHead == *g.Apple {
		g.Score++
		g.generateRandomApple()
		g.Snake.Body = append([]Cell{newHead}, g.Snake.Body...)
	} else {
		g.Snake.Body = append([]Cell{newHead}, g.Snake.Body[:len(g.Snake.Body)-1]...)
	}
}

func (g *Game) saveBestScore() {
	file, err := os.OpenFile(scoreFile, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = fmt.Fprint(file, g.Best)
	if err != nil {
		panic(err)
	}
}

func (g *Game) draw() {
	fmt.Print("\033[2J\033[H") // Clears the screen

	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}

	if g.Apple != nil {
		grid[g.Apple.Y][g.Apple.X] = '*'
	}

	for i, cell := range g.Snake.Body {
		if i == 0 {
			grid[cell.Y][cell.X] = '@'
		} else {
			grid[cell.Y][cell.X] = 'O'
		}
	}

	fmt.Print("+")
	for range width {
		fmt.Print("-")
	}
	fmt.Println("+")

	for y := range height {
		fmt.Print("|")
		for x := range width {
			fmt.Print(string(grid[y][x]))
		}
		fmt.Println("|")
	}

	fmt.Print("+")
	for range width {
		fmt.Print("-")
	}
	fmt.Println("+")

	fmt.Printf("Score: %d Best Score: %d\n", g.Score, g.Best)
	fmt.Println("Use WASD to move. Press CTRL C to quit")
}

func main() {
	fmt.Println("Welcome to the OG Snake Game")
	game := createNewGame()
	game.generateRandomApple()

	input := make(chan rune)
	reader := bufio.NewReader(os.Stdin)

	go func() {
		for {
			r, _, err := reader.ReadRune()
			if err == nil {
				input <- r
			}
		}
	}()

	ticker := time.NewTicker(600 * time.Millisecond)
	defer ticker.Stop()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)

	for !game.Over {
		select {
		case dir := <-input:
			if dir == 'w' || dir == 'a' || dir == 's' || dir == 'd' {
				game.Snake.Direction = dir
			}
			game.update()
			game.draw()
		case <-ticker.C:
			game.update()
			game.draw()
		case <-sigchan:
			exitCall(game)
		}
	}

	exitCall(game)

}

func exitCall(game *Game) {
	if game.Score > game.Best {
		game.Best = game.Score
		game.saveBestScore()
		fmt.Printf("Congrats on the best score %d\n", game.Best)
	} else {
		fmt.Printf("Score of this game %d, Missed by %d to beat best score :(\n", game.Score, game.Best-game.Score+1)
	}
}
