package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const (
	w = 101
	h = 103
)

type Robot struct {
	px, py, vx, vy int
}

func parseInput() ([]Robot, error) {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	var robots []Robot

	re := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) > 0 {
			px, _ := strconv.Atoi(matches[1])
			py, _ := strconv.Atoi(matches[2])
			vx, _ := strconv.Atoi(matches[3])
			vy, _ := strconv.Atoi(matches[4])
			robots = append(robots, Robot{px, py, vx, vy})
		}
	}

	return robots, nil
}

func getPosition(robot Robot, t int) (int, int) {
	x := (robot.px + robot.vx*t) % w
	if x < 0 {
		x += w
	}

	y := (robot.py + robot.vy*t) % h
	if y < 0 {
		y += h
	}

	return x, y
}

func isChristmasTree(positions []Robot) bool {
	grid := make([][]bool, h)
	for i := range grid {
		grid[i] = make([]bool, w)
	}

	for _, robot := range positions {
		x, y := getPosition(robot, 0) 
		grid[y][x] = true
	}

	tree := [][2]int{
		{0, 0}, {-1, 1}, {0, 1}, {1, 1}, {-2, 2}, {-1, 2}, {0, 2}, {1, 2}, {2, 2},
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			allFound := true
			for _, offset := range tree {
				dx, dy := offset[0], offset[1]
				if !grid[(y+dy+h)%h][(x+dx+w)%w] {
					allFound = false
					break
				}
			}
			if allFound {
				return true
			}
		}
	}

	return false
}

func findFewestSecondsForEasterEgg(robots []Robot) int {
	t := 0
	for {
		if isChristmasTree(robots) {
			break
		}
		t++
	}
	return t
}

func calculatePart1(t int, robots []Robot) int {
	var positions []Robot
	for _, robot := range robots {
		x, y := getPosition(robot, t)
		positions = append(positions, Robot{px: x, py: y})
	}

	midX, midY := w/2, h/2
	counts := [4]int{}
	for _, pos := range positions {
		if pos.px == midX || pos.py == midY {
			continue
		}
		index := boolToInt(pos.px >= midX) + 2*boolToInt(pos.py >= midY)
		counts[index]++
	}

	result := 1
	for _, count := range counts {
		result *= count
	}
	return result
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main() {
	robots, err := parseInput()
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Println("Part 1:", calculatePart1(100, robots))
	fmt.Println("Part 2:", findFewestSecondsForEasterEgg(robots))
}
