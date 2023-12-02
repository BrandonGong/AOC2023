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
	var power int
	games := ParseInput()
	for gameId, results := range games {
		if CheckPossible(results, query) {
			sum += gameId
		}
		product := 1
		for _, count := range results {
			product *= count
		}
		power += product
	}
	fmt.Printf("Part 1: Sum of possible Game IDs: %d\n", sum)
	fmt.Printf("Part 2: Power of all games: %d\n", power)
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
