package main

import (
	"bufio"
	"fmt"
	"os"
)

func bfsScores(r, c int, grid [][]int, rows, cols int, moves [][2]int) int {
	q := [][2]int{{r, c}}
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	visited[r][c] = true
	nines := 0

	for len(q) > 0 {
		x, y := q[0][0], q[0][1]
		q = q[1:]

		h := grid[x][y]
		for _, move := range moves {
			nx, ny := x+move[0], y+move[1]
			if nx >= 0 && nx < rows && ny >= 0 && ny < cols && !visited[nx][ny] && grid[nx][ny] == h+1 {
				visited[nx][ny] = true
				if grid[nx][ny] == 9 {
					nines++
				}
				q = append(q, [2]int{nx, ny})
			}
		}
	}
	return nines
}

func dfsRatings(r, c int, grid [][]int, rows, cols int, path [][2]int, moves [][2]int, visitedPaths map[string]bool) int {
	pathKey := ""
	for _, p := range path {
		pathKey += fmt.Sprintf("%d,%d|", p[0], p[1])
	}
	if visitedPaths[pathKey] {
		return 0
	}
	visitedPaths[pathKey] = true

	if grid[r][c] == 9 {
		return 1
	}

	trails := 0
	for _, move := range moves {
		nx, ny := r+move[0], c+move[1]
		if nx >= 0 && nx < rows && ny >= 0 && ny < cols && grid[nx][ny] == grid[r][c]+1 {
			trails += dfsRatings(nx, ny, grid, rows, cols, append(path, [2]int{nx, ny}), moves, visitedPaths)
		}
	}
	return trails
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]int
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, ch := range line {
			row[i] = int(ch - '0')
		}
		grid = append(grid, row)
	}

	rows, cols := len(grid), len(grid[0])
	moves := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	visitedPaths := make(map[string]bool)

	var totalScores, totalRatings int
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 0 {
				totalScores += bfsScores(r, c, grid, rows, cols, moves)
				totalRatings += dfsRatings(r, c, grid, rows, cols, [][2]int{{r, c}}, moves, visitedPaths)
			}
		}
	}

	fmt.Printf("PART 1: %d PART 2: %d\n", totalScores, totalRatings)
}
