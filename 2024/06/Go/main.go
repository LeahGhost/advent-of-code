package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var directions = [4][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1}, 
}

func parseInput(filename string) ([][]string, [2]int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, [2]int{}, 0, err
	}
	defer file.Close()

	var grid [][]string
	var guardPos [2]int
	var guardDir int

	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, strings.Split(line, ""))
		for col, cell := range grid[row] {
			switch cell {
			case "^":
				guardPos = [2]int{row, col}
				guardDir = 0
			case ">":
				guardPos = [2]int{row, col}
				guardDir = 1
			case "v":
				guardPos = [2]int{row, col}
				guardDir = 2
			case "<":
				guardPos = [2]int{row, col}
				guardDir = 3
			}
		}
		row++
	}

	if err := scanner.Err(); err != nil {
		return nil, [2]int{}, 0, err
	}

	return grid, guardPos, guardDir, nil
}

func simulateGuardMovement(grid [][]string, guardPos [2]int, guardDir int) int {
	visited := make(map[string]bool)
	visitedPositions := 0

	visitPosition := func(pos [2]int) {
		posStr := fmt.Sprintf("%d,%d", pos[0], pos[1])
		if !visited[posStr] {
			visited[posStr] = true
			visitedPositions++
		}
	}

	visitPosition(guardPos)

	for {
		x, y := guardPos[0], guardPos[1]
		dx, dy := directions[guardDir][0], directions[guardDir][1]
		newPos := [2]int{x + dx, y + dy}

		if newPos[0] < 0 || newPos[0] >= len(grid) || newPos[1] < 0 || newPos[1] >= len(grid[newPos[0]]) {
			break
		}
		if grid[newPos[0]][newPos[1]] == "#" {
			guardDir = (guardDir + 1) % 4
		} else {
			guardPos = newPos
			visitPosition(guardPos)
		}
	}

	return visitedPositions
}

func main() {
	grid, guardPos, guardDir, err := parseInput("../input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	distinctPositions := simulateGuardMovement(grid, guardPos, guardDir)
	fmt.Printf("Distinct positions visited: %d\n", distinctPositions)
}
