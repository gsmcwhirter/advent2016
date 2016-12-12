package day2

import (
	"fmt"
	"strings"

	"github.com/gsmcwhirter/advent2016/lib"
)

type Delta struct {
	x int
	y int
}

func DeltaFromRune(dir string) Delta {
	switch dir {
	case "U":
		return Delta{0, -1}
	case "D":
		return Delta{0, 1}
	case "L":
		return Delta{-1, 0}
	case "R":
		return Delta{1, 0}
	default:
		return Delta{0, 0}
	}
}

type Step struct {
	Moves []string
}

func (step *Step) deltas() []Delta {
	deltas := []Delta{}
	for _, move := range step.Moves {
		delta := DeltaFromRune(move)
		deltas = append(deltas, delta)
	}
	return deltas
}

type Key struct {
	x int
	y int
}

func (key *Key) value() int {
	return 3*key.y + key.x + 1
}

func numInRow(y int) int {
	return 2*(2-lib.IntAbs(2-y)) + 1
}

func minIndexInRow(y int) int {
	return lib.IntAbs(2 - y)
}

func maxIndexInRow(y int) int {
	return 4 - lib.IntAbs(2-y)
}

func (key *Key) value2() string {
	intVal := 0
	for i := 0; i < key.y; i++ {
		intVal += numInRow(i)
	}
	intVal += key.x - minIndexInRow(key.y) + 1
	// intVal := 2*key.y + 1 + IntAbs((2-key.y)-key.x)
	if intVal < 10 {
		return fmt.Sprintf("%v", intVal)
	}

	return string(byte(int([]byte("A")[0]) + (intVal - 10)))
}

func (key *Key) move(step Step) Key {
	nextKey := Key{key.x, key.y}
	for _, delta := range step.deltas() {
		nextKey.applyDelta(delta)
	}

	return nextKey
}

func (key *Key) move2(step Step) Key {
	nextKey := Key{key.x, key.y}
	for _, delta := range step.deltas() {
		nextKey.applyDelta2(delta)
	}

	return nextKey
}

func (key *Key) applyDelta(delta Delta) {
	key.x += delta.x
	key.x = lib.IntClamp(key.x, 0, 2)

	key.y += delta.y
	key.y = lib.IntClamp(key.y, 0, 2)
}

func (key *Key) applyDelta2(delta Delta) {
	key.x += delta.x
	key.x = lib.IntClamp(key.x, minIndexInRow(key.y), maxIndexInRow(key.y))

	key.y += delta.y
	key.y = lib.IntClamp(key.y, minIndexInRow(key.x), maxIndexInRow(key.x))

	// fmt.Printf("Applying %v (->%v)\n", delta, key)
}

//ParseDataString splits on spaces
func ParseDataString(datString string) []Step {
	lines := []Step{}
	for _, rawLine := range strings.Split(datString, "\n") {
		runes := []rune(rawLine)
		chars := []string{}
		for _, r := range runes {
			chars = append(chars, string(r))
		}
		step := Step{chars}
		lines = append(lines, step)
	}

	return lines
}

//LoadData stuff
func LoadData(filename string) []Step {
	dat := lib.ReadFileData(filename)

	return ParseDataString(string(dat))
}

func RunPartA(filename string) {
	steps := LoadData(filename)
	key := Key{1, 1}

	for _, step := range steps {
		// fmt.Printf("Step: %v\n", step)
		key = key.move(step)
		fmt.Printf("Key: %v\n", key.value())
	}
}

func RunPartB(filename string) {
	steps := LoadData(filename)
	key := Key{0, 2}

	// fmt.Printf("Start Key: %v (%v)\n", key.value2(), key)
	for _, step := range steps {
		// fmt.Printf("Step: %v\n", step)
		key = key.move2(step)
		fmt.Printf("Key: %v (%v)\n", key.value2(), key)
	}
}
