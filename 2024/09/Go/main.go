package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInputFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text(), nil
}

func computeDiskChecksum() {
	input, err := readInputFile("../input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	processingFile := true
	currentFileIndex := 0
	var diskLayout []interface{}

	for _, size := range input {
		blockSize, _ := strconv.Atoi(string(size))
		if processingFile {
			for i := 0; i < blockSize; i++ {
				diskLayout = append(diskLayout, currentFileIndex)
			}
			currentFileIndex++
		} else {
			for i := 0; i < blockSize; i++ {
				diskLayout = append(diskLayout, ".")
			}
		}
		processingFile = !processingFile
	}

	for i := 0; i < len(diskLayout); i++ {
		if diskLayout[i] == "." {
			for j := len(diskLayout) - 1; j > i; j-- {
				if diskLayout[j] != "." {
					diskLayout[i] = diskLayout[j]
					diskLayout[j] = "."
					break
				}
			}
		}
	}

	totalChecksum := 0
	for i, block := range diskLayout {
		if block != "." {
			blockVal, _ := block.(int)
			totalChecksum += blockVal * i
		}
	}

	fmt.Println("Checksum from Part 1:", totalChecksum)
}

func calculateOptimisedChecksum() {
	condensedDiskMap, err := readInputFile("../input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	isFile := true
	fileIndex := 0
	var fullDiskMap [][]interface{}

	for _, entryLength := range condensedDiskMap {
		entryLen, _ := strconv.Atoi(string(entryLength))
		if isFile {
			block := make([]interface{}, entryLen)
			for i := range block {
				block[i] = fileIndex
			}
			fullDiskMap = append(fullDiskMap, block)
			fileIndex++
			isFile = false
		} else {
			if entryLen > 0 {
				block := make([]interface{}, entryLen)
				for i := range block {
					block[i] = "."
				}
				fullDiskMap = append(fullDiskMap, block)
			}
			isFile = true
		}
	}

	for i := len(fullDiskMap) - 1; i >= 0; i-- {
		if !contains(fullDiskMap[i], ".") {
			for j := 0; j < i; j++ {
				if contains(fullDiskMap[j], ".") {
					requiredSegmentLength := len(fullDiskMap[i])
					currentSegmentLength := 0
					segmentStartIndex := -1
					for m := 0; m < len(fullDiskMap[j]); m++ {
						if fullDiskMap[j][m] == "." {
							currentSegmentLength++
							if segmentStartIndex == -1 {
								segmentStartIndex = m
							}
						} else {
							currentSegmentLength = 0
							segmentStartIndex = -1
						}
						if currentSegmentLength == requiredSegmentLength {
							break
						}
					}

					if currentSegmentLength == requiredSegmentLength && segmentStartIndex != -1 {
						for m := 0; m < len(fullDiskMap[i]); m++ {
							fullDiskMap[j][segmentStartIndex] = fullDiskMap[i][m]
							fullDiskMap[i][m] = "."
							segmentStartIndex++
						}
						break
					}
				}
			}
		}
	}

	flattenedDiskMap := flatten(fullDiskMap)
	optimisedChecksum := 0
	for i, block := range flattenedDiskMap {
		if block != "." {
			blockVal, _ := block.(int)
			optimisedChecksum += blockVal * i
		}
	}

	fmt.Println("Optimised Checksum from Part 2:", optimisedChecksum)
}

func contains(arr []interface{}, value interface{}) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func flatten(arr [][]interface{}) []interface{} {
	var flattened []interface{}
	for _, subarr := range arr {
		flattened = append(flattened, subarr...)
	}
	return flattened
}

func main() {
	computeDiskChecksum()
	calculateOptimisedChecksum()
}
