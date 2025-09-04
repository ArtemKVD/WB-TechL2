package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

type config struct {
	After      int
	Before     int
	Context    int
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
}

func ReadFile() []string {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Print("error open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text
}

func LineMatch(check, line string, cfg *config) bool {
	if cfg.IgnoreCase {
		check = strings.ToLower(check)
		line = strings.ToLower(line)
	}

	if cfg.Fixed {
		return strings.Contains(line, check)
	}

	flag, err := regexp.MatchString(check, line)
	if err != nil {
		log.Fatal("failed match line", err)
	}
	return flag
}

func main() {
	cfg := &config{}

	flag.IntVar(&cfg.After, "A", 0, "")
	flag.IntVar(&cfg.Before, "B", 0, "")
	flag.IntVar(&cfg.Context, "C", 0, "")
	flag.BoolVar(&cfg.Count, "c", false, "")
	flag.BoolVar(&cfg.IgnoreCase, "i", false, "")
	flag.BoolVar(&cfg.Invert, "v", false, "")
	flag.BoolVar(&cfg.Fixed, "F", false, "")
	flag.BoolVar(&cfg.LineNum, "n", false, "")

	flag.Parse()

	arguments := flag.Args()

	searchPattern := arguments[0]

	lines := ReadFile()

	var indexs []int
	for i, line := range lines {
		isMatch := LineMatch(searchPattern, line, cfg)

		if cfg.Invert {
			isMatch = !isMatch
		}

		if isMatch {
			indexs = append(indexs, i)
		}
	}

	if cfg.Count {
		fmt.Println(len(indexs))
		return
	}

	if cfg.Context > 0 {
		cfg.Before = cfg.Context
		cfg.After = cfg.Context
	}

	outIndexs := make(map[int]struct{})
	for _, i := range indexs {
		outIndexs[i] = struct{}{}

		if cfg.Before > 0 {
			start := max(0, i-cfg.Before)
			for j := start; j < i; j++ {
				outIndexs[j] = struct{}{}
			}
		}

		if cfg.After > 0 {
			end := min(len(lines)-1, i+cfg.After)
			for j := i + 1; j <= end; j++ {
				outIndexs[j] = struct{}{}
			}
		}
	}

	var result []int
	for i := range outIndexs {
		result = append(result, i)
	}
	sort.Ints(result)

	for _, i := range result {
		if cfg.LineNum {
			fmt.Printf("%d:", i+1)
		}
		fmt.Println(lines[i])
	}
}
