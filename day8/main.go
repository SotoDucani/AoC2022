package main

import (
	"fmt"
	"strconv"
	"time"

	read "github.com/SotoDucani/AoC2022/internal/read"
)

func parseGrid() [][]string {
	array := read.ReadStrArrayByLine("./input.txt")

	var treeGrid [][]string

	for y := 0; y < len(array); y++ {
		chars := read.StrToCharArray(array[y])
		var treeRow []string
		for x := 0; x < len(chars); x++ {
			treeHeight := chars[x]
			treeRow = append(treeRow, treeHeight)
		}
		treeGrid = append(treeGrid, treeRow)
	}

	//for _, line := range treeGrid {
	//	fmt.Printf("%v\n", line)
	//}

	return treeGrid
}

func searchVisible(treeGrid [][]string, y int, x int) bool {
	isVisible := false
	orig, _ := strconv.Atoi(treeGrid[y][x])

	isVisible = searchToLeft(treeGrid, y, x, orig)
	if isVisible {
		return isVisible
	}
	isVisible = searchToRight(treeGrid, y, x, orig)
	if isVisible {
		return isVisible
	}
	isVisible = searchToUp(treeGrid, y, x, orig)
	if isVisible {
		return isVisible
	}
	isVisible = searchToDown(treeGrid, y, x, orig)
	if isVisible {
		return isVisible
	}

	return isVisible
}

func searchToLeft(treeGrid [][]string, y int, x int, orig int) bool {
	res := false
	if x == 0 {
		res = true
	} else {
		curLeft, _ := strconv.Atoi(treeGrid[y][x-1])
		if orig > curLeft {
			res = searchToLeft(treeGrid, y, x-1, orig)
		} else {
			res = false
		}
	}
	return res
}

func searchToRight(treeGrid [][]string, y int, x int, orig int) bool {
	res := false
	if x == len(treeGrid[0])-1 {
		res = true
	} else {
		curRight, _ := strconv.Atoi(treeGrid[y][x+1])
		if orig > curRight {
			res = searchToRight(treeGrid, y, x+1, orig)
		} else {
			res = false
		}
	}
	return res
}

func searchToUp(treeGrid [][]string, y int, x int, orig int) bool {
	res := false
	if y == 0 {
		res = true
	} else {
		curUp, _ := strconv.Atoi(treeGrid[y-1][x])
		if orig > curUp {
			res = searchToUp(treeGrid, y-1, x, orig)
		} else {
			res = false
		}
	}
	return res
}

func searchToDown(treeGrid [][]string, y int, x int, orig int) bool {
	res := false
	if y == len(treeGrid)-1 {
		res = true
	} else {
		curDown, _ := strconv.Atoi(treeGrid[y+1][x])
		if orig > curDown {
			res = searchToDown(treeGrid, y+1, x, orig)
		} else {
			res = false
		}
	}
	return res
}

func part1() {
	treeGrid := parseGrid()
	var visibleGrid [][]bool

	// Edges
	for y := 0; y < len(treeGrid); y++ {
		var visibleRow []bool
		for x := 0; x < len(treeGrid[y]); x++ {
			if (x == 0) || (y == 0) || (x == len(treeGrid[y])-1) || (y == len(treeGrid)-1) {
				visibleRow = append(visibleRow, true)
			} else {
				visibleRow = append(visibleRow, false)
			}
		}
		visibleGrid = append(visibleGrid, visibleRow)
	}

	for y := 0; y < len(treeGrid); y++ {
		for x := 0; x < len(treeGrid[y]); x++ {
			if visibleGrid[y][x] != true {
				visibleGrid[y][x] = searchVisible(treeGrid, y, x)
			}
		}
	}

	//for _, line := range visibleGrid {
	//	fmt.Printf("%v\n", line)
	//}

	count := 0
	for _, line := range visibleGrid {
		for _, res := range line {
			if res {
				count += 1
			}
		}
	}

	fmt.Printf("Part 1 - Visible Trees: %v\n", count)
}

func countScore(treeGrid [][]string, y int, x int) int {
	orig, _ := strconv.Atoi(treeGrid[y][x])

	leftScore := countLeft(treeGrid, y, x, orig, 0)
	rightScore := countRight(treeGrid, y, x, orig, 0)
	upScore := countUp(treeGrid, y, x, orig, 0)
	downScore := countDown(treeGrid, y, x, orig, 0)

	if leftScore == 0 {
		leftScore = 1
	}
	if rightScore == 0 {
		rightScore = 1
	}
	if upScore == 0 {
		upScore = 1
	}
	if downScore == 0 {
		downScore = 1
	}

	totalScore := rightScore * leftScore * upScore * downScore
	return totalScore
}

func countLeft(treeGrid [][]string, y int, x int, orig int, count int) int {
	var next int
	if x != 0 {
		next, _ = strconv.Atoi(treeGrid[y][x-1])
		if next >= orig {
			count += 1
		} else {
			count += 1
			count = countLeft(treeGrid, y, x-1, orig, count)
		}
	}
	return count
}

func countRight(treeGrid [][]string, y int, x int, orig int, count int) int {
	var next int
	if x != len(treeGrid[0])-1 {
		next, _ = strconv.Atoi(treeGrid[y][x+1])
		if next >= orig {
			count += 1
		} else {
			count += 1
			count = countRight(treeGrid, y, x+1, orig, count)
		}
	}
	return count
}

func countUp(treeGrid [][]string, y int, x int, orig int, count int) int {
	var next int
	if y != 0 {
		next, _ = strconv.Atoi(treeGrid[y-1][x])
		if next >= orig {
			count += 1
		} else {
			count += 1
			count = countUp(treeGrid, y-1, x, orig, count)
		}
	}
	return count
}

func countDown(treeGrid [][]string, y int, x int, orig int, count int) int {
	var next int
	if y != len(treeGrid)-1 {
		next, _ = strconv.Atoi(treeGrid[y+1][x])
		if next >= orig {
			count += 1
		} else {
			count += 1
			count = countDown(treeGrid, y+1, x, orig, count)
		}
	}
	return count
}

func part2() {
	treeGrid := parseGrid()
	var scoreGrid [][]int

	maxScore := 0
	for y := 0; y < len(treeGrid); y++ {
		var scoreRow []int
		for x := 0; x < len(treeGrid[y]); x++ {
			score := countScore(treeGrid, y, x)
			scoreRow = append(scoreRow, score)
			if score > maxScore {
				maxScore = score
			}
		}
		scoreGrid = append(scoreGrid, scoreRow)
	}

	//for _, line := range scoreGrid {
	//	fmt.Printf("%v\n", line)
	//}

	fmt.Printf("Part 2 - Best Scenic Score: %v\n", maxScore)
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
