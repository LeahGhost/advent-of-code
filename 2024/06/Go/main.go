package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	part1, part2 := execute(string(input))
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func execute(input string) (int, int) {
	grid := parseGrid(input)
	part1, _ := trackGuard(grid)

	part2 := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid.at(x, y) == '.' {
				updatedGrid := parseGrid(input)
				updatedGrid[y][x] = '#'

				_, loop := trackGuard(updatedGrid)
				if loop {
					part2++
				}
			}
		}
	}

	return part1, part2
}

type Position struct {
	X, Y int
}

type Direction struct {
	X, Y int
	Cell byte
}

type Grid [][]byte

func (g Grid) findGuard() (byte, int, int) {
	for y, row := range g {
		for x, cell := range row {
			if cell == '^' || cell == 'v' || cell == '<' || cell == '>' {
				return cell, x, y
			}
		}
	}
	panic("Guard not found")
}

func (g Grid) at(x, y int) byte {
	if x < 0 || y < 0 || x >= len(g[0]) || y >= len(g) {
		return 0
	}
	return g[y][x]
}

func parseGrid(input string) Grid {
	var grid Grid
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line != "" {
			grid = append(grid, []byte(line))
		}
	}
	return grid
}

func trackGuard(grid Grid) (int, bool) {
	visited := make(map[Position]bool)
	visitedWithDirection := make(map[Direction]bool)

	cell, x, y := grid.findGuard()

	for x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid) {
		if visitedWithDirection[Direction{x, y, cell}] {
			return len(visited), true
		}

		visited[Position{x, y}] = true
		visitedWithDirection[Direction{x, y, cell}] = true

		for {
			switch cell {
			case '^':
				if grid.at(x, y-1) == '#' {
					cell = '>'
					continue
				}
			case '>':
				if grid.at(x+1, y) == '#' {
					cell = 'v'
					continue
				}
			case 'v':
				if grid.at(x, y+1) == '#' {
					cell = '<'
					continue
				}
			case '<':
				if grid.at(x-1, y) == '#' {
					cell = '^'
					continue
				}
			}
			break
		}

		switch cell {
		case '^':
			y--
		case '>':
			x++
		case 'v':
			y++
		case '<':
			x--
		}
	}

	return len(visited), false
}
