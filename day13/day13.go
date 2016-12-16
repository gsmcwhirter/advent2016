package day13

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gsmcwhirter/advent2016/lib"
	"github.com/steakknife/hamming"
)

type Location struct {
	X     int64
	Y     int64
	Steps int
}

func (l *Location) Copy() Location {
	return Location{
		l.X,
		l.Y,
		l.Steps,
	}
}

func (l *Location) IsWall(favNum int64) bool {
	val := uint64(l.X*l.X + 3*l.X + 2*l.X*l.Y + l.Y + l.Y*l.Y + favNum)

	bitsOn := hamming.CountBitsUint64(val)

	if bitsOn%2 == 1 {
		return true
	}

	return false
}

func (l *Location) ValidAdjacentLocs(favNum int64, history []Location) []Location {
	locs := []Location{}

	var nextLoc Location

	nextLoc = Location{l.X + 1, l.Y, l.Steps + 1}
	if !nextLoc.IsWall(favNum) && !LocationInHistory(history, nextLoc) {
		locs = append(locs, nextLoc)
	}

	if l.X > 0 {
		nextLoc = Location{l.X - 1, l.Y, l.Steps + 1}
		if !nextLoc.IsWall(favNum) && !LocationInHistory(history, nextLoc) {
			locs = append(locs, nextLoc)
		}
	}

	nextLoc = Location{l.X, l.Y + 1, l.Steps + 1}
	if !nextLoc.IsWall(favNum) && !LocationInHistory(history, nextLoc) {
		locs = append(locs, nextLoc)
	}

	if l.Y > 0 {
		nextLoc = Location{l.X, l.Y - 1, l.Steps + 1}
		if !nextLoc.IsWall(favNum) && !LocationInHistory(history, nextLoc) {
			locs = append(locs, nextLoc)
		}
	}

	return locs
}

func (l *Location) Equals(other Location) bool {
	if other.X == l.X && other.Y == l.Y {
		return true
	}

	return false
}

func LocationInHistory(history []Location, loc Location) bool {
	for _, hloc := range history {
		if hloc.Equals(loc) {
			return true
		}
	}

	return false
}

type LocationQueue struct {
	Items []*Location
}

func (q *LocationQueue) Push(loc Location) {
	newLoc := loc.Copy()
	q.Items = append(q.Items, &newLoc)
}

func (q *LocationQueue) Pop() *Location {
	if len(q.Items) == 0 {
		return nil
	}

	item := q.Items[0]
	q.Items = q.Items[1:]

	return item
}

func PrintGrid(w, h int, favNum int64, history []Location) {
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			loc := Location{int64(c), int64(r), 0}
			if loc.IsWall(favNum) {
				fmt.Print("#")
			} else {
				if LocationInHistory(history, loc) {
					fmt.Print(".")
				} else {
					fmt.Print(" ")
				}
			}
		}

		fmt.Println()
	}
}

func BFS(startLoc Location, targetLoc Location, favNum int64) {
	queue := LocationQueue{[]*Location{}}
	history := []Location{startLoc}

	newMoves := startLoc.ValidAdjacentLocs(favNum, history)
	fmt.Printf("Adding %v moves to the queue (current length %v)\n", len(newMoves), len(newMoves)+len(queue.Items))
	for _, move := range newMoves {
		// fmt.Println(move)
		queue.Push(move)
		history = append(history, move)
	}

	lastMoveCount := 0

	newLoc := queue.Pop()
	for newLoc != nil {
		// newState.Print()

		if newLoc.Steps > lastMoveCount {
			lastMoveCount = newLoc.Steps
			fmt.Printf("Considering states %v moves out\n", lastMoveCount)
		}

		if targetLoc.Equals(*newLoc) {
			fmt.Println(newLoc.Steps)
			break
		}

		newMoves := newLoc.ValidAdjacentLocs(favNum, history)
		fmt.Printf("Adding %v moves to the queue (current length %v)\n", len(newMoves), len(newMoves)+len(queue.Items))
		for _, move := range newMoves {
			queue.Push(move)
			history = append(history, move)
		}

		newLoc = queue.Pop()
	}
}

func BFS2(startLoc Location, favNum int64) []Location {
	queue := LocationQueue{[]*Location{}}
	history := []Location{startLoc}

	newMoves := startLoc.ValidAdjacentLocs(favNum, history)
	fmt.Printf("Adding %v moves to the queue (current length %v)\n", len(newMoves), len(newMoves)+len(queue.Items))
	for _, move := range newMoves {
		// fmt.Println(move)
		queue.Push(move)
		history = append(history, move)
	}

	lastMoveCount := 0

	newLoc := queue.Pop()
	for newLoc != nil {
		// newState.Print()

		if newLoc.Steps > lastMoveCount {
			lastMoveCount = newLoc.Steps
			fmt.Printf("Considering states %v moves out\n", lastMoveCount)
		}

		if lastMoveCount >= 51 {
			fmt.Println(history)
			fmt.Println(len(history))
			break
		}

		newMoves := newLoc.ValidAdjacentLocs(favNum, history)
		fmt.Printf("Adding %v moves to the queue (current length %v)\n", len(newMoves), len(newMoves)+len(queue.Items))
		for _, move := range newMoves {
			queue.Push(move)
			history = append(history, move)
		}

		newLoc = queue.Pop()
	}

	return history
}

func LoadData(filename string) int64 {
	dat := lib.ReadFileData(filename)

	val, _ := strconv.Atoi(strings.Trim(string(dat), "\n"))
	return int64(val)
}

func RunPartA(filename string) {
	favNum := LoadData(filename)
	start := Location{1, 1, 0}
	target := Location{31, 39, -1}

	BFS(start, target, favNum)
}

func RunPartB(filename string) {
	favNum := LoadData(filename)
	start := Location{1, 1, 0}

	history := BFS2(start, favNum)
	PrintGrid(50, 50, favNum, history)
}
