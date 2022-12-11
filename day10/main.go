package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	read "github.com/SotoDucani/AoC2022/internal/read"
)

func part1() {
	array := read.ReadStrArrayByLine("./input.txt")

	x := 1
	scanner := 0
	instructionStack := read.Stack{}
	cycle := 1
	sigStrengthSum := 0

	for _, line := range array {
		//Read
		lineSplit := strings.Split(line, " ")
		if lineSplit[0] == "noop" {
			instructionStack.Insert("0", 0)
		} else {
			instructionStack.Insert("0", 0)
			instructionStack.Insert(lineSplit[1], 0)
		}
	}

	for ok := true; ok; ok = (len(instructionStack) > 0) {
		// Measure value for cycle
		//fmt.Printf("Cycle: %v; X Val: %v\n", cycle, x)
		if cycle == 20 || cycle == (20+(40*scanner)) {
			scanner += 1
			//fmt.Printf("Cycle: %v; X Val: %v\n", cycle, x)
			sigStrengthSum += x * cycle
		}

		// Process Instruction
		cur, hadValue := instructionStack.Pop()
		if hadValue {
			curInt, _ := strconv.Atoi(cur)
			x += curInt
		}
		cycle += 1
	}

	fmt.Printf("Part 1 - Signal Strength Sum: %v\n", sigStrengthSum)
}

func part2() {
	array := read.ReadStrArrayByLine("./input.txt")

	var resultArray []string

	x := 1
	scanner := 1
	instructionStack := read.Stack{}
	cycle := 1
	crtPos := 0

	for _, line := range array {
		//Read
		lineSplit := strings.Split(line, " ")
		if lineSplit[0] == "noop" {
			instructionStack.Insert("0", 0)
		} else {
			instructionStack.Insert("0", 0)
			instructionStack.Insert(lineSplit[1], 0)
		}
	}

	curCRTLine := ""
	for ok := true; ok; ok = (len(instructionStack) > 0) {
		if crtPos == x-1 || crtPos == x || crtPos == x+1 {
			curCRTLine = curCRTLine + "X"
		} else {
			curCRTLine = curCRTLine + "."
		}

		// Process Instruction
		cur, hadValue := instructionStack.Pop()
		if hadValue {
			curInt, _ := strconv.Atoi(cur)
			x += curInt
		}

		crtPos += 1
		// Inc CRT Line
		if cycle == (40 * scanner) {
			scanner += 1
			crtPos = 0
			resultArray = append(resultArray, curCRTLine)
			curCRTLine = ""
		}
		cycle += 1

	}

	for _, line := range resultArray {
		fmt.Printf("%v\n", line)
	}
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
