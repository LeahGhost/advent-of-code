package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("../../input.txt")
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	rows := len(data)
	cols := len(data[0])

	type Signal struct {
		row, col int
		freq     rune
	}

	var signals []Signal
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			char := rune(data[r][c])
			if char != '.' {
				signals = append(signals, Signal{r, c, char})
			}
		}
	}

	locations := make(map[string]struct{})
	recordAntinode := func(r, c int) {
		if r >= 0 && r < rows && c >= 0 && c < cols {
			key := fmt.Sprintf("%d,%d", r, c)
			locations[key] = struct{}{}
		}
	}

	for i := 0; i < len(signals); i++ {
		for j := i + 1; j < len(signals); j++ {
			s1 := signals[i]
			s2 := signals[j]
			if s1.freq != s2.freq {
				continue
			}

			diffR := s2.row - s1.row
			diffC := s2.col - s1.col

			if diffR%2 == 0 && diffC%2 == 0 {
				midR := s1.row + diffR/2
				midC := s1.col + diffC/2
				recordAntinode(midR, midC)
			}
			recordAntinode(s2.row+diffR, s2.col+diffC)
			recordAntinode(s1.row-diffR, s1.col-diffC)
		}
	}

	fmt.Println(len(locations))
}
