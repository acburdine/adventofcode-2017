package main

import (
	"fmt"
	"math"
)

const (
	RIGHT = 0
	UP    = 1
	LEFT  = 2
	DOWN  = 3
	input = 325489
)

type xy struct {
	x int
	y int
}

type state struct {
	dir    int
	right  int
	top    int
	left   int
	bottom int
}

func main() {
	fmt.Printf("Day 3 Part 1: %d\n", part1())
	fmt.Printf("Day 3 Part 2: %d\n", part2())
}

func buildMem() []xy {
	mem := make([]xy, input)
	mem[0] = xy{0, 0}

	state := &state{}

	for i := 1; i < input; i++ {
		mem[i] = iterator(mem[i-1], state)
	}

	return mem
}

func iterator(last xy, st *state) xy {
	if st.dir == RIGHT {
		next := last.x + 1
		if next > st.right {
			st.right = next
			st.dir = UP
		}
		return xy{next, last.y}
	} else if st.dir == UP {
		next := last.y + 1
		if next > st.top {
			st.top = next
			st.dir = LEFT
		}
		return xy{last.x, next}
	} else if st.dir == LEFT {
		next := last.x - 1
		if next < st.left {
			st.left = next
			st.dir = DOWN
		}
		return xy{next, last.y}
	} else if st.dir == DOWN {
		next := last.y - 1
		if next < st.bottom {
			st.bottom = next
			st.dir = RIGHT
		}
		return xy{last.x, next}
	}

	return xy{}
}

func part1() int {
	mem := buildMem()
	c := mem[input-1]
	return int(math.Abs(float64(c.x)) + math.Abs(float64(c.y)))
}

func part2() int {
	mem := make(map[xy]int, input)
	last := xy{0, 0}
	mem[last] = 1

	state := &state{}
	check := func(val int, x int, y int) int {
		if v, ok := mem[xy{x, y}]; ok {
			return val + v
		}
		return val
	}

	for {
		next := iterator(last, state)
		val := 0

		val = check(val, next.x-1, next.y)
		val = check(val, next.x-1, next.y-1)
		val = check(val, next.x, next.y-1)
		val = check(val, next.x+1, next.y-1)
		val = check(val, next.x+1, next.y)
		val = check(val, next.x+1, next.y+1)
		val = check(val, next.x, next.y+1)
		val = check(val, next.x-1, next.y+1)

		if val > input {
			return val
		}

		mem[next] = val
		last = next
	}
}
