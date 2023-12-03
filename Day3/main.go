package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type PartNumber struct {
	Value   int
	Start   int
	End     int
	Counted bool
}
type Symbol struct {
	Value    string
	Position int
}

func main() {
	numberMap, symbolMap, maxLines := ParseInput("input.txt")
	var sum int
	// Find numbers that are adjacent to a symbol
	for line := 0; line <= maxLines; line++ {
		searchStart := max(0, line-1)
		searchEnd := min(maxLines, line+1)
		for _, symbol := range symbolMap[line] {
			// Search lines above and below current line
			for i := searchStart; i <= searchEnd; i++ {
				for _, number := range numberMap[i] {
					if number.Start-1 <= symbol.Position &&
						number.End >= symbol.Position &&
						!number.Counted {
						// number is adjacent
						sum += number.Value
						number.Counted = true
						fmt.Printf("[%d] %s\t[%d] %d\n", line, symbol.Value, i, number.Value)
					}
				}
			}
		}
	}
	fmt.Printf("Part 1: Sum of part numbers %d\n", sum)
}

func ParseInput(file string) (map[int][]*PartNumber, map[int][]*Symbol, int) {
	allNumbers := make(map[int][]*PartNumber)
	allSymbols := make(map[int][]*Symbol)
	inputFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
	fs := bufio.NewScanner(inputFile)
	reNumbers := regexp.MustCompile(`\d+`)
	reSymbols := regexp.MustCompile(`[^\w\d\.]`)
	lineIndex := 0
	for fs.Scan() {
		line := fs.Text()
		// Parse part numbers on line
		nums := reNumbers.FindAllStringIndex(line, -1)
		if nums != nil {
			lineNumbers := make([]*PartNumber, len(nums))
			for i, match := range nums {
				value, err := strconv.Atoi(line[match[0]:match[1]])
				if err != nil {
					panic(err)
				}
				partNumber := PartNumber{
					Start: match[0],
					End:   match[1],
					Value: value,
				}
				lineNumbers[i] = &partNumber
			}
			allNumbers[lineIndex] = lineNumbers
		}
		// Parse symbos from line
		symbols := reSymbols.FindAllStringIndex(line, -1)
		if symbols != nil {
			lineSymbols := make([]*Symbol, len(symbols))
			for i, match := range symbols {
				value := line[match[0]:match[1]]
				symbol := Symbol{
					Value:    value,
					Position: match[0],
				}
				lineSymbols[i] = &symbol
			}
			allSymbols[lineIndex] = lineSymbols
		}
		lineIndex++
	}
	return allNumbers, allSymbols, lineIndex
}
