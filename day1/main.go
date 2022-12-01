package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	read "github.com/SotoDucani/AoC2021/internal/read"
)

func part1() {
	array := read.ReadStrArrayByLine("./input.txt")

	var calCounts []int
	maxCal := 0
	sum := 0
	for _, line := range array {
		if line == "" {
			calCounts = append(calCounts, sum)
			if sum > maxCal {
				maxCal = sum
			}
			//fmt.Printf("Appending: %v; MaxCal: %v\n", sum, maxCal)
			sum = 0
		} else {
			num, _ := strconv.Atoi(line)
			sum = sum + num
			//fmt.Printf("SameElf\n")
		}
		//fmt.Printf("Sum: %v\n", sum)
	}

	fmt.Printf("Part 1 - Max calorie count: %v\n", maxCal)
}

func part2() {
	array := read.ReadStrArrayByLine("./input.txt")

	var calCounts []int
	maxCal := 0
	sum := 0
	for _, line := range array {
		if line == "" {
			calCounts = append(calCounts, sum)
			if sum > maxCal {
				maxCal = sum
			}
			//fmt.Printf("Appending: %v; MaxCal: %v\n", sum, maxCal)
			sum = 0
		} else {
			num, _ := strconv.Atoi(line)
			sum = sum + num
			//fmt.Printf("SameElf\n")
		}
		//fmt.Printf("Sum: %v\n", sum)
	}
	sort.Ints(calCounts)
	len := len(calCounts)
	totalSum := calCounts[len-1] + calCounts[len-2] + calCounts[len-3]
	fmt.Printf("Part 2 - Sum of top 3 calorie count: %v\n", totalSum)
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
