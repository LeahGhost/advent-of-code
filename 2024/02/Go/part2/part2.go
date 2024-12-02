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
		line := scanner.Text()
		levels := parseLine(line)

		if isSafeWithDampener(levels) {
			safeCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Number of safe reports:", safeCount)
}

func parseLine(line string) []int {
	parts := strings.Fields(line)
	numbers := make([]int, len(parts))
	for i, part := range parts {
		numbers[i], _ = strconv.Atoi(part)
	}
	return numbers
}

func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	increasing := true
	decreasing := true

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

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

	return increasing || decreasing
}

func isSafeWithDampener(levels []int) bool {
	if isSafe(levels) {
		return true
	}

	for i := 0; i < len(levels); i++ {
		modified := make([]int, 0, len(levels)-1)
		modified = append(modified, levels[:i]...)
		modified = append(modified, levels[i+1:]...)

		if isSafe(modified) {
			return true
		}
	}

	return false
}
