package main

import (
	"fmt"
	"sort"
	"strings"
)

func FindAnagrams(words []string) map[string][]string {
	lwords := make([]string, len(words))
	for i, word := range words {
		lwords[i] = strings.ToLower(word)
	}

	group := make(map[string][]string)
	for _, word := range lwords {
		sorted := sortString(word)
		group[sorted] = append(group[sorted], word)
	}

	result := make(map[string][]string)

	for _, group := range group {

		uniqueSorted := Sort(group)

		key := uniqueSorted[0]
		result[key] = uniqueSorted
	}

	return result
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func Sort(words []string) []string {
	set := make(map[string]bool)
	unique := []string{}

	for _, word := range words {
		if !set[word] {
			set[word] = true
			unique = append(unique, word)
		}
	}

	sort.Strings(unique)
	return unique
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	anagrams := FindAnagrams(words)

	for key, group := range anagrams {
		fmt.Println(key, group)
	}
}
