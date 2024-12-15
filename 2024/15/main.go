package main

import (
	"log"
	"os"
	"strings"
)

type GameData struct {
	Grid    [][]string
	Moves   []string
	BotX    int
	BotY    int
}

func loadGame(file string) *GameData {
	data, err := os.ReadFile(file)
	if err != nil {
		return &GameData{}
	}
	parts := strings.Split(string(data), "\n\n")
	gridLines := strings.Split(parts[0], "\n")
	game := &GameData{Grid: [][]string{}, Moves: []string{}}

	for y, line := range gridLines {
		row := strings.Split(line, "")
		for x, cell := range row {
			if cell == "@" {
				game.BotX, game.BotY = x, y
			}
		}
		game.Grid = append(game.Grid, row)
	}
	for _, move := range strings.Split(parts[1], "") {
		if move != "\n" {
			game.Moves = append(game.Moves, move)
		}
	}
	return game
}

func isMovable(obj string) bool {
	return obj == "O" || obj == "[" || obj == "]"
}

func shiftBox(game *GameData, x, y int, dir string) bool {
	obj := game.Grid[y][x]
	if !isMovable(obj) {
		return true
	}
	newX, newY := x, y
	switch dir {
	case "left":
		newX--
	case "right":
		newX++
	case "up":
		newY--
	case "down":
		newY++
	}

	if isMovable(game.Grid[newY][newX]) {
		shiftBox(game, newX, newY, dir)
	}
	if obj == "O" && game.Grid[newY][newX] == "." {
		game.Grid[newY][newX], game.Grid[y][x] = "O", "."
		return true
	}
	if obj == "]" && dir == "left" && game.Grid[newY][newX-1] == "." {
		game.Grid[newY][newX-1], game.Grid[newY][newX] = "[", "]"
		return true
	}
	if obj == "[" && dir == "right" && game.Grid[newY][newX+1] == "." {
		game.Grid[newY][newX+1], game.Grid[newY][newX] = "]", "["
		return true
	}
	return false
}

func calculateScore(game *GameData) int {
	score := 0
	for x, row := range game.Grid {
		for y, cell := range row {
			if cell == "O" {
				score += 100*x + y
			}
		}
	}
	return score
}

func moveBot(game *GameData) int {
	for _, move := range game.Moves {
		newX, newY := game.BotX, game.BotY
		dir := ""
		switch move {
		case "<":
			newX--
			dir = "left"
		case ">":
			newX++
			dir = "right"
		case "^":
			newY--
			dir = "up"
		case "v":
			newY++
			dir = "down"
		}
		if dir == "" {
			log.Fatalf("Invalid move at %d/%d", newX, newY)
		}
		shiftBox(game, newX, newY, dir)
		if game.Grid[newY][newX] == "." {
			game.Grid[newY][newX], game.Grid[game.BotY][game.BotX] = "@", "."
			game.BotX, game.BotY = newX, newY
		}
	}
	return calculateScore(game)
}

func main() {
	game := loadGame("input.txt")
	result := moveBot(game)
	log.Printf("Part One: %v", result)
}
