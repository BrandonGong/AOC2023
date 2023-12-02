package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile(`\d`)
	fs := bufio.NewScanner(input)
	sum := 0
	for fs.Scan() {
		line := fs.Text()
		numbers := re.FindAllString(line, -1)
		calibrationValue, _ := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])

		fmt.Printf("In Line: %s, Found numbers: %v, Calibration Number: %d\n", line, numbers, calibrationValue)
		sum += calibrationValue
	}
	fmt.Printf("Total Value: %d\n", sum)
}
