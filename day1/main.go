package main

import (
	"fmt"

	"github.com/acburdine/adventofcode-2017"
)

func main() {
	line := adventofcode.Input()[0]
	ints := mapInt([]rune(line))

	fmt.Printf("Day 1 Part 1: %d\n", part1(ints))
	fmt.Printf("Day 1 Part 2: %d\n", part2(ints))
}

func mapInt(arr []rune) []int {
	ints := make([]int, len(arr))
	for i, r := range arr {
		ints[i] = int(r - '0')
	}
	return ints
}

func part1(ints []int) int {
	sum := 0
	for i, v := range ints {
		next := (i + 1) % len(ints)
		if v == ints[next] {
			sum += v
		}
	}
	return sum
}

func part2(ints []int) int {
	inc := len(ints) / 2
	sum := 0
	for i, v := range ints {
		next := (i + inc) % len(ints)
		if v == ints[next] {
			sum += v
		}
	}
	return sum
}
