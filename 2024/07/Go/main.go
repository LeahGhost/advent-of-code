package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func evaluateExpression(numbers []int, target, index, currentResult int) bool {
	if index == len(numbers) {
		return currentResult == target
	}
	addResult := evaluateExpression(numbers, target, index+1, currentResult+numbers[index])
	multiplyResult := evaluateExpression(numbers, target, index+1, currentResult*numbers[index])
	return addResult || multiplyResult
}

func solveCalibration(inputFile string) int {
	file, _ := os.Open(inputFile)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalCalibrationResult := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		target, _ := strconv.Atoi(parts[0])
		numberStrings := strings.Split(parts[1], " ")
		numbers := make([]int, len(numberStrings))
		for i, numStr := range numberStrings {
			numbers[i], _ = strconv.Atoi(numStr)
		}
		if evaluateExpression(numbers, target, 1, numbers[0]) {
			totalCalibrationResult += target
		}
	}

	return totalCalibrationResult
}

func main() {
	inputFile := "../input.txt"
	result := solveCalibration(inputFile)
	fmt.Println("Total Calibration Result:", result)
}
