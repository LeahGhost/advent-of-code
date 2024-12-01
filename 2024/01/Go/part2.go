package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Calculate the similarity score between two lists
func calculateSimilarityScore(leftList, rightList []int) int {
	similarityScore := 0

	for _, left := range leftList {
		count := 0
		for _, right := range rightList {
			if left == right {
				count++
			}
		}
		similarityScore += left * count
	}

	return similarityScore
}

// Parse the input into two integer lists
func parseInput(data string) ([]int, []int) {
	lines := strings.Split(data, "\n")
	var leftList, rightList []int

	for _, line := range lines {
		if line == "" {
			continue
		}

		var left, right int
		_, err := fmt.Sscanf(line, "%d %d", &left, &right)
		if err == nil {
			leftList = append(leftList, left)
			rightList = append(rightList, right)
		}
	}

	return leftList, rightList
}

func main() {
	data, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse the input into two lists
	leftList, rightList := parseInput(string(data))

	// Calculate the similarity score
	result := calculateSimilarityScore(leftList, rightList)

	// Output the result
	fmt.Printf("Total Similarity Score: %d\n", result)
}
