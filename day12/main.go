package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	read "github.com/SotoDucani/AoC2022/internal/read"
	"github.com/fatih/color"
)

func Insert(stk []string, str string, pos int) []string {
	stk = append(stk, "")
	copy(stk[pos+1:], stk[pos:])
	stk[pos] = str
	return stk
}

// Ref: https://go-recipes.dev/dijkstras-algorithm-in-go-e1129b2f5c9e
// Literally have never written this from scratch before

type Node struct {
	name     string
	distance int
	through  *Node
}

type Edge struct {
	node   *Node
	weight int
}

type WeightedGraph struct {
	Nodes []*Node
	Edges map[string][]*Edge
	mutex sync.RWMutex
}

type Heap struct {
	elements []*Node
	mutex    sync.RWMutex
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

func (h *Heap) swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func (h *Heap) Size() int {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	return len(h.elements)
}

func (h *Heap) Push(element *Node) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.elements = append(h.elements, element)
	i := len(h.elements) - 1
	for ; h.elements[i].distance < h.elements[parent(i)].distance; i = parent(i) {
		h.swap(i, parent(i))
	}
}

func (h *Heap) rearrange(i int) {
	smallest := i
	left, right, size := leftChild(i), rightChild(i), len(h.elements)
	if left < size && h.elements[left].distance < h.elements[smallest].distance {
		smallest = left
	}
	if right < size && h.elements[right].distance < h.elements[smallest].distance {
		smallest = right
	}
	if smallest != i {
		h.swap(i, smallest)
		h.rearrange(smallest)
	}
}

func (h *Heap) Pop() (i *Node) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	i = h.elements[0]
	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]
	h.rearrange(0)
	return
}

func NewGraph() *WeightedGraph {
	return &WeightedGraph{
		Edges: make(map[string][]*Edge),
	}
}

func (g *WeightedGraph) GetNode(name string) (node *Node) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	for _, n := range g.Nodes {
		if n.name == name {
			node = n
		}
	}
	return
}

func (g *WeightedGraph) AddNode(n *Node) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.Nodes = append(g.Nodes, n)
}

func AddNodes(g *WeightedGraph, names ...string) (nodes map[string]*Node) {
	nodes = make(map[string]*Node)
	for _, name := range names {
		n := &Node{name, math.MaxInt, nil}
		g.AddNode(n)
		nodes[name] = n
	}
	return
}

func (g *WeightedGraph) AddEdge(n1, n2 *Node, weight int) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.Edges[n1.name] = append(g.Edges[n1.name], &Edge{n2, weight})
	g.Edges[n2.name] = append(g.Edges[n2.name], &Edge{n1, weight})
}

func processInput(filename string) [][]int {
	array := read.ReadStrArrayByLine(filename)

	var intMap [][]int
	for _, line := range array {
		charArray := read.StrToCharArray(line)
		var curIntLine []int
		for _, char := range charArray {
			if char == "S" {
				runeChar := []rune("a")
				curIntLine = append(curIntLine, int(runeChar[0])-96)
			} else if char == "E" {
				runeChar := []rune("z")
				curIntLine = append(curIntLine, int(runeChar[0])-96)
			} else {
				runeChar := []rune(char)
				curIntLine = append(curIntLine, int(runeChar[0])-96)
			}
		}
		intMap = append(intMap, curIntLine)
	}

	//for _, line := range intMap {
	//	fmt.Printf("%v\n", line)
	//}

	return intMap
}

