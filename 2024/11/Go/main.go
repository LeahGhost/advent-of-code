package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func processStonesOptimized(stones []int, blinks int) int {
	counts := make(map[int]int)
	for _, stone := range stones {
		counts[stone]++
	}

	for i := 0; i < blinks; i++ {
		newCounts := make(map[int]int)

		for stone, count := range counts {
			if stone == 0 {
				newCounts[1] += count
			} else if len(strconv.Itoa(stone))%2 == 0 {
				numStr := strconv.Itoa(stone)
				mid := len(numStr) / 2
				left, _ := strconv.Atoi(numStr[:mid])
				right, _ := strconv.Atoi(numStr[mid:])
				newCounts[left] += count
				newCounts[right] += count
			} else {
				newStone := stone * 2024
				newCounts[newStone] += count
			}
		}
		counts = newCounts
	}
	total := 0
	for _, count := range counts {
		total += count
	}
	return total
}

func main() {
	data, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	stonesStr := strings.Fields(string(data))
	stones := make([]int, len(stonesStr))
	for i, s := range stonesStr {
		stones[i], _ = strconv.Atoi(s)
	}
	part1 := processStonesOptimized(stones, 25)
	part2 := processStonesOptimized(stones, 75)
	fmt.Printf("Part 1: Number of stones after 25 blinks: %d\n", part1)
	fmt.Printf("Part 2: Number of stones after 75 blinks: %d\n", part2)
}
