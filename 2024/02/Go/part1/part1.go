package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../../input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeCount := 0

	for scanner.Scan() {
		// Parse the line into integers
		line := scanner.Text()
		levels := parseLine(line)

		if isSafe(levels) {
			safeCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Number of safe reports:", safeCount)
}

// parseLine converts a space-separated line of numbers into a slice of integers.
func parseLine(line string) []int {
	parts := strings.Fields(line)
	numbers := make([]int, len(parts))
	for i, part := range parts {
		numbers[i], _ = strconv.Atoi(part)
	}
	return numbers
}

// isSafe checks if a report satisfies the safety rules.
func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	increasing := true
	decreasing := true

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		// Check if the difference is out of bounds
		if diff < -3 || diff > 3 {
			return false
		}
		if diff > 0 {
			decreasing = false
		} else if diff < 0 {
			increasing = false
		} else {
			return false 
		}
	}

	// At least one of increasing or decreasing must be true
	return increasing || decreasing
}
