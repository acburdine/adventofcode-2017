package main

import (
	"fmt"

	"github.com/acburdine/adventofcode-2017"
)

const (
	openGroup  = '{'
	closeGroup = '}'

	openGarbage  = '<'
	closeGarbage = '>'

	ignore = '!'
)

func main() {
	stream := []rune(adventofcode.Input()[0])

	part1, part2 := score(stream)

	fmt.Printf("Day 9 Part 1: %d\n", part1)
	fmt.Printf("Day 9 Part 2: %d\n", part2)
}

func score(stream []rune) (int, int) {
	sum := 0
	garbageChars := 0
	level := 0
	garbage := false

	for i := 0; i < len(stream); i++ {
		char := stream[i]

		if garbage {
			if char == ignore {
				i++
			} else if char == closeGarbage {
				garbage = false
			} else {
				garbageChars++
			}
		} else if char == openGarbage {
			garbage = true
		} else if char == openGroup {
			level++
		} else if char == closeGroup && level > 0 {
			sum += level
			level--
		}
	}

	return sum, garbageChars
}
