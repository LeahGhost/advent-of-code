package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputPath := "./input.txt"
	content, err := os.Open(inputPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer content.Close()

	config := parseData(content)
	fmt.Println(runProgram(config))
}

func parseData(file *os.File) map[string]interface{} {
	result := make(map[string]interface{})
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Register") {
			line = strings.Replace(line, "Register ", "", 1)
		}
		parts := strings.Split(line, ": ")
		if len(parts) == 2 {
			key, value := parts[0], parts[1]
			if key == "Program" {
				program := parseProgram(value)
				result[key] = program
			} else {
				result[key] = parseInt(value)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
	}
	return result
}

func parseProgram(program string) []int {
	strProgram := strings.Split(program, ",")
	var parsedProgram []int
	for _, s := range strProgram {
		val := parseInt(s)
		parsedProgram = append(parsedProgram, val)
	}
	return parsedProgram
}

func parseInt(value string) int {
	val, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("Error parsing int:", err)
	}
	return val
}

func runProgram(data map[string]interface{}) string {
	output := []int{}
	A := data["A"].(int)
	B := data["B"].(int)
	C := data["C"].(int)
	program := data["Program"].([]int)
	pointer := 0

	getValue := func(operand int) int {
		switch {
		case operand <= 3:
			return operand
		case operand == 4:
			return A
		case operand == 5:
			return B
		default:
			return C
		}
	}

	executeInstruction := func(opcode, operand int) {
		switch opcode {
		case 0:
			A = A / (1 << getValue(operand))
		case 1:
			B ^= operand
		case 2:
			B = getValue(operand) % 8
		case 3:
			if A != 0 {
				pointer = operand - 2
			}
		case 4:
			B ^= C
		case 5:
			output = append(output, getValue(operand)%8)
		case 6:
			B = A / (1 << getValue(operand))
		case 7:
			C = A / (1 << getValue(operand))
		}
		pointer += 2
	}

	for pointer < len(program) {
		opcode := program[pointer]
		operand := program[pointer+1]
		executeInstruction(opcode, operand)
	}

	var result []string
	for _, val := range output {
		result = append(result, strconv.Itoa(val))
	}
	return strings.Join(result, ",")
}
