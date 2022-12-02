package main

import (
	"fmt"
	"time"

	read "github.com/SotoDucani/AoC2021/internal/read"
)

func part1() {
	array := read.ReadStrArrayByLine("./input.txt")

	winningState := []string{
		"A Y",
		"B Z",
		"C X",
	}
	tieState := []string{
		"A X",
		"B Y",
		"C Z",
	}

	totalScore := 0
	for _, line := range array {
		roundScore := 0
		if read.SliceContains(read.StrToCharArray(line), "X") {
			roundScore += 1
		} else if read.SliceContains(read.StrToCharArray(line), "Y") {
			roundScore += 2
		} else {
			roundScore += 3
		}

		if read.SliceContains(winningState, line) {
			roundScore += 6
		} else if read.SliceContains(tieState, line) {
			roundScore += 3
		}
		totalScore += roundScore
	}

	fmt.Printf("Part 1 - Total Score: %v\n", totalScore)
}

func part2() {
	array := read.ReadStrArrayByLine("./input.txt")

	totalScore := 0
	for _, line := range array {
		roundScore := 0

		roundSlice := read.StrToCharArray(line)

		if read.SliceContains(roundSlice, "Y") {
			roundScore += 3
			if read.SliceContains(roundSlice, "A") {
				roundScore += 1
			} else if read.SliceContains(roundSlice, "B") {
				roundScore += 2
			} else {
				roundScore += 3
			}
		} else if read.SliceContains(roundSlice, "Z") {
			roundScore += 6
			if read.SliceContains(roundSlice, "A") {
				roundScore += 2
			} else if read.SliceContains(roundSlice, "B") {
				roundScore += 3
			} else {
				roundScore += 1
			}
		} else {
			if read.SliceContains(roundSlice, "A") {
				roundScore += 3
			} else if read.SliceContains(roundSlice, "B") {
				roundScore += 1
			} else {
				roundScore += 2
			}
		}

		totalScore += roundScore
	}

	fmt.Printf("Part 2 - Total Score: %v\n", totalScore)
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
