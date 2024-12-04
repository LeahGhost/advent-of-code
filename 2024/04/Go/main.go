package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var grid []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	word := "XMAS"
	count := countWord(grid, word)
	fmt.Println(count)
}

func countWord(grid []string, word string) int {
	n := len(grid)
	m := len(grid[0])
	wordLen := len(word)
	count := 0
	dirs := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for _, dir := range dirs {
				if checkWord(grid, word, i, j, dir[0], dir[1], n, m, wordLen) {
					count++
				}
			}
		}
	}
	return count
}


func checkWord(grid []string, word string, i, j, dx, dy, n, m, wordLen int) bool {
	for k := 0; k < wordLen; k++ {
		nx, ny := i+dx*k, j+dy*k
		if nx < 0 || nx >= n || ny < 0 || ny >= m || grid[nx][ny] != word[k] {
			return false
		}
	}
	return true
}
	

