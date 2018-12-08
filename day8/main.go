package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/acburdine/adventofcode-2017"
)

const (
	inc = "inc"
	dec = "dec"
)

type instruction struct {
	reg     string
	act     string
	val     int
	condReg string
	condOpt string
	condVal int
}

func main() {
	lines := adventofcode.Input()
	insts := mapLines(lines)

	part1, part2 := run(insts)

	fmt.Printf("Day 8 Part 1: %d\n", part1)
	fmt.Printf("Day 8 Part 2: %d\n", part2)
}

func mapLines(lines []string) []instruction {
	inst := make([]instruction, len(lines))
	for i, l := range lines {
		sp := strings.Split(strings.TrimSpace(l), " ")
		ins := instruction{}
		ins.reg, ins.act, ins.condReg, ins.condOpt = sp[0], sp[1], sp[4], sp[5]
		val, _ := strconv.Atoi(sp[2])
		condVal, _ := strconv.Atoi(sp[6])
		ins.val, ins.condVal = val, condVal
		inst[i] = ins
	}
	return inst
}

func max(mp map[string]int) int {
	max := 0
	for _, v := range mp {
		if v > max {
			max = v
		}
	}
	return max
}

func eval(registers map[string]int, inst instruction) bool {
	switch inst.condOpt {
	case ">":
		return registers[inst.condReg] > inst.condVal
	case "<":
		return registers[inst.condReg] < inst.condVal
	case ">=":
		return registers[inst.condReg] >= inst.condVal
	case "<=":
		return registers[inst.condReg] <= inst.condVal
	case "!=":
		return registers[inst.condReg] != inst.condVal
	case "==":
		return registers[inst.condReg] == inst.condVal
	default:
		return false
	}
}

func run(insts []instruction) (int, int) {
	registers := make(map[string]int, len(insts))
	curMax := 0

	for _, inst := range insts {
		if eval(registers, inst) {
			if inst.act == inc {
				registers[inst.reg] = registers[inst.reg] + inst.val
			} else if inst.act == dec {
				registers[inst.reg] = registers[inst.reg] - inst.val
			}
		}

		tmp := max(registers)
		if tmp > curMax {
			curMax = tmp
		}
	}

	return max(registers), curMax
}
