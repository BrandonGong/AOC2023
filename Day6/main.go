package main

import (
	shared "aoc2023/Shared"
	"fmt"
)

func main() {
	lines := shared.ReadInputFile("input.txt")
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