func setupGraph(intMap [][]int, startNodeName string, destNodeName string) *WeightedGraph {
	graph := NewGraph()

	var nodeNames []string
	// SetupNodes
	// TopLeft origin (0,0) point
	for y := 0; y < len(intMap); y++ {
		for x := 0; x < len(intMap[y]); x++ {
			nodeNames = append(nodeNames, fmt.Sprintf("%d,%d", x, y))
		}
	}
	nodes := AddNodes(graph, nodeNames...)

	// Add Edges
	// TopLeft origin (0,0) point
	for y := 0; y < len(intMap); y++ {
		for x := 0; x < len(intMap[y]); x++ {
			if y-1 >= 0 {
				curNodeName := fmt.Sprintf("%d,%d", x, y)
				nextNodeName := fmt.Sprintf("%d,%d", x, y-1)
				//fmt.Printf("Weight from %s to %s: %d\n", curNodeName, nextNodeName, weight)
				graph.AddEdge(nodes[curNodeName], nodes[nextNodeName], 1)
			}
			if y+1 < len(intMap) {
				curNodeName := fmt.Sprintf("%d,%d", x, y)
				nextNodeName := fmt.Sprintf("%d,%d", x, y+1)
				//fmt.Printf("Weight from %s to %s: %d\n", curNodeName, nextNodeName, weight)
				graph.AddEdge(nodes[curNodeName], nodes[nextNodeName], 1)
			}
			if x-1 >= 0 {
				curNodeName := fmt.Sprintf("%d,%d", x, y)
				nextNodeName := fmt.Sprintf("%d,%d", x-1, y)
				//fmt.Printf("Weight from %s to %s: %d\n", curNodeName, nextNodeName, weight)
				graph.AddEdge(nodes[curNodeName], nodes[nextNodeName], 1)
			}
			if x+1 < len(intMap[y]) {
				curNodeName := fmt.Sprintf("%d,%d", x, y)
				nextNodeName := fmt.Sprintf("%d,%d", x+1, y)
				//fmt.Printf("Weight from %s to %s: %d\n", curNodeName, nextNodeName, weight)
				graph.AddEdge(nodes[curNodeName], nodes[nextNodeName], 1)
			}
		}
	}

	return graph
}

func dijkstras(graph *WeightedGraph, startingNodeName string, destNodeName string) {
	visited := make(map[string]bool)
	heap := &Heap{}

	startNode := graph.GetNode(startingNodeName)
	startNode.distance = 0
	heap.Push(startNode)

	for heap.Size() > 0 {
		current := heap.Pop()
		visited[current.name] = true
		edges := graph.Edges[current.name]
		for _, edge := range edges {
			// Modified to not allow weights over 1
			if !visited[edge.node.name] && edge.weight <= 1 {
				//fmt.Printf("Visited: %s\n", edge.node.name)
				heap.Push(edge.node)
				if current.distance+1 < edge.node.distance {
					edge.node.distance = current.distance + 1
					edge.node.through = current
				}
			}
		}
	}
}

func drawMap(filename string, finalNode *Node) {
	array := read.ReadStrArrayByLine(filename)

	var drawnMap [][]string
	for y := 0; y < len(array); y++ {
		curLine := read.StrToCharArray(array[y])
		var line []string
		for x := 0; x < len(curLine); x++ {
			line = append(line, curLine[x])
		}
		drawnMap = append(drawnMap, line)
	}

	for n := finalNode; n.through != nil; n = n.through {
		split := strings.Split(n.name, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		existingChar := drawnMap[y][x]
		drawnMap[y][x] = color.HiGreenString(existingChar)
	}

	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	for _, line := range drawnMap {
		var curline []string
		for _, char := range line {
			curline = append(curline, char)
		}
		fmt.Printf("%s\n", curline)
	}
}

func part1() {
	//Example Input
	//startNodeName := "0,0"
	//destNodeName := "5,2"

	//Real Input
	startNodeName := "0,20"
	destNodeName := "43,20"

	intMap := processInput("./input.txt")
	graph := setupGraph(intMap, startNodeName, destNodeName)

	dijkstras(graph, startNodeName, destNodeName)

	steps := 0
	var finalNode *Node
	for _, node := range graph.Nodes {
		if node.name == destNodeName {
			finalNode = node
			//fmt.Printf("Shortest weight from %s to %s is %d\n", "0,0", node.name, node.distance)
			for n := node; n.through != nil; n = n.through {
				fmt.Print(n.name, " <- ")
				steps += 1
			}
			fmt.Println(startNodeName)
		}
	}

	drawMap("./input.txt", finalNode)

	fmt.Printf("Part 1 - Steps: %v\n", steps)
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
	fmt.Printf("Part 1 Time: %d??s\nPart 2 Time: %d??s\n", part1Time.Microseconds(), part2Time.Microseconds())
}
