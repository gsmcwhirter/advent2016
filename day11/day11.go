package day11

import (
	"fmt"
	"strings"

	"regexp"

	"sort"

	"github.com/gsmcwhirter/advent2016/lib"
)

type Piece struct {
	Element string
	Type    string
}

type Move struct {
	TargetFloor int
	Pieces      []Piece
}

func (m *Move) HasPiece(p Piece) bool {
	for _, piece := range m.Pieces {
		if piece.Element == p.Element && piece.Type == p.Type {
			return true
		}
	}

	return false
}

func (m *Move) GeneratorTypes() []string {
	types := []string{}
	for _, piece := range m.Pieces {
		if piece.Type == "generator" {
			types = append(types, piece.Element)
		}
	}

	return types
}

func (m *Move) ChipTypes() []string {
	types := []string{}
	for _, piece := range m.Pieces {
		if piece.Type == "chip" {
			types = append(types, piece.Element)
		}
	}

	return types
}

type Floor struct {
	Num    int
	Pieces []Piece
}

func (f *Floor) Equal(other Floor) bool {
	if f.Num != other.Num {
		return false
	}

	for _, piece := range append(f.Pieces, other.Pieces...) {
		if !f.HasPiece(piece) || !other.HasPiece(piece) {
			return false
		}
	}

	return true
}

func (f *Floor) HasPiece(p Piece) bool {
	for _, piece := range f.Pieces {
		if piece.Element == p.Element && piece.Type == p.Type {
			return true
		}
	}

	return false
}

func (f *Floor) HasGenerator(elem string) bool {
	for _, piece := range f.Pieces {
		if piece.Element == elem && piece.Type == "generator" {
			return true
		}
	}

	return false
}

func (f *Floor) GeneratorTypes(excepts ...string) []string {
	types := []string{}
	for _, piece := range f.Pieces {
		if piece.Type == "generator" {
			toAdd := true
			for _, except := range excepts {
				if except == piece.Element {
					toAdd = false
				}
			}

			if toAdd {
				types = append(types, piece.Element)
			}
		}
	}

	return types
}

func (f *Floor) HasChip(elem string) bool {
	for _, piece := range f.Pieces {
		if piece.Element == elem && piece.Type == "chip" {
			return true
		}
	}

	return false
}

func (f *Floor) ChipTypes(excepts ...string) []string {
	types := []string{}
	for _, piece := range f.Pieces {
		if piece.Type == "chip" {
			toAdd := true
			for _, except := range excepts {
				if except == piece.Element {
					toAdd = false
				}
			}

			if toAdd {
				types = append(types, piece.Element)
			}
		}
	}

	return types
}

func (f *Floor) NumPairs() int {
	gens := f.GeneratorTypes("")

	count := 0
	for _, gen := range gens {
		if f.HasChip(gen) {
			count++
		}
	}

	return count
}

func (f *Floor) NumUnpairedChips() int {
	chips := f.ChipTypes("")

	count := 0
	for _, chip := range chips {
		if !f.HasGenerator(chip) {
			count++
		}
	}

	return count
}

func (f *Floor) NumUnpairedGenerators() int {
	gens := f.GeneratorTypes("")

	count := 0
	for _, gen := range gens {
		if !f.HasChip(gen) {
			count++
		}
	}

	return count
}

type State struct {
	Elevator  int
	Floors    []Floor
	Elements  []string
	MoveCount int
}

func (s *State) GenerateMoves(history []State) []Move {
	moves := []Move{}

	floor := s.Floors[s.Elevator-1]

	var move Move
	for i, piece := range floor.Pieces {
		move = Move{s.Elevator + 1, []Piece{piece}}
		if s.MoveIsValid(move, history) {
			moves = append(moves, move)
		}

		move = Move{s.Elevator - 1, []Piece{piece}}
		if s.MoveIsValid(move, history) {
			moves = append(moves, move)
		}

		for j, piece2 := range floor.Pieces {
			if i == j {
				continue
			}

			move = Move{s.Elevator + 1, []Piece{piece, piece2}}
			if s.MoveIsValid(move, history) {
				moves = append(moves, move)
			}

			move = Move{s.Elevator - 1, []Piece{piece, piece2}}
			if s.MoveIsValid(move, history) {
				moves = append(moves, move)
			}
		}
	}

	return moves
}

