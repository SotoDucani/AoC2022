package main

import (
	"fmt"
	"strconv"
	"time"

	read "github.com/SotoDucani/AoC2022/internal/read"
)

func part1() {
	array := read.ReadStrArrayByLine("./input.txt")

	stacks := make(map[int]read.Stack)
	stack1 := read.Stack{}
	stack2 := read.Stack{}
	stack3 := read.Stack{}
	stack4 := read.Stack{}
	stack5 := read.Stack{}
	stack6 := read.Stack{}
	stack7 := read.Stack{}
	stack8 := read.Stack{}
	stack9 := read.Stack{}

	stack1.Push("W")
	stack1.Push("B")
	stack1.Push("D")
	stack1.Push("N")
	stack1.Push("C")
	stack1.Push("F")
	stack1.Push("J")

	stack2.Push("P")
	stack2.Push("Z")
	stack2.Push("V")
	stack2.Push("Q")
	stack2.Push("L")
	stack2.Push("S")
	stack2.Push("T")

	stack3.Push("P")
	stack3.Push("Z")
	stack3.Push("B")
	stack3.Push("G")
	stack3.Push("J")
	stack3.Push("T")

	stack4.Push("D")
	stack4.Push("T")
	stack4.Push("L")
	stack4.Push("J")
	stack4.Push("Z")
	stack4.Push("B")
	stack4.Push("H")
	stack4.Push("C")

	stack5.Push("G")
	stack5.Push("V")
	stack5.Push("B")
	stack5.Push("J")
	stack5.Push("S")

	stack6.Push("P")
	stack6.Push("S")
	stack6.Push("Q")

	stack7.Push("B")
	stack7.Push("V")
	stack7.Push("D")
	stack7.Push("F")
	stack7.Push("L")
	stack7.Push("M")
	stack7.Push("P")
	stack7.Push("N")

	stack8.Push("P")
	stack8.Push("S")
	stack8.Push("M")
	stack8.Push("F")
	stack8.Push("B")
	stack8.Push("D")
	stack8.Push("L")
	stack8.Push("R")

	stack9.Push("V")
	stack9.Push("D")
	stack9.Push("T")
	stack9.Push("R")

	stacks[1] = stack1
	stacks[2] = stack2
	stacks[3] = stack3
	stacks[4] = stack4
	stacks[5] = stack5
	stacks[6] = stack6
	stacks[7] = stack7
	stacks[8] = stack8
	stacks[9] = stack9

	for _, line := range array {
		words := read.StrToWordArray(line)
		moveCount, _ := strconv.Atoi(words[1])
		source, _ := strconv.Atoi(words[3])
		dest, _ := strconv.Atoi(words[5])
		srcStack := stacks[source]
		destStack := stacks[dest]

		for i := 0; i < moveCount; i++ {
			cur, _ := srcStack.Pop()
			destStack.Push(cur)
		}
		stacks[source] = srcStack
		stacks[dest] = destStack
		//fmt.Printf("CurStacks: %v\n", stacks)
	}

	finalString := ""
	for i := 1; i < 10; i++ {
		curStack := stacks[i]
		curChar, _ := curStack.Pop()
		finalString = finalString + curChar
	}

	//Doesn't work, wtf
	//for _, stack := range stacks {
	//	curChar, _ := stack.Pop()
	//	finalString = finalString + curChar
	//}

	fmt.Printf("Part 1 - Top Crates: %v\n", finalString)
}

func part2() {
	array := read.ReadStrArrayByLine("./input.txt")

	stacks := make(map[int]read.Stack)
	stack1 := read.Stack{}
	stack2 := read.Stack{}
	stack3 := read.Stack{}
	stack4 := read.Stack{}
	stack5 := read.Stack{}
	stack6 := read.Stack{}
	stack7 := read.Stack{}
	stack8 := read.Stack{}
	stack9 := read.Stack{}

	stack1.Push("W")
	stack1.Push("B")
	stack1.Push("D")
	stack1.Push("N")
	stack1.Push("C")
	stack1.Push("F")
	stack1.Push("J")

	stack2.Push("P")
	stack2.Push("Z")
	stack2.Push("V")
	stack2.Push("Q")
	stack2.Push("L")
	stack2.Push("S")
	stack2.Push("T")

	stack3.Push("P")
	stack3.Push("Z")
	stack3.Push("B")
	stack3.Push("G")
	stack3.Push("J")
	stack3.Push("T")

	stack4.Push("D")
	stack4.Push("T")
	stack4.Push("L")
	stack4.Push("J")
	stack4.Push("Z")
	stack4.Push("B")
	stack4.Push("H")
	stack4.Push("C")

	stack5.Push("G")
	stack5.Push("V")
	stack5.Push("B")
	stack5.Push("J")
	stack5.Push("S")

	stack6.Push("P")
	stack6.Push("S")
	stack6.Push("Q")

	stack7.Push("B")
	stack7.Push("V")
	stack7.Push("D")
	stack7.Push("F")
	stack7.Push("L")
	stack7.Push("M")
	stack7.Push("P")
	stack7.Push("N")

	stack8.Push("P")
	stack8.Push("S")
	stack8.Push("M")
	stack8.Push("F")
	stack8.Push("B")
	stack8.Push("D")
	stack8.Push("L")
	stack8.Push("R")

	stack9.Push("V")
	stack9.Push("D")
	stack9.Push("T")
	stack9.Push("R")

	stacks[1] = stack1
	stacks[2] = stack2
	stacks[3] = stack3
	stacks[4] = stack4
	stacks[5] = stack5
	stacks[6] = stack6
	stacks[7] = stack7
	stacks[8] = stack8
	stacks[9] = stack9

	for _, line := range array {
		words := read.StrToWordArray(line)
		moveCount, _ := strconv.Atoi(words[1])
		source, _ := strconv.Atoi(words[3])
		dest, _ := strconv.Atoi(words[5])
		srcStack := stacks[source]
		destStack := stacks[dest]

		middlingStack := read.Stack{}

		for i := 0; i < moveCount; i++ {
			cur, _ := srcStack.Pop()
			middlingStack.Push(cur)
		}
		for i := 0; i < moveCount; i++ {
			cur, _ := middlingStack.Pop()
			destStack.Push(cur)
		}

		stacks[source] = srcStack
		stacks[dest] = destStack
		//fmt.Printf("CurStacks: %v\n", stacks)
	}

	finalString := ""
	for i := 1; i < 10; i++ {
		curStack := stacks[i]
		curChar, _ := curStack.Pop()
		finalString = finalString + curChar
	}

	//Doesn't work, wtf
	//for _, stack := range stacks {
	//	curChar, _ := stack.Pop()
	//	finalString = finalString + curChar
	//}

	fmt.Printf("Part 2 - Top Crates: %v\n", finalString)
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
