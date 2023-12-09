package main

import (
	shared "aoc2023/Shared"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := shared.ReadInputFile("input.txt")
	Part1(lines)
	Part2(lines)
}
func Part1(lines []string) {
	times := shared.ParseIntsFromLine(lines[0])
	recordDistances := shared.ParseIntsFromLine(lines[1])
	part1Product := 1
	for i := 0; i < len(times); i++ {
		winOptions := 0
		for t := 1; t < times[i]; t++ {

			distance := t * (times[i] - t)
			if distance > recordDistances[i] {
				winOptions++
			}
		}
		part1Product *= winOptions
	}
	fmt.Printf("Part1: %d\n", part1Product)
}
func Part2(lines []string) {
	time, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", ""))
	recordDistance, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", ""))
	winOptions := 0
	for t := 1; t < time; t++ {
		distance := t * (time - t)
		if distance > recordDistance {
			winOptions++
		}
	}
	fmt.Printf("Part2: %d\n", winOptions)
}
