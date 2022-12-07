package main

import (
	"fmt"
	"time"

	read "github.com/SotoDucani/AoC2022/internal/read"
)

func part1() {
	array := read.ReadStrArrayByLine("./input.txt")
	characters := read.StrToCharArray(array[0])

	reqCharCount := 0
	for i := 0; i < len(characters)-4; i++ {
		fmt.Printf("Current Seek: %s %s %s %s\n", characters[i], characters[i+1], characters[i+2], characters[i+3])
		if (characters[i] != characters[i+1]) && (characters[i] != characters[i+2]) && (characters[i] != characters[i+3]) && (characters[i+1] != characters[i+2]) && (characters[i+1] != characters[i+3]) && (characters[i+2] != characters[i+3]) {
			reqCharCount = i + 4
			break
		} else {
			fmt.Printf("Miss\n")
		}
	}
	fmt.Printf("Part 1 - SoP Marker Char Count: %v\n", reqCharCount)
}

func part2() {
	array := read.ReadStrArrayByLine("./input.txt")
	characters := read.StrToCharArray(array[0])

	reqCharCount := 0
	for i := 0; i < len(characters)-14; i++ {
		fmt.Printf("Current Seek: %v\n", characters[i:i+14])
		searchHash := make(map[string]int)
		var miss bool
		for _, letter := range characters[i : i+14] {
			searchHash[letter] += 1
			if searchHash[letter] > 1 {
				miss = true
				break
			}
		}
		if !miss {
			reqCharCount = i + 14
			break
		}
	}

	fmt.Printf("Part 2 - SoM Marker Char Count: %v\n", reqCharCount)
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
