package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Machine struct {
	a, b, p, p2 []int
}

func calc(A [2][2]int, b [2]int) [2]float64 {
	det := A[0][0]*A[1][1] - A[0][1]*A[1][0]
	if det == 0 {
		return [2]float64{0, 0}
	}
	return [2]float64{
		float64(b[0]*A[1][1]-b[1]*A[0][1]) / float64(det),
		float64(A[0][0]*b[1]-A[1][0]*b[0]) / float64(det),
	}
}

func (m *Machine) cost(part2 bool) int {
	var p []int
	if part2 {
		p = m.p2
	} else {
		p = m.p
	}

	result := calc(
		[2][2]int{{m.a[0], m.b[0]}, {m.a[1], m.b[1]}},
		[2]int{p[0], p[1]},
	)

	x, y := result[0], result[1]
	if isInteger(x) && isInteger(y) {
		return int(x)*3 + int(y)
	}
	return 0
}

func isInteger(value float64) bool {
	return value == float64(int(value))
}

func parseNumbers(line string) []int {
	numStrs := strings.FieldsFunc(line, func(r rune) bool {
		return r < '0' || r > '9'
	})
	nums := make([]int, len(numStrs))
	for i, numStr := range numStrs {
		num, _ := strconv.Atoi(numStr)
		nums[i] = num
	}
	return nums
}

func main() {
	file, err := os.Open(filepath.Join("..", "input.txt"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input [][]string
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			input = append(input, lines)
			lines = nil
		} else {
			lines = append(lines, line)
		}
	}
	if len(lines) > 0 {
		input = append(input, lines)
	}

	var machines []Machine
	for _, lines := range input {
		machine := Machine{
			a: parseNumbers(lines[0]),
			b: parseNumbers(lines[1]),
			p: parseNumbers(lines[2]),
		}
		machine.p2 = make([]int, len(machine.p))
		for i, x := range machine.p {
			machine.p2[i] = x + int(1e13)
		}
		machines = append(machines, machine)
	}

	sum := 0
	for _, m := range machines {
		sum += m.cost(false)
	}
	fmt.Println(sum)

	sumPart2 := 0
	for _, m := range machines {
		sumPart2 += m.cost(true)
	}
	fmt.Println(sumPart2)
}
