package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	read "github.com/SotoDucani/AoC2022/internal/read"
)

func Insert(stk []string, str string, pos int) []string {
	stk = append(stk, "")
	copy(stk[pos+1:], stk[pos:])
	stk[pos] = str
	return stk
}

func catchHead(head map[string]int, tail map[string]int) map[string]int {
	//fmt.Printf("Catching Head: %v\n", head)
	// xDiff positive - Tail is left of Head
	// xDiff negative - Tail is right of Head
	xDiff := head["x"] - tail["x"]
	// yDiff positive - Tail is below Head
	// yDiff negative - Tail is above Head
	yDiff := head["y"] - tail["y"]

	if (xDiff == 2 && (yDiff == 1 || yDiff == 2)) || (xDiff == 1 && yDiff == 2) {
		// Head is up and to the right
		tail["x"] += 1
		tail["y"] += 1
	} else if (xDiff == 2 && (yDiff == -2 || yDiff == -1)) || (xDiff == 1 && yDiff == -2) {
		// Head is below and to the right
		tail["x"] += 1
		tail["y"] -= 1
	} else if (xDiff == -2 && (yDiff == -2 || yDiff == -1)) || (xDiff == -1 && yDiff == -2) {
		// Head is below and to the left
		tail["x"] -= 1
		tail["y"] -= 1
	} else if (xDiff == -2 && (yDiff == 2 || yDiff == 1)) || (xDiff == -1 && yDiff == 2) {
		// Head is up and to the left
		tail["x"] -= 1
		tail["y"] += 1
	} else if xDiff == 2 {
		// Head is to the right
		tail["x"] += 1
	} else if xDiff == -2 {
		// Head is to the left
		tail["x"] -= 1
	} else if yDiff == 2 {
		// Head is up
		tail["y"] += 1
	} else if yDiff == -2 {
		// Head is down
		tail["y"] -= 1
	}

	return tail
}

func draw(tailVisitedLocations map[string]int) {
	var drawnMap [][]string
	for y := -20; y < 21; y++ {
		var curLine []string
		for x := -20; x < 21; x++ {
			curLine = append(curLine, ".")
		}
		drawnMap = append(drawnMap, curLine)
	}

	//for _, line := range drawnMap {
	//	curLine := ""
	//	for _, char := range line {
	//		curLine = curLine + char
	//	}
	//	fmt.Printf("%s\n", curLine)
	//}

	for pair, _ := range tailVisitedLocations {
		split := strings.Split(pair, ",")
		y, _ := strconv.Atoi(split[1])
		x, _ := strconv.Atoi(split[0])
		drawnMap[y+20][x+20] = "X"
	}

	drawnMap[0+20][0+20] = "S"

	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	var printStack []string
	for _, line := range drawnMap {
		curLine := ""
		for _, char := range line {
			curLine = curLine + char
		}
		printStack = Insert(printStack, curLine, 0)
	}

	for _, line := range printStack {
		fmt.Printf("%s\n", line)
	}
}

func part1() {
	array := read.ReadStrArrayByLine("./input.txt")
	tailVisitedLocations := make(map[string]int)

	// (x,y)
	head := make(map[string]int)
	head["x"] = 0
	head["y"] = 0
	tail := make(map[string]int)
	tail["x"] = 0
	tail["y"] = 0

	for _, line := range array {
		instructionSplit := strings.Split(line, " ")
		steps, _ := strconv.Atoi(instructionSplit[1])
		for step := 0; step < steps; step++ {
			switch instructionSplit[0] {
			case "U":
				head["y"] += 1
				tail = catchHead(head, tail)
				tailVisitedLocations[fmt.Sprintf("%v,%v", tail["x"], tail["y"])] += 1
			case "D":
				head["y"] -= 1
				tail = catchHead(head, tail)
				tailVisitedLocations[fmt.Sprintf("%v,%v", tail["x"], tail["y"])] += 1
			case "L":
				head["x"] -= 1
				tail = catchHead(head, tail)
				tailVisitedLocations[fmt.Sprintf("%v,%v", tail["x"], tail["y"])] += 1
			case "R":
				head["x"] += 1
				tail = catchHead(head, tail)
				tailVisitedLocations[fmt.Sprintf("%v,%v", tail["x"], tail["y"])] += 1
			}
			//fmt.Printf("Head: %v\n", head)
			//fmt.Printf("Tail: %v\n", tail)
			//fmt.Printf("finishedMove\n")
		}
	}

	uniqueLocations := len(tailVisitedLocations)
	//for key, _ := range tailVisitedLocations {
	//fmt.Printf("Location: %v\n", key)
	//}
	fmt.Printf("Part 1 - Unique Locations: %v\n", uniqueLocations)
}

