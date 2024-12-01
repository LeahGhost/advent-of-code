package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

// Calculate the total distance between two lists of numbers
func calculateTotalDistance(leftList, rightList []int) int {
	// Sort both lists
	sort.Ints(leftList)
	sort.Ints(rightList)

	// Calculate the total distance
	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		totalDistance += int(math.Abs(float64(leftList[i] - rightList[i])))
	}

	return totalDistance
}

// Reads input from the file and returns two slices of integers
func ReadFile(filename string) ([]int, []int, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}

	// Split the input into lines
	lines := strings.Split(string(data), "\n")

	// Prepare two slices for the left and right lists
	var leftList, rightList []int

	// Parse the lines into integers for both lists
	for i, line := range lines {
		if line == "" {
			continue
		}
		// Assume each line has two numbers, one for each list
		var left, right int
		_, err := fmt.Sscanf(line, "%d %d", &left, &right)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing line %d: %v", i+1, err)
		}
		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	return leftList, rightList, nil
}

func main() {
	inputFilePath := "../input.txt"
	
	// Read the input from the file
	leftList, rightList, err := ReadFile(inputFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Calculate the total distance
	result := calculateTotalDistance(leftList, rightList)
	fmt.Printf("Total Distance: %d\n", result)
}
