package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type GameSet struct {
	gameNumber int
	games      []map[string]int
}

var VALID_COLORS = []string{"red", "blue", "green"}
var MAX_VALUES = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	file, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(string(file), "\n")
	Matches := []GameSet{}

	parseGameNumber := regexp.MustCompile(`\d+$`)
	for _, line := range lines {
		Game := GameSet{}

		// Parse game number from the rest of the line
		parts := strings.Split(line, ": ")
		firstPart, secondPart := parts[0], parts[1]
		Game.gameNumber, _ = strconv.Atoi(parseGameNumber.FindString(firstPart))

		// Parse colors in each line
		rounds := strings.Split(secondPart, "; ")
		for _, round := range rounds {
			sets := strings.Split(round, ", ")
			gameInfo := map[string]int{}

			// Separate color and amount
			for _, color := range sets {
				info := strings.Split(color, " ")
				color := info[1]
				amount, _ := strconv.Atoi(info[0])
				gameInfo[color] = amount
			}

			// Store game info
			Game.games = append(Game.games, gameInfo)
		}
		Matches = append(Matches, Game)
	}

	// Parse valid matches
	part1sum := 0
	for _, match := range Matches {
		part1sum += validMatch(match)
	}
	part2sum := 0

	for _, match := range Matches {
		part2sum += cubePower(match)
	}
	fmt.Println("part1: ", part1sum)
	fmt.Println("part2: ", part2sum)
}

func validMatch(game GameSet) int {
	for _, round := range game.games {
		for key, value := range round {
			// Exit if color is not valid or exceeds the max per key
			if !slices.Contains(VALID_COLORS, key) || value > MAX_VALUES[key] {
				return 0
			}
		}
	}
	return game.gameNumber
}

func cubePower(game GameSet) int {
	min_cubes := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, round := range game.games {
		for key, value := range round {
			// Exit if color is not valid or exceeds the max per key
			if value > min_cubes[key] {
				min_cubes[key] = value
			}
		}
	}
	return min_cubes["red"] * min_cubes["green"] * min_cubes["blue"]
}
