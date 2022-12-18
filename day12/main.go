package main

import (
	"fmt"
	"time"

	read "github.com/SotoDucani/AoC2022/internal/read"
)

func part1() {
	array := read.ReadStrArrayByLine("./input.txt")

	fmt.Printf("Part 1 - String: %v\n", var)
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