func part2() {
	array := read.ReadStrArrayByLine("./input.txt")
	tailVisitedLocations := make(map[string]int)

	// (x,y)
	var knots []map[string]int

	for ok := true; ok; ok = (len(knots) < 10) {
		cur := make(map[string]int)
		cur["x"] = 0
		cur["y"] = 0
		knots = append(knots, cur)
	}

	for _, line := range array {
		instructionSplit := strings.Split(line, " ")
		steps, _ := strconv.Atoi(instructionSplit[1])
		for step := 0; step < steps; step++ {
			//fmt.Printf("Ins:%v Step: %v\n", instructionSplit[0], step)
			switch instructionSplit[0] {
			case "U":
				// get copy of our head knot
				prevKnot := knots[0]
				// move our head knot
				prevKnot["y"] += 1
				// reassign our head knot value back into chain
				knots[0] = prevKnot

				// for the rest of the knots
				for i := 1; i < len(knots); i++ {
					// extract copy of the current knot
					curKnot := knots[i]
					// catch it up to the previously moved knot
					curKnot = catchHead(prevKnot, curKnot)
					// reassign the new values back into the chain
					knots[i] = curKnot
					// also assign it to prevKnot for the next loop
					prevKnot = curKnot
					// if last knot, track it
					if i == len(knots)-1 {
						tailVisitedLocations[fmt.Sprintf("%v,%v", curKnot["x"], curKnot["y"])] += 1
					}
				}
			case "D":
				// get copy of our head knot
				prevKnot := knots[0]
				// move our head knot
				prevKnot["y"] -= 1
				// reassign our head knot value back into chain
				knots[0] = prevKnot

				// for the rest of the knots
				for i := 1; i < len(knots); i++ {
					// extract copy of the current knot
					curKnot := knots[i]
					// catch it up to the previously moved knot
					curKnot = catchHead(prevKnot, curKnot)
					// reassign the new values back into the chain
					knots[i] = curKnot
					// also assign it to prevKnot for the next loop
					prevKnot = curKnot
					// if last knot, track it
					if i == len(knots)-1 {
						tailVisitedLocations[fmt.Sprintf("%v,%v", curKnot["x"], curKnot["y"])] += 1
					}
				}
			case "L":
				// get copy of our head knot
				prevKnot := knots[0]
				// move our head knot
				prevKnot["x"] -= 1
				// reassign our head knot value back into chain
				knots[0] = prevKnot

				// for the rest of the knots
				for i := 1; i < len(knots); i++ {
					// extract copy of the current knot
					curKnot := knots[i]
					// catch it up to the previously moved knot
					curKnot = catchHead(prevKnot, curKnot)
					// reassign the new values back into the chain
					knots[i] = curKnot
					// also assign it to prevKnot for the next loop
					prevKnot = curKnot
					// if last knot, track it
					if i == len(knots)-1 {
						tailVisitedLocations[fmt.Sprintf("%v,%v", curKnot["x"], curKnot["y"])] += 1
					}
				}
			case "R":
				// get copy of our head knot
				prevKnot := knots[0]
				// move our head knot
				prevKnot["x"] += 1
				// reassign our head knot value back into chain
				knots[0] = prevKnot

				// for the rest of the knots
				for i := 1; i < len(knots); i++ {
					// extract copy of the current knot
					curKnot := knots[i]
					// catch it up to the previously moved knot
					curKnot = catchHead(prevKnot, curKnot)
					// reassign the new values back into the chain
					knots[i] = curKnot
					// also assign it to prevKnot for the next loop
					prevKnot = curKnot
					// if last knot, track it
					if i == len(knots)-1 {
						tailVisitedLocations[fmt.Sprintf("%v,%v", curKnot["x"], curKnot["y"])] += 1
					}
				}
			}
		}
		// Draw current chain
		/*
			currentLocations := make(map[string]int)
			for _, knot := range knots {
				currentLocations[fmt.Sprintf("%v,%v", knot["x"], knot["y"])] += 1
			}
			draw(currentLocations)
			fmt.Printf("====\n")
		*/
	}

	uniqueLocations := len(tailVisitedLocations)
	//draw(tailVisitedLocations)
	//for key, _ := range tailVisitedLocations {
	//	fmt.Printf("Location: %v\n", key)
	//}
	fmt.Printf("Part 2 - Unique Locations: %v\n", uniqueLocations)
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
