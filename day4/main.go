package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/thoas/go-funk"

	"github.com/acburdine/adventofcode-2017"
)

func main() {
	lines := adventofcode.Input()

	fmt.Printf("Day 4 Part 1: %d\n", part1(lines))
	fmt.Printf("Day 4 Part 2: %d\n", part2(lines))
}

func part1(lines []string) int {
	sum := 0

	for _, l := range lines {
		words := strings.Split(strings.TrimSpace(l), " ")

		if len(funk.Uniq(words).([]string)) == len(words) {
			sum++
		}
	}

	return sum
}

func part2(lines []string) int {
	sum := 0

	for _, l := range lines {
		words := sortWords(strings.Split(strings.TrimSpace(l), " "))

		if len(funk.Uniq(words).([]string)) == len(words) {
			sum++
		}
	}

	return sum
}

type runeSort []rune

func (s runeSort) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s runeSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s runeSort) Len() int {
	return len(s)
}

func sortWords(arr []string) []string {
	result := make([]string, len(arr))
	for i, s := range arr {
		r := []rune(s)
		sort.Sort(runeSort(r))
		result[i] = string(r)
	}
	return result
}
