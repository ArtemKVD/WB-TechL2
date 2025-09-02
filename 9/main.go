package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

func unpacking(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	flag := true
	for _, char := range s {
		if !unicode.IsDigit(char) {
			flag = false
			break
		}
	}
	if flag {
		return "", errors.New("only digits")
	}

	var result string
	i := 0

	for i < len(s)-1 {
		if unicode.IsLetter(rune(s[i])) {
			char := s[i]
			count := 1
			if unicode.IsDigit(rune(s[i+1])) {
				num, _ := strconv.Atoi(string(s[i+1]))
				count = num
				i++
			}
			for j := 0; j < count; j++ {
				result += string(char)
			}
		}
		i++
	}
	if unicode.IsLetter(rune(s[len(s)-1])) {
		result += string(s[len(s)-1])
	}
	return result, nil
}

func main() {
	s := ""
	str, err := unpacking(s)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(str)
}
