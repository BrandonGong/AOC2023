package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var searchStrings []string

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	searchStrings = []string{"\\d", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	re := regexp.MustCompile(strings.Join(searchStrings, "|"))
	fs := bufio.NewScanner(input)
	sum := 0
	for fs.Scan() {
		line := fs.Text()
		// Handle edge cases where two digits are merged together
		line = strings.ReplaceAll(line, "eightwo", "82")
		line = strings.ReplaceAll(line, "twone", "21")
		line = strings.ReplaceAll(line, "oneight", "18")
		// Find all matches of search strings in line
		numbers := re.FindAllString(line, -1)
		first := ParseNumber(numbers[0])
		last := ParseNumber(numbers[len(numbers)-1])
		// Make calibration value as two digit number from first and last digits found
		calibrationValue, _ := strconv.Atoi(fmt.Sprintf("%d%d", first, last))

		//fmt.Printf("In Line: %s, Found numbers: %v, Calibration Number: %d\n", line, numbers, calibrationValue)
		//fmt.Println(calibrationValue)
		sum += calibrationValue
	}
	fmt.Printf("Total Value: %d\n", sum)
}

func ParseNumber(value string) int {
	num, err := strconv.Atoi(value)
	if err == nil {
		return num
	}
	for i, s := range searchStrings {
		if s == value {
			return i
		}
	}
	panic("Value not found: " + value)
}
