package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(string(file), "\n")

	
	rows := len(lines)
	cols := len(lines[0])

	findNumber := regexp.MustCompile(`\d+`)
	findSymbol := regexp.MustCompile(`[^.\d]`)
	findGears := regexp.MustCompile(`\*`)

	symbolLocations := make([][]bool, rows)
	for i := range lines {
		symbolLocations[i] = make([]bool, cols)
	}

	// find all symbols
	for lineNumber, line := range lines {
		// indexes are returned as an array of (first, last)
		foundSymbols := findSymbol.FindAllIndex([]byte(line), -1)
		for _, s := range foundSymbols {
			symbolLocations[lineNumber][s[0]] = true
		}
	}

	// find all numbers
	partNumbers := 0
	for lineNumber, line := range lines {
		numberLocations := findNumber.FindAllIndex([]byte(line), -1)
		if numberLocations == nil {
			continue
		}

		for _, locations := range numberLocations {
			start, end := locations[0], locations[1]
			validNumber := false

			// Search the box (including diagonals) around the number
			for x := max(lineNumber-1, 0); x <= min(rows-1, lineNumber+1); x++ {
				if slices.Contains(symbolLocations[x][max(start-1, 0):min(end+1, cols-1)], true) {
					validNumber = true
					continue
				}
			}

			// If valid number, parse integer
			if validNumber {
				value, _ := strconv.Atoi(line[start:end])
				partNumbers += value
			}
		}
	}
	fmt.Println("Part 1:", partNumbers)

	gearRatioSum := 0
	for lineNumber, line := range lines {
		gears := findGears.FindAllIndex([]byte(line), -1)

		allPartNumbers := map[int][][]int{} // cursed

		for x := max(lineNumber-1, 0); x <= min(lineNumber+1, rows-1); x++ {
			allPartNumbers[x] = findNumber.FindAllStringIndex(lines[x], -1)
		}

		for _, gear := range gears {
			goldenPartNumbers := []GoldenDigit{}
			gearLocation := []int{max(gear[0]-1, 0), min(gear[0]+1, rows-1)}
			// Search the box (including diagonals) around the number
			for x := max(lineNumber-1, 0); x <= min(rows-1, lineNumber+1); x++ {
				for _, digit := range allPartNumbers[x] {
					if rangesOverlap(digit, gearLocation) {
						goldenPartNumbers = append(goldenPartNumbers, GoldenDigit{row: x, digitRange: digit})
					}
				}
			}

			if len(goldenPartNumbers) == 2 {
				part1 := goldenPartNumbers[0]
				part2 := goldenPartNumbers[1]
				value1, _ := strconv.Atoi(lines[part1.row][part1.digitRange[0]:part1.digitRange[1]])
				value2, _ := strconv.Atoi(lines[part2.row][part2.digitRange[0]:part2.digitRange[1]])
				gearRatioSum += value1 * value2
			}
		}
	}

	fmt.Println("Part 2:", gearRatioSum)

}

type GoldenDigit struct {
	row        int
	digitRange []int
}

func rangesOverlap(digit, gear []int) bool {
	// We minus one from the digit range here because its needed for slicing but not for actual range
	return digit[0] <= gear[1] && digit[1]-1 >= gear[0]
}
