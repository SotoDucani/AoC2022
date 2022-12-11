package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	read "github.com/SotoDucani/AoC2022/internal/read"
)

type monkey struct {
	items       read.Stack
	operation   string
	operationBy string
	testBy      int
	trueTarget  int
	falseTarget int
}

func processInput() map[int]monkey {
	array := read.ReadStrArrayByLine("./input.txt")

	monkeyMap := make(map[int]monkey)
	inputParseLine := 1
	var curMonkey monkey
	monkeyNum := 0
	for _, line := range array {
		//fmt.Printf("%s\n", line)
		switch inputParseLine {
		case 1:
			//fmt.Printf("Case 1\n")
		case 2:
			//fmt.Printf("Case 2\n")
			headSplit := strings.Split(line, ":")[1]
			items := strings.Split(headSplit, ",")
			for _, item := range items {
				curMonkey.items.Insert(strings.Trim(item, " "), 0)
			}
		case 3:
			//fmt.Printf("Case 3\n")
			headSplit := strings.Split(line, ":")
			operationSplit := strings.Split(headSplit[1], " ")
			curMonkey.operation = operationSplit[4]
			curMonkey.operationBy = operationSplit[5]
		case 4:
			//fmt.Printf("Case 4\n")
			headSplit := strings.Split(line, ":")
			curMonkey.testBy, _ = strconv.Atoi(strings.Split(headSplit[1], " ")[3])
		case 5:
			//fmt.Printf("Case 5\n")
			headSplit := strings.Split(line, ":")
			curMonkey.trueTarget, _ = strconv.Atoi(strings.Split(headSplit[1], " ")[4])
		case 6:
			//fmt.Printf("Case 6\n")
			headSplit := strings.Split(line, ":")
			curMonkey.falseTarget, _ = strconv.Atoi(strings.Split(headSplit[1], " ")[4])
		case 7:
			//fmt.Printf("Case 7\n")
		}

		// Append Monkey and reset
		if inputParseLine == 7 {
			monkeyMap[monkeyNum] = curMonkey
			inputParseLine = 1
			monkeyNum += 1
			curMonkey = monkey{}
		} else {
			inputParseLine += 1
		}
	}

	for _, m := range monkeyMap {
		fmt.Printf("%v\n", m)
	}

	return monkeyMap
}

func part1() {
	monkeys := processInput()

	monkeyActiveCount := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
	}

	for round := 0; round < 20; round++ {
		for monkeyNum := 0; monkeyNum < len(monkeys); monkeyNum++ {

		}
	}

	//fmt.Printf("Part 1 - Monkeys: %v\n", monkeys)
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
