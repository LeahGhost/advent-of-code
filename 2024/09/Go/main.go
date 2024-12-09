package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func computeDiskChecksum() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()

	processingFile := true
	currentFileIndex := 0
	diskLayout := []interface{}{} 

	for _, sizeChar := range input {
		blockSize, err := strconv.Atoi(string(sizeChar))
		if err != nil {
			fmt.Println("Error converting block size:", err)
			return
		}

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
	checksum := 0
	for i, block := range diskLayout {
		if block != "." {
			checksum += block.(int) * i
		}
	}
	fmt.Println("Checksum:", checksum)
}

func main() {
	computeDiskChecksum()
}
