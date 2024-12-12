package main

import (
	"bufio"
	"image"
	"io"
	"log"
	"os"
	"sort"
)

func readGrid(input io.Reader) [][]rune {
	scanner := bufio.NewScanner(input)
	var grid [][]rune
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	return grid
}

func isInside(grid [][]rune, pt image.Point) bool {
	return pt.In(image.Rect(0, 0, len(grid[0]), len(grid)))
}

var dirs = []image.Point{
	image.Pt(0, 1), 
	image.Pt(1, 0),
	image.Pt(0, -1),
	image.Pt(-1, 0), 
}

func dfs(grid [][]rune, pt image.Point, visited map[image.Point]struct{}) (int64, int64) {
	var area, perim int64
	for _, d := range dirs {
		adj := pt.Add(d)
		if !isInside(grid, adj) || grid[adj.Y][adj.X] != grid[pt.Y][pt.X] {
			perim++
		} else if _, seen := visited[adj]; !seen {
			visited[adj] = struct{}{}
			a, p := dfs(grid, adj, visited)
			area += a
			perim += p
		}
	}
	return area + 1, perim
}

func regionArea(grid [][]rune, pt image.Point, visited map[image.Point]struct{}, edges map[Side]struct{}) int64 {
	var area int64
	for _, d := range dirs {
		adj := pt.Add(d)
		if !isInside(grid, adj) || grid[adj.Y][adj.X] != grid[pt.Y][pt.X] {
			edges[Side{d, pt}] = struct{}{}
			continue
		}
		if _, seen := visited[adj]; seen {
			continue
		}
		visited[adj] = struct{}{}
		area += regionArea(grid, adj, visited, edges)
	}
	return area + 1
}

func countEdges(edges map[Side]struct{}) int64 {
	var count int64
	for _, d := range dirs {
		lines := make(map[int][]int)
		for e := range edges {
			if e.dir == d {
				axis, pos := getAxis(d, e.pt)
				lines[axis] = append(lines[axis], pos)
			}
		}
		for _, positions := range lines {
			sort.Ints(positions)
			count++
			for i := 1; i < len(positions); i++ {
				if positions[i]-positions[i-1] > 1 {
					count++
				}
			}
		}
	}
	return count
}

func getAxis(dir, pt image.Point) (int, int) {
	if dir.X == 0 {
		return pt.Y, pt.X
	}
	return pt.X, pt.Y
}

type Side struct {
	dir image.Point
	pt  image.Point
}

func part1(grid [][]rune) int64 {
	visited := make(map[image.Point]struct{})
	var result int64
	for y := range grid {
		for x := range grid[y] {
			pt := image.Pt(x, y)
			if _, seen := visited[pt]; !seen {
				visited[pt] = struct{}{}
				area, perim := dfs(grid, pt, visited)
				result += area * perim
			}
		}
	}
	return result
}

func part2(grid [][]rune) int64 {
	visited := make(map[image.Point]struct{})
	var result int64
	for y := range grid {
		for x := range grid[y] {
			pt := image.Pt(x, y)
			if _, seen := visited[pt]; seen {
				continue
			}
			visited[pt] = struct{}{}
			edges := make(map[Side]struct{})
			area := regionArea(grid, pt, visited, edges)
			result += area * countEdges(edges)
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

	grid := readGrid(file)

	log.Printf("Part 1: %d", part1(grid))
	log.Printf("Part 2: %d", part2(grid))
}
