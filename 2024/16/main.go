package main

import (
	"fmt"
	"os"
	"strings"
)

func readFileLines() [][]string {
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	var filteredLines []string
	for _, line := range lines {
		if line != "" {
			filteredLines = append(filteredLines, line)
		}
	}
	matrix := make([][]string, len(filteredLines))
	for i, line := range filteredLines {
		matrix[i] = strings.Split(line, "")
	}

	return matrix
}

func getAdjacentCoordinates(x, y int, direction string) (int, int) {
	switch direction {
	case "^":
		return x - 1, y
	case "v":
		return x + 1, y
	case ">":
		return x, y + 1
	case "<":
		return x, y - 1
	}
	return -1, -1
}

func getAlternativeDirections(direction string) [3]string {
	switch direction {
	case "^":
		return [3]string{">", "<", "v"}
	case ">":
		return [3]string{"^", "<", "v"}
	case "v":
		return [3]string{"^", ">", "<"}
	case "<":
		return [3]string{"^", ">", "v"}
	}
	fmt.Println("Invalid direction")
	return [3]string{"", "", ""}
}

func updateAlternativeDirections(labyrinth map[[2]int]map[string]int, x, y, score int, direction string) {
	alternatives := getAlternativeDirections(direction)
	for _, alternative := range alternatives {
		if existingScore, exists := labyrinth[[2]int{x, y}][alternative]; exists {
			if score < existingScore {
				labyrinth[[2]int{x, y}][alternative] = score
			}
		} else {
			labyrinth[[2]int{x, y}][alternative] = score
		}
	}
}

func calculateNextMoves(matrix [][]string, labyrinth map[[2]int]map[string]int, x, y int, direction string) {
	currentScore := labyrinth[[2]int{x, y}][direction]
	possibleMoves := []string{}

	if matrix[x+1][y] != "#" {
		possibleMoves = append(possibleMoves, "v")
	}
	if matrix[x-1][y] != "#" {
		possibleMoves = append(possibleMoves, "^")
	}
	if matrix[x][y+1] != "#" {
		possibleMoves = append(possibleMoves, ">")
	}
	if matrix[x][y-1] != "#" {
		possibleMoves = append(possibleMoves, "<")
	}

	for _, move := range possibleMoves {
		newScore := currentScore + 1001
		if move == direction {
			newScore = currentScore + 1
		}

		newX, newY := getAdjacentCoordinates(x, y, move)
		if _, exists := labyrinth[[2]int{newX, newY}]; exists {
			if oldScore, exists := labyrinth[[2]int{newX, newY}][move]; exists && newScore < oldScore {
				labyrinth[[2]int{newX, newY}][move] = newScore
				updateAlternativeDirections(labyrinth, newX, newY, newScore+1000, move)
				calculateNextMoves(matrix, labyrinth, newX, newY, move)
			}
		} else {
			labyrinth[[2]int{newX, newY}] = map[string]int{move: newScore}
			updateAlternativeDirections(labyrinth, newX, newY, newScore+1000, move)
			calculateNextMoves(matrix, labyrinth, newX, newY, move)
		}
	}
}


func initialiseLabyrinth(matrix [][]string) map[[2]int]map[string]int {
	labyrinth := make(map[[2]int]map[string]int)
	for i, row := range matrix {
		for j, cell := range row {
			if cell == "S" {
				labyrinth[[2]int{i, j}] = map[string]int{">": 0}
				calculateNextMoves(matrix, labyrinth, i, j, ">")
				return labyrinth
			}
		}
	}
	return labyrinth
}

func findMinimalScore(matrix [][]string, labyrinth map[[2]int]map[string]int) int {
	for i, row := range matrix {
		for j, cell := range row {
			if cell == "E" {
				if position, exists := labyrinth[[2]int{i, j}]; exists {
					minScore := position["^"]
					for _, score := range position {
						if score < minScore {
							minScore = score
						}
					}
					return minScore
				}
				fmt.Println("End position not calculated")
				return -1
			}
		}
	}
	return -1
}

func getPreviousCoordinates(x, y int, direction string) (int, int) {
	switch direction {
	case "^":
		return x + 1, y
	case ">":
		return x, y - 1
	case "<":
		return x, y + 1
	case "v":
		return x - 1, y
	}
	return -1, -1
}

func determineDirection(x, y, nextX, nextY int) string {
	if x == nextX+1 { return "^" }
	if x == nextX-1 { return "v" }
	if y == nextY+1 { return "<" }
	if y == nextY-1 { return ">" }
	return ""
}

func getStartingDirection(matrix [][]string, labyrinth map[[2]int]map[string]int) (int, int, [][2]int) {
	var endX, endY int
	for i, row := range matrix {
		for j, cell := range row {
			if cell == "E" {
				endX, endY = i, j
				break
			}
		}
	}

	previousCoords := make([][2]int, 0)
	minScore := labyrinth[[2]int{endX, endY}]["^"]
	for _, score := range labyrinth[[2]int{endX, endY}] {
		if score < minScore {
			minScore = score
		}
	}

	for direction, score := range labyrinth[[2]int{endX, endY}] {
		if score == minScore {
			switch direction {
			case "^":
				previousCoords = append(previousCoords, [2]int{endX + 1, endY})
			case "v":
				previousCoords = append(previousCoords, [2]int{endX - 1, endY})
			case ">":
				previousCoords = append(previousCoords, [2]int{endX, endY - 1})
			case "<":
				previousCoords = append(previousCoords, [2]int{endX, endY + 1})
			}
		}
	}
	return endX, endY, previousCoords
}

func traceMinimalPath(matrix [][]string, labyrinth map[[2]int]map[string]int, x, y, nextX, nextY int) {
	direction := determineDirection(x, y, nextX, nextY)
	currentScore := labyrinth[[2]int{x, y}][direction]
	matrix[x][y] = "O"
	moves := map[string][2]int{"^": {x + 1, y}, "v": {x - 1, y}, "<": {x, y + 1}, ">": {x, y - 1}}
	for move, coords := range moves {
		if matrix[coords[0]][coords[1]] != "#" {
			scoreGap := 1
			if move != direction {
				scoreGap = 1001
			}
			oldX, oldY := getPreviousCoordinates(x, y, move)
			if previousScore, exists := labyrinth[[2]int{oldX, oldY}][move]; exists && previousScore == currentScore-scoreGap {
				matrix[oldX][oldY] = "O"
				traceMinimalPath(matrix, labyrinth, oldX, oldY, x, y)
			}
		}
	}
}


func calculateChecksum(matrix [][]string) int {
	sum := 0
	for _, row := range matrix {
		for _, cell := range row {
			if cell == "O" || cell == "E" || cell == "S" {
				sum++
			}
		}
	}
	return sum
}

func main() {
	matrix := readFileLines()
	labyrinth := initialiseLabyrinth(matrix)
	fmt.Printf("Part 1 : %d\n", findMinimalScore(matrix, labyrinth))
	startX, startY, previousCoords := getStartingDirection(matrix, labyrinth)
	for _, prevCoord := range previousCoords {
		traceMinimalPath(matrix, labyrinth, prevCoord[0], prevCoord[1], startX, startY)
	}
	fmt.Printf("Part 2 : %d\n", calculateChecksum(matrix))
}
