package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	read "github.com/SotoDucani/AoC2022/internal/read"
)

func part1() {
	array := read.ReadStrArrayByLine("./exinput.txt")

	x := 1
	scanner := 0
	nextInst := 0
	cycle := 1
	sigStrengthSum := 0
	for line := 0; line < len(array); line++ {
		var curInst int
		curInst = nextInst
		// Instruction parse
		if array[line] != "noop" {
			incVal, _ := strconv.Atoi(strings.Split(array[line], " ")[1])
			nextInst = incVal
		} else {
			nextInst = 0
		}

		// Measure Current Value
		//fmt.Printf("Cycle: %v; X Val: %v\n", cycle, x)
		if cycle == 20 || cycle == (20+(40*scanner)) {
			scanner += 1
			fmt.Printf("Cycle: %v; X Val: %v\n", cycle, x)
			sigStrengthSum += x * (20 + (40 * scanner))
		}

		if array[line] != "noop" {
			cycle += 1
		}

		// Process instruction
		//fmt.Printf("Finished command: %v\n", curInst)
		x = x + curInst
		cycle += 1
	}

	fmt.Printf("Part 1 - Signal Strength Sum: %v\n", sigStrengthSum)
}

func part2() {
	//array := read.ReadStrArrayByLine("./input.txt")

	//fmt.Printf("Part 2 - String: %v\n", var)
}

func main() {
	p1b := time.Now()
	part1()
	mid := time.Now()
	part2()
	p2a := time.Now()
	part1Time := mid.Sub(p1b)
	part2Time := p2a.Sub(mid)
	fmt.Printf("Part 1 Time: %dμs\nPart 2 Time: %dμs\n", part1Time.Microseconds(), part2Time.Microseconds())
}
