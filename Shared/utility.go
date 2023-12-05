package shared

import (
	"bufio"
	"os"
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
