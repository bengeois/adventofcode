package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getFilePath() string {
	inputFile := flag.String("file", "input.txt", "Input file")

	flag.Parse()

	return *inputFile
}

func isDampenerMode() bool {
	dampenerMode := flag.Bool("dampener", false, "Dampener mode")

	flag.Parse()

	return *dampenerMode
}

func getFileContent(filePath string) ([][]int) {
	var content [][]int

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	// Close the file
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		var parsedLine = strings.Fields(line)

		var values []int
		for _, val := range parsedLine {
			column, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			values = append(values, column)
		}
		content = append(content, values)
	}

	return content
}

func absDiffInt(x, y int) int {
	if x < y {
	   return y - x
	}
	return x - y
}

func checkLineIncrease(line []int) bool {
	// Check if the values of a line are increasing
	for i := 0; i < len(line) - 1; i++ {
		if line[i] >= line[i + 1] {
			return false
		}
	}
	return true
}

func checkLineDecrease(line []int) bool {
	// Check if the values of a line are decreasing
	for i := 0; i < len(line) - 1; i++ {
		if line[i] <= line[i + 1] {
			return false
		}
	}
	return true
}

func checkLineDiff(line []int) bool {
	// Check if the difference between the values of a line is between 1 and 3
	for i := 0; i < len(line) - 1; i++ {
		if absDiffInt(line[i], line[i + 1]) < 1 || absDiffInt(line[i], line[i + 1]) > 3 {
			return false
		}
	}
	return true
}

func isSafe(line []int) bool {
	increasing := checkLineIncrease(line)
	decreasing := checkLineDecrease(line)
	diff := checkLineDiff(line)

	return (increasing || decreasing) && diff
}

func isSafeDampener(line []int) bool {
	// For each value in the line
	// Try to remove it and check if the line is still safe without it
	for i := 0; i < len(line); i++ {
		var newLine []int
		newLine = append(newLine, line[:i]...)
		newLine = append(newLine, line[i+1:]...)

		if isSafe(newLine) {
			return true
		}
	}
	return false
}

func countSafeReport(content [][]int, dampenerMode bool) int {
	var safeReport int
	for _, line := range content {
		if dampenerMode {
			if (isSafeDampener(line)) {
				safeReport++
			}
		} else {
			if isSafe(line) {
				safeReport++
			}
		}
	}
	return safeReport
}

func main() {
	dampenerMode := isDampenerMode()
	inputFilePath := getFilePath()

	// Get the content of the file
	content := getFileContent(inputFilePath)

	safeReport := countSafeReport(content, dampenerMode)

	fmt.Print("Safe report: ", safeReport)
}