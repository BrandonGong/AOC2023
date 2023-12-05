package shared

import (
	"bufio"
	"os"
	"strconv"
)

func ReadInputFile(file string) []string {
	lines := make([]string, 0)

	fs, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fs.Close()

	sr := bufio.NewScanner(fs)
	for sr.Scan() {
		line := sr.Text()
		lines = append(lines, line)
	}
	return lines
}
func StringsToInts(s []string) []int {
	ints := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		value, _ := strconv.Atoi(s[i])
		ints[i] = value
	}
	return ints
}
