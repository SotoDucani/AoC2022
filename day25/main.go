package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	read "github.com/SotoDucani/AoC2022/internal/read"
)

func part1() {
	valueMap := map[string]int64{
		"2": 2,
		"1": 1,
		"0": 0,
		"-": -1,
		"=": -2,
	}
	revMap := map[string]string{
		"2":  "2",
		"1":  "1",
		"0":  "0",
		"-1": "-",
		"-2": "=",
	}

	array := read.ReadStrArrayByLine("./input.txt")

	var sum int64

	for _, line := range array {
		var decValue int64
		charArray := read.StrToCharArray(line)
		lineStack := read.Stack{}
		for _, character := range charArray {
			lineStack.Push(character)
		}

		totalLength := len(lineStack)
		for i := 0; i < totalLength; i++ {
			cur, notEmpty := lineStack.Pop()
			if !notEmpty {
				log.Panic("overran the stack")
			}
			var place int64
			if i == 0 {
				place = int64(1)
			} else {
				place = int64(math.Pow(5, float64(i)))
			}

			//fmt.Printf("equation: %v * %v\n", place, valueMap[cur])
			decValue = decValue + (place * valueMap[cur])
		}
		//fmt.Printf("DecValue now: %v\n", decValue)
		sum += decValue
	}

	//convert to base5
	base5 := strconv.FormatInt(sum, 5)
	b5Digits := read.StrToCharArray(base5)
	var FinalString read.Stack

	carry := 0
	for i := len(b5Digits) - 1; i >= 0; i-- {
		curDig, _ := strconv.Atoi(b5Digits[i])
		curDig = curDig + carry
		carry = 0
		if curDig > 2 {
			curDig = curDig - 5
			carry += 1
		}
		FinalString.Push(strconv.Itoa(curDig))
	}

	//Assemble final string
	var answer string
	stackLen := len(FinalString)
	for i := 0; i < stackLen; i++ {
		cur, _ := FinalString.Pop()
		answer = answer + revMap[cur]
	}
	fmt.Printf("Part 1 - Sum of fuel: %v\n", answer)
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
