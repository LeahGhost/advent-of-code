package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fileName := "input.txt"
	data := loadInputData(fileName)
	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}

func expandWarehouseLayout(warehouse [][]rune) [][]rune {
	var modified [][]rune
	for _, row := range warehouse {
		var newRow []rune
		for _, cell := range row {
			switch cell {
			case '.':
				newRow = append(newRow, '.', '.')
			case '#':
				newRow = append(newRow, '#', '#')
			case '@':
				newRow = append(newRow, '@', '.')
			default:
				newRow = append(newRow, '[', ']')
			}
		}
		modified = append(modified, newRow)
	}
	return modified
}

func locateCharacter(target rune, warehouse [][]rune) (int, int, error) {
	for i, row := range warehouse {
		for j, cell := range row {
			if cell == target {
				return i, j, nil
			}
		}
	}
	return -1, -1, fmt.Errorf("character '%c' not found", target)
}

func translateDirectionToSteps(direction rune) (int, int, error) {
	switch direction {
	case '^':
		return -1, 0, nil
	case 'v':
		return 1, 0, nil
	case '<':
		return 0, -1, nil
	case '>':
		return 0, 1, nil
	}
	return 0, 0, fmt.Errorf("invalid direction '%c'", direction)
}

func part1(data []string) int {
	warehouse, instructions := parseInput(data)
	row, col, err := locateCharacter('@', warehouse)
	if err != nil {
		log.Fatal(err)
	}
	for _, instruction := range instructions {
		dRow, dCol, err := translateDirectionToSteps(instruction)
		if err != nil {
			log.Fatal(err)
		}
		nRow, nCol := row+dRow, col+dCol
		steps := 0
		for {
			if newSquare := warehouse[nRow+steps*dRow][nCol+steps*dCol]; newSquare == 'O' {
				steps++
			} else if newSquare == '.' {
				warehouse[row][col] = '.'
				warehouse[nRow][nCol] = '@'
				for i := 1; i <= steps; i++ {
					warehouse[nRow+i*dRow][nCol+i*dCol] = 'O'
				}
				row, col = nRow, nCol
				break
			} else if newSquare == '#' {
				break
			}
		}
	}
	total := 0
	for i, row := range warehouse {
		for j, cell := range row {
			if cell == 'O' {
				total += i*100 + j
			}
		}
	}
	return total
}

func findMovableBlocks(warehouse [][]rune, instruction rune, row, col int) [][]int {
	dRow, dCol, err := translateDirectionToSteps(instruction)
	if err != nil {
		log.Fatal(err)
	}
	newRow, newCol := row+dRow, col+dCol
	var blocksToCheck [][]int
	if warehouse[newRow][newCol] == '#' {
		return nil
	}
	if warehouse[newRow][newCol] == ']' {
		blocksToCheck = append(blocksToCheck, []int{newRow, newCol - 1, newCol})
	} else if warehouse[newRow][newCol] == '[' {
		blocksToCheck = append(blocksToCheck, []int{newRow, newCol, newCol + 1})
	} else if warehouse[newRow][newCol] == '.' {
		return [][]int{{row, col}}
	}
	blocks := [][]int{{row, col}}
	for len(blocksToCheck) > 0 {
		block := blocksToCheck[0]
		blocksToCheck = blocksToCheck[1:]
		if instruction == 'v' || instruction == '^' {
			if warehouse[block[0]+dRow][block[1]] == '#' || warehouse[block[0]+dRow][block[2]] == '#' {
				return nil
			}
			if warehouse[block[0]+dRow][block[1]] == ']' {
				blocksToCheck = append(blocksToCheck, []int{block[0] + dRow, block[1] - 1, block[1]})
			} else if warehouse[block[0]+dRow][block[1]] == '[' {
				blocksToCheck = append(blocksToCheck, []int{block[0] + dRow, block[1], block[2]})
			}
			if warehouse[block[0]+dRow][block[2]] == '[' {
				blocksToCheck = append(blocksToCheck, []int{block[0] + dRow, block[2], block[2] + 1})
			}
		} else if instruction == '<' {
			if warehouse[block[0]][block[1]-1] == '#' {
				return nil
			} else if warehouse[block[0]][block[1]-1] == ']' {
				blocksToCheck = append(blocksToCheck, []int{block[0], block[1] - 2, block[1] - 1})
			}
		} else if instruction == '>' {
			if warehouse[block[0]][block[2]+1] == '#' {
				return nil
			} else if warehouse[block[0]][block[2]+1] == '[' {
				blocksToCheck = append(blocksToCheck, []int{block[0], block[2] + 1, block[2] + 2})
			}
		}
		blocks = append(blocks, block)
	}
	return blocks
}

func calculateGPSScore(warehouse [][]rune) int {
	total := 0
	for i, row := range warehouse {
		for j, cell := range row {
			if cell == '[' {
				total += i*100 + j
			}
		}
	}
	return total
}

func moveBlocks(warehouse *[][]rune, blocks [][]int, instruction rune) *[][]rune {
	dRow, dCol, err := translateDirectionToSteps(instruction)
	if err != nil {
		log.Fatal(err)
	}
	for n := len(blocks) - 1; n >= 0; n-- {
		block := blocks[n]
		if len(block) == 2 {
			(*warehouse)[block[0]+dRow][block[1]+dCol] = '@'
			(*warehouse)[block[0]][block[1]] = '.'
		} else {
			(*warehouse)[block[0]][block[1]] = '.'
			(*warehouse)[block[0]][block[2]] = '.'
			(*warehouse)[block[0]+dRow][block[1]+dCol] = '['
			(*warehouse)[block[0]+dRow][block[2]+dCol] = ']'
		}
	}
	return warehouse
}

func part2(data []string) int {
	warehouse, instructions := parseInput(data)
	warehouse = expandWarehouseLayout(warehouse)
	rx, ry, err := locateCharacter('@', warehouse)
	if err != nil {
		log.Fatal(err)
	}
	for _, instruction := range instructions {
		blocks := findMovableBlocks(warehouse, instruction, rx, ry)
		moveBlocks(&warehouse, blocks, instruction)
		if len(blocks) > 0 {
			dx, dy, err := translateDirectionToSteps(instruction)
			if err != nil {
				log.Fatal(err)
			}
			rx, ry = rx+dx, ry+dy
		}
	}
	return calculateGPSScore(warehouse)
}

func parseInput(data []string) ([][]rune, []rune) {
	var warehouse [][]rune
	var j int
	for i, line := range data {
		j = i
		if line == "" {
			break
		}
		warehouse = append(warehouse, []rune(line))
	}
	return warehouse, []rune(strings.Join(data[j+1:], ""))
}

func loadInputData(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
