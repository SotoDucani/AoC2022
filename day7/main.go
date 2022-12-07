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
	curDirStack := read.Stack{}
	sizeMap := make(map[string]int)

	for _, line := range array {
		spaceSplit := strings.Split(line, " ")

		if (spaceSplit[0] == "$") && (spaceSplit[1] == "cd") && (spaceSplit[2] == "..") {
			//Leave dir
			curDirStack.Pop()
			//fmt.Printf("Current Dir: %v\n", curDirStack)
		} else if (spaceSplit[0] == "$") && (spaceSplit[1] == "cd") && (spaceSplit[2] != "..") {
			//Enter dir
			curDirStack.Push(spaceSplit[2])
			//fmt.Printf("Current Dir: %v\n", curDirStack)
		} else if (spaceSplit[0] == "$") && (spaceSplit[1] == "ls") {
			//fmt.Printf("Listing Dir\n")
		} else if spaceSplit[0] == "dir" {
			//fmt.PrintF("Found Directory")
		} else {
			copyDirStack := read.Stack{}
			for _, dir := range curDirStack {
				copyDirStack.Push(dir)
			}
			for i := 0; i < len(curDirStack); i++ {
				//fmt.Printf("Adding file: %v; Size: %v; Dir: %v\n", spaceSplit[1], spaceSplit[0], copyDirStack)
				fileSize, _ := strconv.Atoi(spaceSplit[0])
				sizeMap[fmt.Sprintf("%v", copyDirStack)] += fileSize
				copyDirStack.Pop()
			}
		}
	}

	sizeSum := 0
	for _, dirSize := range sizeMap {
		if dirSize <= 100000 {
			//fmt.Printf("YES - Dir: %v; Size: %v\n", dirName, dirSize)
			sizeSum += dirSize
		} else {
			//fmt.Printf("no - Dir: %v; Size: %v\n", dirName, dirSize)
		}
	}

	fmt.Printf("Part 1 - Sum of sizes under 100000: %v\n", sizeSum)
}

func part2() {
	array := read.ReadStrArrayByLine("./input.txt")
	curDirStack := read.Stack{}
	sizeMap := make(map[string]int)

	for _, line := range array {
		spaceSplit := strings.Split(line, " ")

		if (spaceSplit[0] == "$") && (spaceSplit[1] == "cd") && (spaceSplit[2] == "..") {
			//Leave dir
			curDirStack.Pop()
			//fmt.Printf("Current Dir: %v\n", curDirStack)
		} else if (spaceSplit[0] == "$") && (spaceSplit[1] == "cd") && (spaceSplit[2] != "..") {
			//Enter dir
			curDirStack.Push(spaceSplit[2])
			//fmt.Printf("Current Dir: %v\n", curDirStack)
		} else if (spaceSplit[0] == "$") && (spaceSplit[1] == "ls") {
			//fmt.Printf("Listing Dir\n")
		} else if spaceSplit[0] == "dir" {
			//fmt.PrintF("Found Directory")
		} else {
			copyDirStack := read.Stack{}
			for _, dir := range curDirStack {
				copyDirStack.Push(dir)
			}
			for i := 0; i < len(curDirStack); i++ {
				//fmt.Printf("Adding file: %v; Size: %v; Dir: %v\n", spaceSplit[1], spaceSplit[0], copyDirStack)
				fileSize, _ := strconv.Atoi(spaceSplit[0])
				sizeMap[fmt.Sprintf("%v", copyDirStack)] += fileSize
				copyDirStack.Pop()
			}
		}
	}

	root := read.Stack{}
	root.Push("/")
	totalSizeUsed := sizeMap[fmt.Sprintf("%v", root)]

	minimumSizeNeeded := totalSizeUsed - (70000000 - 30000000)

	fmt.Printf("Used Size: %v; Minimum Size Needed: %v\n", totalSizeUsed, minimumSizeNeeded)

	sizeToDelete := 90000000 // just an abnormally large value to start
	for _, size := range sizeMap {
		if (size >= minimumSizeNeeded) && (size < sizeToDelete) {
			sizeToDelete = size
		}
	}

	fmt.Printf("Part 2 - Size To Delete: %v\n", sizeToDelete)
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
