package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type LottoTicket struct {
	numbers         []int
	winningNumbers  []int
	score           int
	matchingNumbers int
	copies          int
}

func main() {
	numberRegex := regexp.MustCompile(`\d+`)

	tickets := []LottoTicket{}

	file, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(string(file), "\n")

	for _, line := range lines {

		_, data, _ := strings.Cut(line, ":")
		splitData := strings.Split(data, "|")

		numberStrings := numberRegex.FindAll([]byte(splitData[0]), -1)
		winningNumberStrings := numberRegex.FindAll([]byte(splitData[1]), -1)

		numbers, winningNumbers := []int{}, []int{}

		for _, n := range numberStrings {
			value, _ := strconv.Atoi(string(n))
			numbers = append(numbers, value)
		}
		for _, n := range winningNumberStrings {
			value, _ := strconv.Atoi(string(n))
			winningNumbers = append(winningNumbers, value)
		}

		tickets = append(tickets, LottoTicket{
			numbers:         numbers,
			winningNumbers:  winningNumbers,
			score:           0,
			matchingNumbers: 0,
			copies:          1,
		})
	}

	sum := 0
	for i := range tickets {
		ticket := &tickets[i]
		for _, n := range ticket.numbers {
			if slices.Contains(ticket.winningNumbers, n) {
				ticket.matchingNumbers = ticket.matchingNumbers + 1
				if ticket.score == 0 {
					ticket.score = 1
				} else {
					ticket.score *= 2
				}
			}
		}
		sum += ticket.score
	}
	fmt.Println("Part1:", sum)

	for index, ticket := range tickets {
		for x := 0; x < ticket.copies; x++ {
			for i := index + 1; i <= min(index+ticket.matchingNumbers, len(tickets)-1); i++ {
				tickets[i].copies += 1
			}
		}
	}

	sum = 0
	for _, ticket := range tickets {
		sum += ticket.copies
	}

	fmt.Println("Part2:", sum)
}
