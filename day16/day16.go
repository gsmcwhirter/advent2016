package day16

import (
	"fmt"
	"strings"

	"github.com/gsmcwhirter/advent2016/lib"
)

func LoadData(filename string) []rune {
	dat := lib.ReadFileData(filename)

	return []rune(strings.Trim(string(dat), "\n"))
}

func AdvanceState(state []rune) []rune {
	newState := make([]rune, len(state)*2+1)

	var j int
	for i := 0; i < len(state); i++ {
		j = 2*len(state) - i

		newState[i] = state[i]
		if state[i] == '1' {
			newState[j] = '0'
		} else {
			newState[j] = '1'
		}
	}

	newState[len(state)] = '0'

	return newState
}

func CheckSum(state []rune) []rune {
	chksum := make([]rune, len(state)/2)

	for i := 0; i < len(chksum); i++ {
		if state[2*i] == state[2*i+1] {
			chksum[i] = '1'
		} else {
			chksum[i] = '0'
		}
	}

	return chksum
}

func RunPartA(filename string) {
	state := LoadData(filename)

	for len(state) < 272 {
		state = AdvanceState(state)
	}

	chksum := CheckSum(state[:272])

	for len(chksum)%2 == 0 {
		chksum = CheckSum(chksum)
	}
	fmt.Println(string(chksum))
}

func RunPartB(filename string) {
	state := LoadData(filename)

	for len(state) < 35651584 {
		state = AdvanceState(state)
	}

	chksum := CheckSum(state[:35651584])

	for len(chksum)%2 == 0 {
		chksum = CheckSum(chksum)
	}
	fmt.Println(string(chksum))
}
