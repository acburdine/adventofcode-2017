package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/acburdine/adventofcode-2017"
)

var regex = regexp.MustCompile("([0-9]+)")

func main() {
	initial := regex.FindAllString(strings.TrimSpace(adventofcode.Input()[0]), -1)
	banks := mapInt(initial)

	part1, part2 := run(banks)

	fmt.Printf("Day 6 Part 1: %d\n", part1)
	fmt.Printf("Day 6 Part 2: %d\n", part2)
}

func mapInt(arr []string) []int {
	result := make([]int, len(arr))
	for i, s := range arr {
		num, _ := strconv.Atoi(s)
		result[i] = num
	}
	return result
}

func join(arr []int) string {
	result := make([]string, len(arr))
	for i, x := range arr {
		result[i] = strconv.Itoa(x)
	}
	return strings.Join(result, "|")
}

func maxPos(arr []int) (int, int) {
	max, pos := 0, 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max, pos = arr[i], i
		}
	}
	return max, pos
}

func run(banks []int) (int, int) {
	found := make(map[string]int)
	found[join(banks)] = 1

	for cycles := 1; true; cycles++ {
		max, pos := maxPos(banks)
		banks[pos] = 0

		for i := 0; i < max; i++ {
			banks[(pos+1+i)%len(banks)] += 1
		}

		key := join(banks)
		i, ok := found[key]
		if ok {
			return cycles, cycles - i
		}
		found[key] = cycles
	}
	return 0, 0
}
