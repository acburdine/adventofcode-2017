package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/acburdine/adventofcode-2017"
)

var regex = regexp.MustCompile("([a-z0-9]+)")

type program struct {
	name     string
	weight   int
	subs     []string
	parent   *program
	children []*program
}

func (p *program) weightsum() int {
	return p.weight + sum(p.children)
}

func main() {
	lines := adventofcode.Input()
	programs := mapPrograms(lines)

	root := root(programs)
	fmt.Printf("Day 7 Part 1: %s\n", root.name)
	fmt.Printf("Day 7 Part 2: %d\n", part2(programs, root))
}

func mapPrograms(lines []string) map[string]*program {
	result := make(map[string]*program, len(lines))
	for _, l := range lines {
		split := regex.FindAllString(l, -1)
		weight, _ := strconv.Atoi(split[1])
		result[split[0]] = &program{name: split[0], weight: weight, subs: split[2:]}
	}
	for _, prog := range result {
		prog.children = make([]*program, len(prog.subs))
		for i, sub := range prog.subs {
			result[sub].parent = prog
			prog.children[i] = result[sub]
		}
	}
	return result
}

func sum(arr []*program) int {
	sum := 0
	for _, i := range arr {
		sum += i.weightsum()
	}
	return sum
}

func groupBy(arr []int) map[int]int {
	result := make(map[int]int, len(arr))
	for _, i := range arr {
		result[i] = result[i] + 1
	}
	return result
}

func mapWeight(weights map[string]int, progs []*program) []int {
	result := make([]int, len(progs))
	for i, p := range progs {
		result[i] = weights[p.name]
	}
	return result
}

func root(programs map[string]*program) *program {
	for _, prog := range programs {
		if prog.parent == nil {
			return prog
		}
	}
	return nil
}

func recursor(prog program) int {
	if len(prog.children) < 2 {
		return prog.weightsum()
	}
}

func part2(programs map[string]*program, root program) int {
	weights := make(map[string]int, len(programs))

	for name, prog := range programs {
		weights[name] = prog.weight + sum(prog.subweights)
	}

	pt := root

	for {
		group := groupBy(mapWeight(weights, pt.children))
		if len(group) == 1 {

		}
	}

	for name, prog := range programs {
		if len(prog.subs) < 2 {
			continue
		}
		tmp := make([]int, len(prog.subweights))
		copy(tmp, prog.subweights)
		sort.Ints(tmp)
		if tmp[0] != tmp[1] {

		}
	}
}
