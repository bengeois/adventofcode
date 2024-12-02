package main

import (
	"bufio"
	"flag"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getFilePath() string {
	inputFile := flag.String("file", "input.txt", "Input file")

	flag.Parse()

	return *inputFile
}
func getFileContent(filePath string) ([]int, []int) {
	var column1 []int
	var column2 []int

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
		var columns = strings.Fields(line)
		
		// Convert the string to int
		val1, err := strconv.Atoi(columns[0])
		if err != nil {
			panic(err)
		}
		val2, err := strconv.Atoi(columns[1])
		if err != nil {
			panic(err)
		}
		
		column1 = append(column1, val1)
		column2 = append(column2, val2)
	}

	return column1, column2
}

func absDiffInt(x, y int) int {
	if x < y {
	   return y - x
	}
	return x - y
}

func countValueOccurence(arr []int) map[int]int {
	// Create a dictionary to store the occurence of each value
    dict:= make(map[int]int)
    for _ , num :=  range arr {
        dict[num] = dict[num]+1
    }
	return dict
}

func main() {
	inputFilePath := getFilePath()

	// Get the content of the file
	column1, column2 := getFileContent(inputFilePath)

	// Sort the columns
	sort.Ints(column1)
	sort.Ints(column2)

	var finalDistance int
	var similarityScore int
	occurence := countValueOccurence(column2)

	for i := 0; i < len(column1); i++ {
		finalDistance += absDiffInt(column1[i], column2[i])
		similarityScore += column1[i] * occurence[column1[i]]
	}

	// Print the final distance
	println("Total distance:", finalDistance)
	println("Similarity score:", similarityScore)
}