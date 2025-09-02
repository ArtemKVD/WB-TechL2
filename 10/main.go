package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	k := flag.Int("k", 0, "column number to sort")
	n := flag.Bool("n", false, "value sort")
	r := flag.Bool("r", false, "reverse sort")
	u := flag.Bool("u", false, "unique")
	flag.Parse()

	lines := readLines("file.txt")

	if *u {
		lines = removeDuplicates(lines)
	}

	sortLines(lines, *k, *n, *r)

	for _, line := range lines {
		fmt.Println(line)
	}
}

func readLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func removeDuplicates(lines []string) []string {
	set := make(map[string]bool)
	result := []string{}

	for _, line := range lines {
		if !set[line] {
			set[line] = true
			result = append(result, line)
		}
	}

	return result
}

func sortLines(lines []string, k int, n bool, r bool) {
	if k > 0 {
		k--
	}

	sort.Slice(lines, func(i, j int) bool {
		var keyI, keyJ string

		if k >= 0 {
			keyI = getKey(lines[i], k)
			keyJ = getKey(lines[j], k)
		} else {
			keyI = lines[i]
			keyJ = lines[j]
		}

		var result bool

		if n {
			numI, errI := strconv.Atoi(strings.TrimSpace(keyI))
			numJ, errJ := strconv.Atoi(strings.TrimSpace(keyJ))

			if errI == nil && errJ == nil {
				result = numI < numJ
			} else if errI != nil && errJ != nil {
				result = keyI < keyJ
			} else {
				result = errI == nil
			}
		} else {
			result = keyI < keyJ
		}

		if r {
			return !result
		}
		return result
	})
}

func getKey(line string, i int) string {
	columns := strings.Fields(line)
	if i < len(columns) {
		return columns[i]
	}
	return ""
}
