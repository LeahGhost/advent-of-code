package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var rules []string
	var updates [][]int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break 
		}
		rules = append(rules, line)
	}

	for scanner.Scan() {
		line := scanner.Text()
		update := parseUpdate(line)
		updates = append(updates, update)
	}

	// Map rules into a usable format
	ruleMap := parseRules(rules)

	// Part 1
	var validMiddleSum int
	var invalidUpdates [][]int
	for _, update := range updates {
		if isValidUpdate(update, ruleMap) {
			midPage := findMiddlePage(update)
			validMiddleSum += midPage
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	// Part 2
	var fixedMiddleSum int
	for _, update := range invalidUpdates {
		fixedUpdate := fixUpdate(update, ruleMap)
		midPage := findMiddlePage(fixedUpdate)
		fixedMiddleSum += midPage
	}

	fmt.Println("Part 1:", validMiddleSum)  // Sum of middle pages of valid updates
	fmt.Println("Part 2:", fixedMiddleSum) // Sum of middle pages of fixed updates
}

func parseUpdate(line string) []int {
	parts := strings.Split(line, ",")
	var update []int
	for _, part := range parts {
		num, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			panic(err)
		}
		update = append(update, num)
	}
	return update
}

func parseRules(rules []string) map[int]map[int]bool {
	ruleMap := make(map[int]map[int]bool)
	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		if len(parts) != 2 {
			panic("Invalid rule format")
		}
		x, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err1 != nil || err2 != nil {
			panic("Invalid numbers in rule")
		}
		if ruleMap[x] == nil {
			ruleMap[x] = make(map[int]bool)
		}
		ruleMap[x][y] = true
	}
	return ruleMap
}

func isValidUpdate(update []int, ruleMap map[int]map[int]bool) bool {
	position := make(map[int]int)
	for i, page := range update {
		position[page] = i
	}

	for x, targets := range ruleMap {
		for y := range targets {
			// If both x and y are in the update, check their order
			posX, okX := position[x]
			posY, okY := position[y]
			if okX && okY && posX >= posY {
				return false 
			}
		}
	}
	return true
}

func fixUpdate(update []int, ruleMap map[int]map[int]bool) []int {
	graph := make(map[int][]int)
	inDegree := make(map[int]int)
	pageSet := make(map[int]bool)

	// Filter rules for pages in this update
	for _, page := range update {
		pageSet[page] = true
	}

	for x, targets := range ruleMap {
		if pageSet[x] {
			for y := range targets {
				if pageSet[y] {
					graph[x] = append(graph[x], y)
					inDegree[y]++
				}
			}
		}
	}

	var sorted []int
	queue := make([]int, 0)

	for _, page := range update {
		if inDegree[page] == 0 {
			queue = append(queue, page)
		}
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		sorted = append(sorted, curr)

		for _, neighbor := range graph[curr] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return sorted
}

func findMiddlePage(update []int) int {
	mid := len(update) / 2
	return update[mid]
}
