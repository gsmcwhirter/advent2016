package day15

import (
	"fmt"
	"strings"

	"regexp"
	"strconv"

	"github.com/gsmcwhirter/advent2016/lib"
)

type Disc struct {
	NumSlots    int
	CurrentSlot int
}

func (d *Disc) SlotAfterRotations(steps int) int {
	return (d.CurrentSlot + steps) % d.NumSlots
}

type State struct {
	Discs []Disc
}

func (s *State) DropSuccessful(time int) bool {
	for i, disc := range s.Discs {
		if disc.SlotAfterRotations(time+i+1) > 0 {
			return false
		}
	}

	return true
}

func ParseDataString(dat string) State {
	discs := []Disc{}

	r := regexp.MustCompile("has (\\d+) positions; at time=0, it is at position (\\d+)")
	for _, line := range strings.Split(dat, "\n") {
		matches := r.FindStringSubmatch(line)
		numSlots, _ := strconv.Atoi(matches[1])
		pos, _ := strconv.Atoi(matches[2])
		discs = append(discs, Disc{
			NumSlots:    numSlots,
			CurrentSlot: pos,
		})
	}

	return State{discs}
}

func LoadData(filename string) State {
	dat := lib.ReadFileData(filename)

	return ParseDataString(strings.Trim(string(dat), "\n"))
}

func RunPartA(filename string) {
	state := LoadData(filename)

	time := 0
	for {
		if state.DropSuccessful(time) {
			fmt.Println(time)
			break
		}
		time++
	}
}

func RunPartB(filename string) {
	state := LoadData(filename)

	time := 0
	for {
		if state.DropSuccessful(time) {
			fmt.Println(time)
			break
		}

		time++
	}
}
