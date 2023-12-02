package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, _ := os.ReadFile(os.Args[1])

	digitOnly := regexp.MustCompile(`\d`)
	spelledDigit := regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine|[1-9]`)

	lines := strings.Split(string(file), "\n")

	findValue := func(reg regexp.Regexp, lines []string) int {
		sum := 0
		for _, line := range lines {
			matches := reg.FindAllString(line, -1)

			firstMatch := matchValue(matches[0])
			lastMatch := matchValue(matches[len(matches)-1])

			value := firstMatch*10 + lastMatch
			sum += value
		}
		return sum
	}

	fmt.Println("Part 1:", findValue(*digitOnly, lines))
	fmt.Println("Part 2:", findValue(*spelledDigit, lines))
}

func matchValue(s string) int {
	if len(s) == 1 {
		return int(s[0] - '0')
	}

	switch s {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		return 0
	}
}
