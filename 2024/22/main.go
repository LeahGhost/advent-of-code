package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile(filename string) ([]int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return nil, err
		}
		data = append(data, num)
	}
	return data, scanner.Err()
}

func transformValue(v int64) int64 {
	v = (v ^ (v * 64)) % 16777216
	v = (v ^ (v / 32)) % 16777216
	return (v ^ (v * 2048)) % 16777216
}

func calculatePart1(nums []int64) int64 {
	var sum int64
	for _, num := range nums {
		for i := 0; i < 2000; i++ {
			num = transformValue(num)
		}
		sum += num
	}
	return sum
}

func calculatePart2(numbers []int64) int64 {
	frequencyMap := make(map[string]int64)
	for _, num := range numbers {
		sequence := []int64{0}
		seenKeys := make(map[string]struct{})
		remainder := num % 10
		for i := 0; i < 3; i++ {
			num = transformValue(num)
			digit := num % 10
			sequence = append(sequence, digit-remainder)
			remainder = digit
		}
		for i := 3; i < 2000; i++ {
			num = transformValue(num)
			digit := num % 10
			sequence = sequence[1:]
			sequence = append(sequence, digit-remainder)
			key := fmt.Sprintf("%v", sequence)
			if _, seen := seenKeys[key]; !seen {
				seenKeys[key] = struct{}{}
				frequencyMap[key] += digit
			}
			remainder = digit
		}
	}

	var max int64
	for _, value := range frequencyMap {
		if value > max {
			max = value
		}
	}
	return max
}

func main() {
	inputData, err := readFile("./input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(calculatePart1(inputData))
	fmt.Println(calculatePart2(inputData))
}
