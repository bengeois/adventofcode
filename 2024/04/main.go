package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func getFilePath() string {
	inputFile := flag.String("file", "input.txt", "Input file")

	flag.Parse()

	return *inputFile
}

func getFileContent(filePath string) []string {
	var content []string

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
		content = append(content, line)
	}

	return content
}

func findOccurences(content []string) int {
	word := "XMAS"
	rows := len(content)
	cols := len(content[0])
	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{-1, -1},
		{1, -1},
		{-1, 1},
	}

	wordLen := len(word)
	gridRunes := make([][]rune, rows)
	for i := 0; i < rows; i++ {
		gridRunes[i] = []rune(content[i])
	}

	count := 0

	// Helper function to check if a word matches in a given direction
	checkDirection := func(x, y, dx, dy int) bool {
		for k := 0; k < wordLen; k++ {
			nx, ny := x+k*dx, y+k*dy
			if nx < 0 || nx >= rows || ny < 0 || ny >= cols || gridRunes[nx][ny] != rune(word[k]) {
				return false
			}
		}
		return true
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// Check all directions from this cell
			for _, dir := range directions {
				dx, dy := dir[0], dir[1]
				if checkDirection(i, j, dx, dy) {
					count++
				}
			}
		}
	}

	return count
}

func findXmasFigure(content []string) int {
	rows := len(content)
	cols := len(content[0])
	count := 0

	isXMAS := func(x, y int) bool {
		if x-1 < 0 || x+1 >= rows || y-1 < 0 || y+1 >= cols {
			return false
		}

		// Extract diagonals
		topLeft := string(content[x-1][y-1])
		topRight := string(content[x-1][y+1])
		center := string(content[x][y])
		bottomLeft := string(content[x+1][y-1])
		bottomRight := string(content[x+1][y+1])

		validDiagonal1 := (topLeft == "M" && center == "A" && bottomRight == "S") || (topLeft == "S" && center == "A" && bottomRight == "M")
		validDiagonal2 := (topRight == "M" && center == "A" && bottomLeft == "S") || (topRight == "S" && center == "A" && bottomLeft == "M")

		return validDiagonal1 && validDiagonal2
	}
	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if isXMAS(i, j) {
				count++
			}
		}
	}

	return count
}

func main() {
	inputFilePath := getFilePath()

	content := getFileContent(inputFilePath)

	count := findOccurences(content)
	countFigure := findXmasFigure(content)

	fmt.Printf("Count: %d\n", count)
	fmt.Printf("Figure: %d\n", countFigure)
}
