package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/acburdine/adventofcode-2017"
)

func main() {
	inst := mapInt(adventofcode.Input())

	fmt.Printf("Day 5 Part 1: %d\n", part1(inst))
	fmt.Printf("Day 5 Part 2: %d\n", part2(inst))
}

func mapInt(arr []string) []int {
	result := make([]int, len(arr))
	for i, s := range arr {
		num, _ := strconv.Atoi(strings.TrimSpace(s))
		result[i] = num
	}
	return result
}

func part1(inst []int) int {
	list := make([]int, len(inst))
	copy(list, inst)

	steps, i := 0, 0
	for i < len(list) && i >= 0 {
		i, list[i] = i+list[i], list[i]+1
		steps++
	}

	return steps
}

func part2(inst []int) int {
	list := make([]int, len(inst))
	copy(list, inst)

	steps, i := 0, 0
	for i < len(list) && i >= 0 {
		jump := list[i]
		if jump >= 3 {
			list[i]--
		} else {
			list[i]++
		}
		i += jump
		steps++
	}

	return steps
}
