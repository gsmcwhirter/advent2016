package day19

import (
	"fmt"
	"strings"

	"strconv"

	"github.com/gsmcwhirter/advent2016/lib"
)

type CircularBuffer struct {
	Next      *CircularBuffer
	ID        int
	Presents  int
	StartSize int
}

func ParseDataString(dat string) *CircularBuffer {
	num, _ := strconv.Atoi(dat)

	first := &CircularBuffer{
		Next:      nil,
		ID:        1,
		Presents:  1,
		StartSize: num,
	}

	prev := first
	for i := num; i > 1; i-- {
		curr := &CircularBuffer{
			Next:      prev,
			ID:        i,
			Presents:  1,
			StartSize: num,
		}
		prev = curr
	}
	first.Next = prev

	return first
}

func (s *CircularBuffer) Steal() *CircularBuffer {
	if s.Next.ID == s.ID {
		return s
	}

	// fmt.Printf("%v stealing %v's presents\n", s.ID, s.Next.ID)
	s.Presents += s.Next.Presents
	oldNext := s.Next
	s.Next = s.Next.Next
	oldNext.Next = nil

	return s.Next.Steal()
}

func (s *CircularBuffer) Steal2(prevPreMid *CircularBuffer, skip2 bool) (*CircularBuffer, *CircularBuffer) {
	if s.Next.ID == s.ID {
		return s, nil
	}

	preTarget := prevPreMid
	target := prevPreMid.Next
	if skip2 {
		preTarget = target
		target = target.Next
	}

	// fmt.Printf("%v stealing %v's presents\n", s.ID, target.ID)
	postTarget := target.Next

	s.Presents += target.Presents

	preTarget.Next = postTarget
	target.Next = nil

	return s.Next, preTarget
}

func LoadData(filename string) *CircularBuffer {
	dat := lib.ReadFileData(filename)

	return ParseDataString(strings.Trim(string(dat), "\n"))
}

func RunPartA(filename string) {
	start := LoadData(filename)

	winner := start.Steal()

	fmt.Println(winner.ID)
}

func RunPartB(filename string) {
	state := LoadData(filename)

	skip2 := state.StartSize%2 == 0
	// Find the location just before the midpoint
	mod := 1
	if skip2 {
		mod = 2
	}
	prevMid := state
	for i := 0; i < state.StartSize/2-mod; i++ {
		prevMid = prevMid.Next
	}

	for {
		state, prevMid = state.Steal2(prevMid, skip2)
		skip2 = !skip2
		if prevMid == nil {
			break
		}
	}

	fmt.Println(state.ID)
}
