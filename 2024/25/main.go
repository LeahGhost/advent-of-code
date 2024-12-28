package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("./input.txt")
	lines := strings.TrimSpace(string(data))
	fmt.Println(processData(lines))
}

func processData(data string) int {
	groups := strings.Split(data, "\n\n")
	locks := [][]int{}
	keys := [][]int{}

	for _, group := range groups {
		rows := strings.Split(group, "\n")
		counts := make([]int, len(rows[0]))
		for _, row := range rows {
			for i, ch := range row {
				if ch == '#' {
					counts[i]++
				}
			}
		}
		if rows[0][0] == '#' {
			locks = append(locks, counts)
		} else {
			keys = append(keys, counts)
		}
	}

	matchCount := 0
	for _, lock := range locks {
		for _, key := range keys {
			if isValid(lock, key) {
				matchCount++
			}
		}
	}
	return matchCount
}

func isValid(lock, key []int) bool {
	for i := 0; i < len(lock); i++ {
		if lock[i]+key[i] > 7 {
			return false
		}
	}
	return true
}
