package day18

import (
	"fmt"
	"strings"

	"strconv"

	"github.com/gsmcwhirter/advent2016/lib"
)

func IsSafe(oldState []rune, newSquare int) bool {
	var leftSafe, rightSafe bool

	if newSquare-1 < 0 || oldState[newSquare-1] == '.' {
		leftSafe = true
	} else {
		leftSafe = false
	}

	if newSquare+1 >= len(oldState) || oldState[newSquare+1] == '.' {
		rightSafe = true
	} else {
		rightSafe = false
	}

	if leftSafe && !rightSafe {
		return false
	}

	if rightSafe && !leftSafe {
		return false
	}

	return true

}

func AdvanceState(state []rune) []rune {
	nextState := make([]rune, len(state))

	for i := 0; i < len(nextState); i++ {
		if IsSafe(state, i) {
			nextState[i] = '.'
		} else {
			nextState[i] = '^'
		}
	}

	return nextState
}

func CountSafe(state []rune) int {
	ct := 0
	for i := 0; i < len(state); i++ {
		if state[i] == '.' {
			ct++
		}
	}

	return ct
}

func ParseDataString(dat string) (int, []rune) {
	lines := strings.Split(dat, "\n")
	iterCt, _ := strconv.Atoi(lines[0])

	return iterCt, []rune(lines[1])
}

func LoadData(filename string) (int, []rune) {
	dat := lib.ReadFileData(filename)

	return ParseDataString(strings.Trim(string(dat), "\n"))
}

func RunPartA(filename string) {
	iterCt, state := LoadData(filename)

	ct := 0
	for i := 0; i < iterCt; i++ {
		fmt.Println(string(state))
		ct += CountSafe(state)
		state = AdvanceState(state)

	}

	fmt.Println(ct)
}

func RunPartB(filename string) {
	iterCt, state := LoadData(filename)

	ct := 0
	for i := 0; i < iterCt; i++ {
		// fmt.Println(string(state))
		ct += CountSafe(state)
		state = AdvanceState(state)

	}

	fmt.Println(ct)
}
