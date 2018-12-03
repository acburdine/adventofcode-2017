package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/acburdine/adventofcode-2017"
)

var regex = regexp.MustCompile("([0-9]+)")

func main() {
	sheet := mapLines(adventofcode.Input())

	fmt.Printf("Day 2 Part 1: %d\n", part1(sheet))
	fmt.Printf("Day 2 Part 2: %d\n", part2(sheet))
}

func mapLines(lines []string) [][]int {
	result := make([][]int, len(lines))
	for i, l := range lines {
		matches := regex.FindAllString(l, -1)
		row := make([]int, len(matches))

		for j, m := range matches {
			val, _ := strconv.Atoi(m)
			row[j] = val
		}

		result[i] = row
	}
	return result
}

func part1(sheet [][]int) int {
	sum := 0

	for _, row := range sheet {
		min, max := row[0], row[0]

		for _, i := range row {
			if i < min {
				min = i
			} else if i > max {
				max = i
			}
		}

		sum += max - min
	}

	return sum
}

func part2(sheet [][]int) int {
	sum := 0

	for _, row := range sheet {
		for i := 0; i < len(row); i++ {
			for j := i + 1; j < len(row); j++ {
				if row[i]%row[j] == 0 {
					sum += row[i] / row[j]
				} else if row[j]%row[i] == 0 {
					sum += row[j] / row[i]
				}
			}
		}
	}

	return sum
}
