package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions = [][2]int{
	{0, 1}, {1, 1}, {1, 0}, {1, -1},
	{0, -1}, {-1, -1}, {-1, 0}, {-1, 1},
}

func loadFile() []string {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	return grid
}

func isValid(row, col, numRows, numCols int) bool {
	return row >= 0 && row < numRows && col >= 0 && col < numCols
}

func part1(grid []string, numRows, numCols int) int {
	count := 0
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			for _, direction := range directions {
				if checkDirection(grid, row, col, direction[0], direction[1], numRows, numCols, "XMAS") {
					count++
				}
			}
		}
	}
	return count
}

func checkDirection(grid []string, startX, startY, deltaX, deltaY, numRows, numCols int, word string) bool {
	for i := 0; i < len(word); i++ {
		currRow, currCol := startX+i*deltaX, startY+i*deltaY
		if !isValid(currRow, currCol, numRows, numCols) || grid[currRow][currCol] != word[i] {
			return false
		}
	}
	return true
}

// part2 checks for the presence of the patterns "MAS" or "SAM" in an X pattern within a grid of strings.
func part2(grid []string, numRows, numCols int) int {
	count := 0
	checkDirection := func(startX, startY, deltaX, deltaY int) bool {
		startX = startX - deltaX
		startY = startY - deltaY
		word := ""
		for i := 0; i < 3; i++ {
			currRow, currCol := startX+i*deltaX, startY+i*deltaY
			if !isValid(currRow, currCol, numRows, numCols) {
				return false
			}
			word += string(grid[currRow][currCol])
		}
		return word == "MAS" || word == "SAM"
	}

	checkXPattern := func(row, col int) bool {
		if grid[row][col] != 'A' {
			return false
		}
		// Check all four diagonal directions for the patterns "MAS" or "SAM"
		return checkDirection(row, col, -1, 1) &&  // Top-left to bottom-right
			checkDirection(row, col, 1, -1) &&   // Bottom-right to top-left
			checkDirection(row, col, -1, -1) &&  // Top-right to bottom-left
			checkDirection(row, col, 1, 1)       // Bottom-left to top-right
	}

	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			if checkXPattern(row, col) {
				count++
			}
		}
	}
	return count
}

func main() {
	grid := loadFile()
	numRows := len(grid)
	numCols := len(grid[0])
	fmt.Println(part1(grid, numRows, numCols))
	fmt.Println(part2(grid, numRows, numCols))
}