package read

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func IntSliceInsert(targetStack []int, num int, pos int) []int {
	targetStack = append(targetStack, 0)
	copy(targetStack[pos+1:], targetStack[pos:])
	targetStack[pos] = num
	return targetStack
}

func IntSlicePop(targetStack []int) ([]int, int, bool) {
	if len(targetStack) == 0 {
		return targetStack, 0, false
	} else {
		index := len(targetStack) - 1       // Get the index of the top most element.
		element := (targetStack)[index]     // Index into the slice and obtain the element.
		targetStack = (targetStack)[:index] // Remove it from the stack by slicing it off.
		return targetStack, element, true
	}
}

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *Stack) Insert(str string, pos int) {
	var stk []string
	stk = *s
	stk = append(stk, "")
	copy(stk[pos+1:], stk[pos:])
	stk[pos] = str
	*s = stk
}

func SliceContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func ReadIntArrayByLine(filepath string) []int {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var intarray []int
	for scanner.Scan() {
		cur, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		intarray = append(intarray, cur)
	}

	return intarray
}

func ReadStrArrayByLine(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var strarray []string
	for scanner.Scan() {
		cur := scanner.Text()
		strarray = append(strarray, cur)
	}

	return strarray
}

func StrToCharArray(str string) []string {
	runes := []rune(str)
	var chars []string
	for i := 0; i < len(runes); i++ {
		chars = append(chars, string(runes[i]))
	}
	return chars
}

func RemoveFromStringSlice(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func IntArrayToString(intarray []int) string {
	result := ""
	for _, num := range intarray {
		result = result + strconv.Itoa(num)
	}
	return result
}

func CharArrayToIntArray(str []string) []int {
	var ints []int
	for i := 0; i < len(str); i++ {
		cur, _ := strconv.Atoi(str[i])
		ints = append(ints, cur)
	}
	return ints
}

func IntToIntArray(number int) []int {
	numstr := strconv.Itoa(number)
	//fmt.Printf("%#v", numstr)
	strarray := strings.Split(numstr, "")
	//fmt.Printf("%#v", strarray)
	var intarray []int
	for _, i := range strarray {
		j, _ := strconv.Atoi(i)
		intarray = append(intarray, j)
	}
	return intarray
}

func StrToWordArray(str string) []string {
	words := strings.Split(str, " ")
	return words
}
