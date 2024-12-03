package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getFilePath() string {
	inputFile := flag.String("file", "input.txt", "Input file")

	flag.Parse()

	return *inputFile
}

func getFileContent(filePath string) string {
	var content string

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
		content = content + line
	}

	return content
}

func getAllPatterns(content string) []string {
	// Get all the patterns
	// Regex pattern
	pattern := `mul\(\d{1,3},\d{1,3}\)`

	// Compile the regex
	re := regexp.MustCompile(pattern)
	
	// Find all matches
	matches := re.FindAllString(content, -1)

	return matches
}

func getNumbers(content string) []int {
	// Get the numbers
	// Regex pattern
	pattern := `\d{1,3}`

	// Compile the regex
	re := regexp.MustCompile(pattern)
	
	// Find all matches
	matches := re.FindAllString(content, -1)

	var numbers []int
	for _, match := range matches {
		number, err := strconv.Atoi(match)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}

	return numbers
}

func calculateTotal(patterns []string) int {
	total := 0

	for _, pattern := range patterns {
		// Get the numbers
		numbers := getNumbers(pattern)

		// Multiply the numbers
		result := numbers[0] * numbers[1]

		// Add the result to the total
		total += result
	}

	return total
}

func getMode() bool {
	mode := flag.Bool("domode", false, "Mode")

	flag.Parse()

	return *mode
}

func cleanDoMode(content string) string {
	doContent := strings.Split(content, "do()")
	var cleanContent string

    for i := range doContent {
		cleanPart := strings.Split(doContent[i], "don't()")
        cleanContent = cleanContent + cleanPart[0]
    }

	return cleanContent
}

func main() {
	doMode := getMode()
	inputFilePath := getFilePath()

	// Get the content of the file
	content := getFileContent(inputFilePath)

	if doMode {
		content = cleanDoMode(content)
	}

	patterns := getAllPatterns(content)

	total := calculateTotal(patterns)

	fmt.Print("Total: ", total)
}