func (s *State) MoveIsValid(move Move, history []State) bool {
	// Are we moving within bounds?
	if move.TargetFloor > 4 || move.TargetFloor < 1 {
		return false
	}

	// Are we moving enough things?
	if len(move.Pieces) < 1 || len(move.Pieces) > 2 {
		return false
	}

	currentFloor := s.Floors[s.Elevator-1]
	targetFloor := s.Floors[move.TargetFloor-1]

	moveGenerators := move.GeneratorTypes()
	moveChips := move.ChipTypes()

	// Can we move these things together?
	if len(moveGenerators) > 0 && len(moveChips) > 0 {
		for _, gen := range moveGenerators {
			for _, chip := range moveChips {
				if gen == chip {
					return true
				}
			}
		}

		return false
	}

	//Will removing these things fry the current floor?
	currentFloorRemainingGenerators := currentFloor.GeneratorTypes(moveGenerators...)
	for _, chip := range currentFloor.ChipTypes(moveChips...) {
		if !currentFloor.HasGenerator(chip) && len(currentFloorRemainingGenerators) > 0 {
			return false
		}
	}

	// Will we fry something on the target floor?
	for _, gen := range moveGenerators {
		if len(targetFloor.ChipTypes(gen)) > 0 {
			return false
		}
	}

	// Will we fry a chip in the elevator on the target floor?
	for _, chip := range moveChips {
		if !targetFloor.HasGenerator(chip) && len(targetFloor.GeneratorTypes(chip)) > 0 {
			return false
		}
	}

	newState := AdvanceState(*s, move)
	if StateInHistory(history, newState) {
		return false
	}

	return true
}

func (s *State) Print() {
	var floor Floor
	for floorNum := len(s.Floors); floorNum > 0; floorNum-- {
		floor = s.Floors[floorNum-1]
		fmt.Printf("F%v", floorNum)
		if s.Elevator == floorNum {
			fmt.Print(" E")
		} else {
			fmt.Print(" .")
		}

		for _, elem := range s.Elements {
			if floor.HasGenerator(elem) {
				fmt.Printf(" %v-G", string([]rune(elem)[0]))
			} else {
				fmt.Print("  . ")
			}

			if floor.HasChip(elem) {
				fmt.Printf(" %v-C", string([]rune(elem)[0]))
			} else {
				fmt.Print("  . ")
			}
		}

		fmt.Println()
	}
	fmt.Printf("\nMoves: %v\n\n", s.MoveCount)
}

func (s *State) GoalReached() bool {
	for _, floor := range s.Floors[:3] {
		if len(floor.Pieces) > 0 {
			return false
		}
	}

	return true
}

func (s *State) Equal(other State) bool {
	if s.Elevator != other.Elevator {
		return false
	}

	for i, sflr := range s.Floors {
		oflr := other.Floors[i]

		if !sflr.Equal(oflr) {
			return false
		}
	}

	return true
}

func (s *State) Isomorphic(other State) bool {
	if s.Elevator != other.Elevator {
		return false
	}

	for i, sflr := range s.Floors {
		oflr := other.Floors[i]

		if len(sflr.Pieces) != len(oflr.Pieces) {
			return false
		}

		if sflr.NumPairs() != oflr.NumPairs() {
			return false
		}

		if sflr.NumUnpairedChips() != oflr.NumUnpairedChips() {
			return false
		}

		// don't need to check unpaired generators at this point -- they must match
	}

	return true
}

func StateInHistory(history []State, ns State) bool {
	for _, oldS := range history {
		if oldS.Isomorphic(ns) { // was Equal
			return true
		}
	}

	return false
}

