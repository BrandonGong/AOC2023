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

	var pt1Sum int
	for _, line := range lines {
		numbers, card := ParseLine(line)
		power := -1.0
		for _, number := range numbers {
			if slices.Contains(card, number) {
				power++
			}
		}
		if power > -1 {
			pt1Sum += int(math.Pow(2, power))
		}
	}
	fmt.Printf("Part 1: Total points %d\n", pt1Sum)
}

func ParseLine(line string) (numbers []string, card []string) {
	re := regexp.MustCompile(`\d+`)

	start := strings.Index(line, ":")
	parts := strings.Split(line[start:], "|")
	numbers = re.FindAllString(parts[0], -1)
	card = re.FindAllString(parts[1], -1)
	return numbers, card
}
