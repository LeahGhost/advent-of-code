package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInput(file string) [][]int {
	f, _ := os.Open(file)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var grid [][]int
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, ch := range line {
			row[i], _ = strconv.Atoi(string(ch))
		}
		grid = append(grid, row)
	}
	return grid
}

func explore(grid [][]int, r, c int) int {
	rows, cols := len(grid), len(grid[0])
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	q := [][2]int{{r, c}}
	visited[r][c] = true
	count := 0

	for len(q) > 0 {
		x, y := q[0][0], q[0][1]
		q = q[1:]
		h := grid[x][y]
		for _, d := range dirs {
			nx, ny := x+d[0], y+d[1]
			if nx >= 0 && nx < rows && ny >= 0 && ny < cols && !visited[nx][ny] && grid[nx][ny] == h+1 {
				visited[nx][ny] = true
				if grid[nx][ny] == 9 {
					count++
				}
				q = append(q, [2]int{nx, ny})
			}
		}
	}
	return count
}

func main() {
	grid := readInput("../input.txt")
	rows, cols := len(grid), len(grid[0])
	total := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 0 {
				total += explore(grid, r, c)
			}
		}
	}

	fmt.Println(total)
}
