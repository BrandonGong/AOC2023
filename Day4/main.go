package main

import (
	shared "aoc2023/Shared"
	"fmt"
	"math"
	"regexp"
	"slices"
	"strings"
)

func main() {
	lines := shared.ReadInputFile("input.txt")
	// Part 1: Get the total number of points
	var pt1Sum int
	for _, line := range lines {
		numbers, card := ParseLine(line)
		wins := 0
		for _, number := range numbers {
			if slices.Contains(card, number) {
				wins++

			}
		}
		if wins > 0 {
			pt1Sum += int(math.Pow(2, float64(wins)-1))
		}
	}
	fmt.Printf("Part 1: Total points %d\n", pt1Sum)
	// Part 2: Winning cards return copies of cards equal to that number of points
	// How many total scratch cards do you end up with?
	var totalCopies int
	for i := 0; i < len(lines); i++ {
		totalCopies += GetCopies(lines, i)
	}
	fmt.Printf("Part 2: Total number of cards: %d\n", totalCopies)
}

func ParseLine(line string) (numbers []int, card []int) {
	re := regexp.MustCompile(`\d+`)

	start := strings.Index(line, ":")
	parts := strings.Split(line[start:], "|")
	numbers = shared.StringsToInts(re.FindAllString(parts[0], -1))
	card = shared.StringsToInts(re.FindAllString(parts[1], -1))
	return numbers, card
}

func GetCopies(lines []string, gameNumber int) int {
	numbers, card := ParseLine(lines[gameNumber])
	total := 1
	wins := 0
	for _, number := range numbers {
		if slices.Contains(card, number) {
			wins++
			total += GetCopies(lines, (gameNumber + wins))
		}
	}
	return total
}
