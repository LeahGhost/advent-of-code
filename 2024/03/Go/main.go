package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func mainPart1() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content strings.Builder
	for scanner.Scan() {
		content.WriteString(scanner.Text())
	}

	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(content.String(), -1)

	total := 0
	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		total += x * y
	}

	fmt.Println("Sum of all multiplications:", total)
}

func mainPart2() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var content string
	for scanner.Scan() {
		content += scanner.Text()
	}

	pattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	matches := pattern.FindAllStringSubmatch(content, -1)

	isEnabled := true // Multiplications start enabled
	total := 0

	for _, match := range matches {
		if match[0] == "do()" {
			isEnabled = true
		} else if match[0] == "don't()" {
			isEnabled = false
		} else if match[1] != "" && match[2] != "" { // mul(X,Y) case
			if isEnabled {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])
				total += x * y
			}
		}
	}

	fmt.Println("Sum of enabled multiplications:", total)
}

func main() {
    mainPart1() 
    mainPart2() 
}
