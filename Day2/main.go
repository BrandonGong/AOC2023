package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	query := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	var sum int
	games := ParseInput()
	for gameId, results := range games {
		if CheckPossible(results, query) {
			sum += gameId
		}
	}
	fmt.Printf("Sum of Game IDs: %d\n", sum)
}

func ParseInput() map[int]map[string]int {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err) // Could not open file
	}
	var games = make(map[int]map[string]int)
	fs := bufio.NewScanner(input)
	for fs.Scan() {
		line := fs.Text()

		// Line is in format: "Game ##: # color1, # color2; # color1, # color2, # color3"
		// Game number separated by :, colors drawn, with rounds separated by ;
		var gameId int
		var color string
		var count int

		parts := strings.Split(line, ":")
		fmt.Sscanf(parts[0], "Game %d", &gameId) // Get game ID
		// Split games
		var results = make(map[string]int)
		rounds := strings.Split(parts[1], ";")
		for _, round := range rounds {
			draws := strings.Split(round, ",")
			for _, draw := range draws {
				draw = strings.Trim(draw, " ")
				fmt.Sscanf(draw, "%d %s", &count, &color)
				if count > results[color] {
					results[color] = count
				}
			}
		}

		games[gameId] = results
	}
	return games
}
func CheckPossible(game map[string]int, query map[string]int) bool {
	for color, number := range query {
		if game[color] > number {
			return false
		}
	}
	return true
}
