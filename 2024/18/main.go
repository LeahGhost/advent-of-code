package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	GRID_SIZE              = 71
	MAX_CORRUPTED_POSITIONS = 1024
)

var directions = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func bfs(grid [][]rune, start, end [2]int) int {
	queue := [][2]int{start}
	visited := map[[2]int]bool{start: true}
	for steps := 0; len(queue) > 0; steps++ {
		nextQueue := [][2]int{}
		for _, pos := range queue {
			if pos == end {
				return steps
			}
			for _, d := range directions {
				next := [2]int{pos[0] + d[0], pos[1] + d[1]}
				if next[0] >= 0 && next[0] < GRID_SIZE && next[1] >= 0 && next[1] < GRID_SIZE && grid[next[0]][next[1]] == '.' && !visited[next] {
					visited[next] = true
					nextQueue = append(nextQueue, next)
				}
			}
		}
		queue = nextQueue
	}
	return -1
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var corruptedPositions [][2]int
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		corruptedPositions = append(corruptedPositions, [2]int{x, y})
	}

	grid := make([][]rune, GRID_SIZE)
	for i := range grid {
		grid[i] = make([]rune, GRID_SIZE)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	for _, pos := range corruptedPositions[:MAX_CORRUPTED_POSITIONS] {
		grid[pos[0]][pos[1]] = '#'
	}

	fmt.Printf("Minimum steps to reach the exit: %d\n", bfs(grid, [2]int{0, 0}, [2]int{70, 70}))

	for _, pos := range corruptedPositions {
		grid[pos[0]][pos[1]] = '#'
		if bfs(grid, [2]int{0, 0}, [2]int{70, 70}) == -1 {
			fmt.Printf("First byte that blocks the path: (%d, %d)\n", pos[0], pos[1])
			break
		}
	}
}
