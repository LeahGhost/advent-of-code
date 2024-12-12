package main

import (
	"bufio"
	"image"
	"io"
	"log"
	"os"
)

func parseInput(r io.Reader) [][]rune {
	var grid [][]rune
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return grid
}

func isValid(grid [][]rune, p image.Point) bool {
	return p.In(image.Rect(0, 0, len(grid[0]), len(grid)))
}

func explore(grid [][]rune, p image.Point, visited map[image.Point]struct{}) (int64, int64) {
	var area, perimeter int64
	for _, d := range []image.Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		np := p.Add(d)
		if !isValid(grid, np) || grid[np.Y][np.X] != grid[p.Y][p.X] {
			perimeter++
		} else if _, ok := visited[np]; !ok {
			visited[np] = struct{}{}
			na, np := explore(grid, np, visited)
			area += na
			perimeter += np
		}
	}
	return area + 1, perimeter
}

func calculateResult(r io.Reader) int64 {
	grid := parseInput(r)
	visited := make(map[image.Point]struct{})
	var result int64
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			p := image.Pt(x, y)
			if _, ok := visited[p]; !ok {
				visited[p] = struct{}{}
				area, perimeter := explore(grid, p, visited)
				result += area * perimeter
			}
		}
	}
	return result
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := calculateResult(file)
	log.Printf("Result: %d", result)
}
