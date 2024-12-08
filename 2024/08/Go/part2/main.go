package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	rows := len(data)
	cols := len(data[0])

	type Antenna struct {
		row, col int
	}

	frequencies := make(map[rune][]Antenna)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			char := rune(data[r][c])
			if char != '.' {
				frequencies[char] = append(frequencies[char], Antenna{r, c})
			}
		}
	}

	antinodes := make(map[string]struct{})
	for _, antennas := range frequencies {
		if len(antennas) < 2 {
			continue
		}

		for i := 0; i < len(antennas); i++ {
			a1 := antennas[i]
			// Each antenna is itself an antinode
			antinodes[fmt.Sprintf("%d,%d", a1.row, a1.col)] = struct{}{}

			for j := i + 1; j < len(antennas); j++ {
				a2 := antennas[j]
				dr := a2.row - a1.row
				dc := a2.col - a1.col

				for k := 1; ; k++ {
					r := a1.row + k*dr
					c := a1.col + k*dc
					if r < 0 || r >= rows || c < 0 || c >= cols {
						break
					}
					antinodes[fmt.Sprintf("%d,%d", r, c)] = struct{}{}
				}

				for k := 1; ; k++ {
					r := a1.row - k*dr
					c := a1.col - k*dc
					if r < 0 || r >= rows || c < 0 || c >= cols {
						break
					}
					antinodes[fmt.Sprintf("%d,%d", r, c)] = struct{}{}
				}
			}
		}
	}

	fmt.Println(len(antinodes))
}
