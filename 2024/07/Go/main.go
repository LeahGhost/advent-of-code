package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func evaluateExpression(numbers []int, target, index, currentResult, part int) bool {
	if index == len(numbers) {
		return currentResult == target
	}

	addResult := evaluateExpression(numbers, target, index+1, currentResult+numbers[index], part)
	multiplyResult := evaluateExpression(numbers, target, index+1, currentResult*numbers[index], part)

	if part == 2 {
		concatResult := evaluateExpression(numbers, target, index+1, atoiConcat(currentResult, numbers[index]), part)
		return addResult || multiplyResult || concatResult
	}

	return addResult || multiplyResult
}

func atoiConcat(left, right int) int {
	return atoi(fmt.Sprintf("%d%d", left, right))
}

func atoi(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func solveCalibration(inputFile string, part int) int {
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

		if evaluateExpression(numbers, target, 1, numbers[0], part) {
			totalCalibrationResult += target
		}
	}

	return totalCalibrationResult
}

func main() {
	inputFile := "../input.txt"

	part1Result := solveCalibration(inputFile, 1)
	fmt.Println("Part 1 - Total Calibration Result:", part1Result)

	part2Result := solveCalibration(inputFile, 2)
	fmt.Println("Part 2 - Total Calibration Result:", part2Result)
}
