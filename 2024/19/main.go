package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countWays(design string, patterns []string, memo map[string]int) int {
	if val, ok := memo[design]; ok {
		return val
	}
	if design == "" {
		return 1
	}

	totalWays := 0
	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) {
			totalWays += countWays(design[len(pattern):], patterns, memo)
		}
	}
	memo[design] = totalWays
	return totalWays
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	patterns := strings.Split(scanner.Text(), ", ")

	var designs []string
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			designs = append(designs, line)
		}
	}

	possibleDesigns := 0
	totalArrangements := 0

	for _, design := range designs {
		memo := make(map[string]int)
		ways := countWays(design, patterns, memo)
		if ways > 0 {
			possibleDesigns++
			totalArrangements += ways
		}
	}

	fmt.Print("Part 1:", possibleDesigns)
	fmt.Print("Part 2:", totalArrangements)
}
