package main

import (
	"fmt"
	"time"

	read "github.com/SotoDucani/AoC2022/internal/read"
)

func part1() {
	points := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
		"A": 27,
		"B": 28,
		"C": 29,
		"D": 30,
		"E": 31,
		"F": 32,
		"G": 33,
		"H": 34,
		"I": 35,
		"J": 36,
		"K": 37,
		"L": 38,
		"M": 39,
		"N": 40,
		"O": 41,
		"P": 42,
		"Q": 43,
		"R": 44,
		"S": 45,
		"T": 46,
		"U": 47,
		"V": 48,
		"W": 49,
		"X": 50,
		"Y": 51,
		"Z": 52,
	}

	array := read.ReadStrArrayByLine("./input.txt")
	totalSum := 0

	for _, line := range array {
		charArray := read.StrToCharArray(line)
		left := charArray[0:(len(charArray) / 2)]
		right := charArray[(len(charArray) / 2):]
		fmt.Printf("Left: %v; Right: %v\n", left, right)
		var dupedChar string
		for _, char := range left {
			if read.SliceContains(right, char) {
				dupedChar = char
				fmt.Printf("DupChar: %v\n", char)
			}
		}
		totalSum += points[dupedChar]
	}

	fmt.Printf("Part 1 - Total priority sum: %v\n", totalSum)
}

func part2() {
	points := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
		"A": 27,
		"B": 28,
		"C": 29,
		"D": 30,
		"E": 31,
		"F": 32,
		"G": 33,
		"H": 34,
		"I": 35,
		"J": 36,
		"K": 37,
		"L": 38,
		"M": 39,
		"N": 40,
		"O": 41,
		"P": 42,
		"Q": 43,
		"R": 44,
		"S": 45,
		"T": 46,
		"U": 47,
		"V": 48,
		"W": 49,
		"X": 50,
		"Y": 51,
		"Z": 52,
	}

	array := read.ReadStrArrayByLine("./input.txt")
	totalSum := 0

	for i := 0; i < len(array); i = i + 3 {
		ruck1 := read.StrToCharArray(array[i])
		ruck2 := read.StrToCharArray(array[i+1])
		ruck3 := read.StrToCharArray(array[i+2])
		var all []string
		all = append(all, ruck1...)
		all = append(all, ruck2...)
		all = append(all, ruck3...)
		var badgeChar string
		for _, char := range all {
			if read.SliceContains(ruck1, char) && read.SliceContains(ruck2, char) && read.SliceContains(ruck3, char) {
				badgeChar = char
			}
		}
		totalSum += points[badgeChar]
	}

	fmt.Printf("Part 2 - Total priority sum: %v\n", totalSum)
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