func AdvanceState(oldState State, move Move) State {
	// oldState.Print()

	newFloors := []Floor{}

	for i, floor := range oldState.Floors {
		// fmt.Print(i, " ")
		// fmt.Println(floor)
		switch i + 1 {
		case oldState.Elevator:
			// fmt.Println("Cleaning out old floor")
			newPieces := []Piece{}
			for _, piece := range floor.Pieces {
				if !move.HasPiece(piece) {
					newPieces = append(newPieces, piece)
				}
			}

			newFloors = append(newFloors, Floor{
				Num:    floor.Num,
				Pieces: newPieces,
			})
		case move.TargetFloor:
			// fmt.Println("Creating new floor")
			newFloors = append(newFloors, Floor{
				Num:    floor.Num,
				Pieces: append(floor.Pieces[:], move.Pieces...),
			})
		default:
			// fmt.Println("Floor unchanged")
			newFloors = append(newFloors, Floor{
				Num:    floor.Num,
				Pieces: floor.Pieces[:],
			})
		}
	}

	newState := State{
		Elevator:  move.TargetFloor,
		Floors:    newFloors,
		Elements:  oldState.Elements,
		MoveCount: oldState.MoveCount + 1,
	}

	// newState.Print()

	return newState
}

type MoveStatePair struct {
	Move  Move
	State State
}

type MoveStateQueue struct {
	Items []*MoveStatePair
}

func (q *MoveStateQueue) push(move Move, state State) {
	msp := &MoveStatePair{move, state}
	q.Items = append(q.Items, msp)
}

func (q *MoveStateQueue) pop() *MoveStatePair {
	if len(q.Items) == 0 {
		return nil
	}

	item := q.Items[0]
	q.Items = q.Items[1:]

	return item
}

func BFS(startState State) {
	state := startState
	queue := MoveStateQueue{[]*MoveStatePair{}}
	history := []State{state}

	newMoves := state.GenerateMoves(history)
	// fmt.Printf("Adding %v moves to the queue (current length %v)\n", len(newMoves), len(newMoves)+len(queue.Items))
	for _, move := range newMoves {
		queue.push(move, state)
	}

	item := queue.pop()
	var newState State
	for item != nil {
		// fmt.Println(item.Move)
		newState = AdvanceState(item.State, item.Move)
		// newState.Print()

		if StateInHistory(history, newState) {
			// fmt.Println("Skipping state")
			item = queue.pop()
			continue
		} else {
			history = append(history, newState)
		}

		if newState.GoalReached() {
			newState.Print()
			break
		}

		newMoves := newState.GenerateMoves(history)
		// fmt.Printf("Adding %v moves to the queue (current length %v)\n", len(newMoves), len(newMoves)+len(queue.Items))
		for _, move := range newMoves {
			queue.push(move, newState)
		}

		item = queue.pop()
	}
}

func ParseDataString(dat string) ([]Floor, []string) {
	floors := []Floor{}
	typesMap := map[string]bool{}

	generatorRegex := regexp.MustCompile("a ([^\\s]+) generator")
	chipRegex := regexp.MustCompile("a ([^\\s]+)-compatible microchip")

	for num, line := range strings.Split(dat, "\n") {
		pieces := []Piece{}
		generatorMatches := generatorRegex.FindAllStringSubmatch(line, -1)
		for _, match := range generatorMatches {
			pieces = append(pieces, Piece{
				Element: match[1],
				Type:    "generator",
			})
			typesMap[match[1]] = true
		}

		chipMatches := chipRegex.FindAllStringSubmatch(line, -1)
		for _, match := range chipMatches {
			pieces = append(pieces, Piece{
				Element: match[1],
				Type:    "chip",
			})
			typesMap[match[1]] = true
		}

		floors = append(floors, Floor{
			Num:    num + 1,
			Pieces: pieces,
		})
	}

	types := make([]string, len(typesMap))
	i := -1

	for key := range typesMap {
		i++
		types[i] = key
	}

	sort.Strings(types)

	return floors, types
}

func LoadData(filename string) ([]Floor, []string) {
	dat := lib.ReadFileData(filename)

	return ParseDataString(strings.Trim(string(dat), "\n"))
}

func RunPartA(filename string) {
	floors, allTypes := LoadData(filename)

	state := State{
		Elevator:  1,
		Floors:    floors,
		Elements:  allTypes,
		MoveCount: 0,
	}

	BFS(state)
}

func RunPartB(filename string) {

}
