package day17

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/gsmcwhirter/advent2016/lib"
)

type Location struct {
	X int
	Y int
}

func (l *Location) Copy() Location {
	return Location{l.X, l.Y}
}

func (l *Location) Equals(other Location) bool {
	return l.X == other.X && l.Y == other.Y
}

type State struct {
	Passcode string
	Location Location
	Moves    []rune
}

func NewState(passcode string) State {
	return State{
		Passcode: passcode,
		Location: Location{1, 1},
		Moves:    []rune{},
	}
}

func (s *State) ValidMoves() []rune {
	moves := []rune{}
	hash := []rune(fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%v%v", s.Passcode, string(s.Moves))))))
	for i, r := range hash[:4] {
		switch r {
		case 'b', 'c', 'd', 'e', 'f':
			switch i {
			case 0:
				if s.Location.Y > 1 {
					moves = append(moves, 'U')
				}
			case 1:
				if s.Location.Y < 4 {
					moves = append(moves, 'D')
				}
			case 2:
				if s.Location.X > 1 {
					moves = append(moves, 'L')
				}
			case 3:
				if s.Location.X < 4 {
					moves = append(moves, 'R')
				}
			}
		}
	}

	return moves
}

func (s *State) Copy() State {
	newS := NewState(s.Passcode)
	newS.Location = s.Location.Copy()
	newS.Moves = append([]rune{}, s.Moves...)

	return newS
}

func (s *State) ApplyMove(move rune) {
	switch move {
	case 'U':
		s.Location.Y--
	case 'D':
		s.Location.Y++
	case 'L':
		s.Location.X--
	case 'R':
		s.Location.X++
	}
	s.Moves = append(s.Moves, move)
}

func (s *State) IsAt(target Location) bool {
	return s.Location.Equals(target)
}

type StateQueue struct {
	Items []*State
}

func (q *StateQueue) Push(state State) {
	newState := state.Copy()
	q.Items = append(q.Items, &newState)
}

func (q *StateQueue) Pop() *State {
	if len(q.Items) == 0 {
		return nil
	}

	item := q.Items[0]
	q.Items = q.Items[1:]

	return item
}

func BFS(start State, target Location, stopAtFirst bool) {
	queue := StateQueue{[]*State{}}
	queue.Push(start)

	var nextState State
	for {
		cState := queue.Pop()

		if cState == nil {
			break
		}

		if cState.IsAt(target) {
			if stopAtFirst {
				fmt.Println(string(cState.Moves))
				break
			} else {
				fmt.Println(len(cState.Moves))
				continue
			}
		}

		for _, move := range cState.ValidMoves() {
			nextState = cState.Copy()
			nextState.ApplyMove(move)
			queue.Push(nextState)
		}
	}
}

func LoadData(filename string) string {
	dat := lib.ReadFileData(filename)

	return strings.Trim(string(dat), "\n")
}

func RunPartA(filename string) {
	passcode := LoadData(filename)
	start := NewState(passcode)

	BFS(start, Location{4, 4}, true)
}

func RunPartB(filename string) {
	passcode := LoadData(filename)
	start := NewState(passcode)

	BFS(start, Location{4, 4}, false)
}
