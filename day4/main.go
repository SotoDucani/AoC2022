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
	containedSetsCount := 0

	for _, line := range array {
		isContained := false
		assignmentSplit := strings.Split(line, ",")
		partner1 := strings.Split(assignmentSplit[0], "-")
		partner2 := strings.Split(assignmentSplit[1], "-")
		partner1Low, _ := strconv.Atoi(partner1[0])
		partner1High, _ := strconv.Atoi(partner1[1])
		partner2Low, _ := strconv.Atoi(partner2[0])
		partner2High, _ := strconv.Atoi(partner2[1])

		// partner1 contained in partner2
		if partner1Low >= partner2Low && partner1High <= partner2High {
			//fmt.Printf("Partner 1;")
			isContained = true
		}
		// partner2 contained in partner1
		if partner2Low >= partner1Low && partner2High <= partner1High {
			//fmt.Printf("Partner 2;")
			isContained = true
		}

		// Add to count
		if isContained {
			//fmt.Printf(" contained; ")
			//fmt.Printf("Set: %v\n", assignmentSplit)
			containedSetsCount++
		}
	}
	fmt.Printf("Part 1 - Total contained sets: %v\n", containedSetsCount)
}

func part2() {
	array := read.ReadStrArrayByLine("./input.txt")
	overlappedSetsCount := 0

	for _, line := range array {
		isOverlapped := false
		assignmentSplit := strings.Split(line, ",")
		partner1 := strings.Split(assignmentSplit[0], "-")
		partner2 := strings.Split(assignmentSplit[1], "-")
		partner1Low, _ := strconv.Atoi(partner1[0])
		partner1High, _ := strconv.Atoi(partner1[1])
		partner2Low, _ := strconv.Atoi(partner2[0])
		partner2High, _ := strconv.Atoi(partner2[1])

		// partner1 overlapped in partner2
		if partner1Low >= partner2Low && partner1Low <= partner2High {
			//fmt.Printf("Partner 1 Low; ")
			isOverlapped = true
		}
		if partner1High >= partner2Low && partner1High <= partner2High {
			//fmt.Printf("Partner 1 High; ")
			isOverlapped = true
		}
		// partner2 contained in partner1
		if partner2Low >= partner1Low && partner2Low <= partner1High {
			//fmt.Printf("Partner 2 Low; ")
			isOverlapped = true
		}
		if partner2High >= partner1Low && partner2High <= partner1High {
			//fmt.Printf("Partner 2 High; ")
			isOverlapped = true
		}

		// Add to count
		if isOverlapped {
			//fmt.Printf("contained; ")
			//fmt.Printf("Set: %v\n", assignmentSplit)
			overlappedSetsCount++
		}
	}
	fmt.Printf("Part 2 - Total overlapped sets: %v\n", overlappedSetsCount)
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
