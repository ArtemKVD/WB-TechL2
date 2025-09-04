package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fieldf := flag.String("f", "", "fields")
	delimiter := flag.String("d", "\t", "delimiter")
	separated := flag.Bool("s", false, "separated")

	flag.Parse()

	var Fields []int
	if *fieldf != "" {
		parts := strings.Split(*fieldf, ",")
		for _, part := range parts {
			if strings.Contains(part, "-") {
				rangeParts := strings.Split(part, "-")
				if len(rangeParts) == 2 {
					start, _ := strconv.Atoi(rangeParts[0])
					end, _ := strconv.Atoi(rangeParts[1])
					for i := start; i <= end; i++ {
						Fields = append(Fields, i)
					}
				}
			} else {
				field, _ := strconv.Atoi(part)
				Fields = append(Fields, field)
			}
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if *separated && !strings.Contains(line, *delimiter) {
			continue
		}

		columns := strings.Split(line, *delimiter)

		if len(Fields) == 0 {
			fmt.Println(line)
			continue
		}

		var result []string
		for _, field := range Fields {
			if field > 0 && field <= len(columns) {
				result = append(result, columns[field-1])
			}
		}

		fmt.Println(strings.Join(result, *delimiter))
	}
}
