package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	read "github.com/SotoDucani/AoC2022/internal/read"
)

var divisor int = 1

type monkey struct {
	items       []int
	operation   string
	operationBy string
	testBy      int
	trueTarget  int
	falseTarget int
}

func processInput() map[int]monkey {
	divisor = 1

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
				num, _ := strconv.Atoi(strings.Trim(item, " "))
				curMonkey.items = read.IntSliceInsert(curMonkey.items, num, 0)
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
			divisor *= curMonkey.testBy
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
	monkeyMap[monkeyNum] = curMonkey

	//for _, m := range monkeyMap {
	//	fmt.Printf("%v\n", m)
	//}

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

	// 20 Rounds
	for round := 0; round < 20; round++ {
		for monkeyNum := 0; monkeyNum < len(monkeys); monkeyNum++ {
			//fmt.Printf("Monkey: %v\n", monkeyNum)
			curMonkey := monkeys[monkeyNum]
			origItemCount := len(curMonkey.items)
			for itemNum := 0; itemNum < origItemCount; itemNum++ {
				var curItem int
				var hadValue bool
				curMonkey.items, curItem, hadValue = read.IntSlicePop(curMonkey.items)
				//fmt.Printf("CurItem: %v\n", curItem)
				if hadValue {
					monkeyActiveCount[monkeyNum] += 1
					// Inspect increase
					switch curMonkey.operation {
					case "*":
						switch curMonkey.operationBy {
						case "old":
							curItem = curItem * curItem
						default:
							opByInt, _ := strconv.Atoi(curMonkey.operationBy)
							curItem = curItem * opByInt
						}
					case "+":
						switch curMonkey.operationBy {
						case "old":
							curItem = curItem + curItem
						default:
							opByInt, _ := strconv.Atoi(curMonkey.operationBy)
							curItem = curItem + opByInt
						}
					}

					// Worry Drop
					curItem = curItem / 3

					// Pass To Other Monkey
					//fmt.Printf("Testing By: %v", curMonkey.testBy)
					modulo := curItem % curMonkey.testBy
					if modulo == 0 {
						targetMonkey := monkeys[curMonkey.trueTarget]
						targetMonkey.items = read.IntSliceInsert(targetMonkey.items, curItem, 0)
						monkeys[curMonkey.trueTarget] = targetMonkey
					} else {
						targetMonkey := monkeys[curMonkey.falseTarget]
						targetMonkey.items = read.IntSliceInsert(targetMonkey.items, curItem, 0)
						monkeys[curMonkey.falseTarget] = targetMonkey
					}

					//Update Monkey Map
					monkeys[monkeyNum] = curMonkey
				}
			}
			//fmt.Printf("====PostMonkeyState====\n")
			//for _, m := range monkeys {
			//	fmt.Printf("Monkey: %v\n", m)
			//}
		}
	}

	//for _, m := range monkeyActiveCount {
	//	fmt.Printf("Monkey: %v\n", m)
	//}
	//fmt.Printf("Part 1 - Monkeys: %v\n", monkeyActiveCount)
}

func part2() {
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
	fmt.Printf("Divisor: %v\n", divisor)
	// 10000 Rounds
	for round := 0; round < 10000; round++ {
		for monkeyNum := 0; monkeyNum < len(monkeys); monkeyNum++ {
			//fmt.Printf("Monkey: %v\n", monkeyNum)
			curMonkey := monkeys[monkeyNum]
			origItemCount := len(curMonkey.items)
			for itemNum := 0; itemNum < origItemCount; itemNum++ {
				var curItem int
				var hadValue bool
				curMonkey.items, curItem, hadValue = read.IntSlicePop(curMonkey.items)
				//fmt.Printf("CurItem: %v\n", curItem)
				if hadValue {
					monkeyActiveCount[monkeyNum] += 1
					// Inspect increase
					switch curMonkey.operation {
					case "*":
						switch curMonkey.operationBy {
						case "old":
							curItem = curItem * curItem
						default:
							opByInt, _ := strconv.Atoi(curMonkey.operationBy)
							curItem = curItem * opByInt
						}
					case "+":
						switch curMonkey.operationBy {
						case "old":
							curItem = curItem + curItem
						default:
							opByInt, _ := strconv.Atoi(curMonkey.operationBy)
							curItem = curItem + opByInt
						}
					}

					// Worry Drop
					curItem = curItem % divisor

					// Pass To Other Monkey
					//fmt.Printf("Testing By: %v", curMonkey.testBy)
					modulo := curItem % curMonkey.testBy
					if modulo == 0 {
						targetMonkey := monkeys[curMonkey.trueTarget]
						targetMonkey.items = read.IntSliceInsert(targetMonkey.items, curItem, 0)
						monkeys[curMonkey.trueTarget] = targetMonkey
					} else {
						targetMonkey := monkeys[curMonkey.falseTarget]
						targetMonkey.items = read.IntSliceInsert(targetMonkey.items, curItem, 0)
						monkeys[curMonkey.falseTarget] = targetMonkey
					}

					//Update Monkey Map
					monkeys[monkeyNum] = curMonkey
				}
			}
			//fmt.Printf("====PostMonkeyState====\n")
			//for _, m := range monkeys {
			//	fmt.Printf("Monkey: %v\n", m)
			//}
		}
	}

	var values []int
	for _, m := range monkeyActiveCount {
		values = append(values, m)
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})

	final := values[0] * values[1]
	fmt.Printf("Part 2 - Monkeys: %v\n", final)
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
