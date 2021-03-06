package day11

// import (
// 	"fmt"
// 	"strings"

// 	"regexp"

// 	"sort"

// 	"github.com/gsmcwhirter/advent2016/lib"
// )

// type Piece struct {
// 	Element string
// 	Type    string
// }

// type Move struct {
// 	TargetFloor int
// 	Pieces      []Piece
// }

// func (m *Move) HasPiece(p Piece) bool {
// 	for _, piece := range m.Pieces {
// 		if piece.Element == p.Element && piece.Type == p.Type {
// 			return true
// 		}
// 	}

// 	return false
// }

// func (m *Move) GeneratorTypes() []string {
// 	types := []string{}
// 	for _, piece := range m.Pieces {
// 		if piece.Type == "generator" {
// 			types = append(types, piece.Element)
// 		}
// 	}

// 	return types
// }

// func (m *Move) ChipTypes() []string {
// 	types := []string{}
// 	for _, piece := range m.Pieces {
// 		if piece.Type == "chip" {
// 			types = append(types, piece.Element)
// 		}
// 	}

// 	return types
// }

// type Floor struct {
// 	Num    int
// 	Pieces []Piece
// }

// func (f *Floor) Equal(other Floor) bool {
// 	if f.Num != other.Num {
// 		return false
// 	}

// 	allPieces := append(f.Pieces, other.Pieces...)
// 	for _, piece := range allPieces {
// 		if !f.HasPiece(piece) || !other.HasPiece(piece) {
// 			return false
// 		}
// 	}

// 	return true
// }

// func (f *Floor) HasPiece(p Piece) bool {
// 	for _, piece := range f.Pieces {
// 		if piece.Element == p.Element && piece.Type == p.Type {
// 			return true
// 		}
// 	}

// 	return false
// }

// func (f *Floor) HasGenerator(elem string) bool {
// 	for _, piece := range f.Pieces {
// 		if piece.Element == elem && piece.Type == "generator" {
// 			return true
// 		}
// 	}

// 	return false
// }

// func (f *Floor) GeneratorTypes(excepts ...string) []string {
// 	types := []string{}
// 	for _, piece := range f.Pieces {
// 		if piece.Type == "generator" {
// 			toAdd := true
// 			for _, except := range excepts {
// 				if except == piece.Element {
// 					toAdd = false
// 				}
// 			}

// 			if toAdd {
// 				types = append(types, piece.Element)
// 			}
// 		}
// 	}

// 	return types
// }

// func (f *Floor) HasChip(elem string) bool {
// 	for _, piece := range f.Pieces {
// 		if piece.Element == elem && piece.Type == "chip" {
// 			return true
// 		}
// 	}

// 	return false
// }

// func (f *Floor) ChipTypes(excepts ...string) []string {
// 	types := []string{}
// 	for _, piece := range f.Pieces {
// 		if piece.Type == "chip" {
// 			toAdd := true
// 			for _, except := range excepts {
// 				if except == piece.Element {
// 					toAdd = false
// 				}
// 			}

// 			if toAdd {
// 				types = append(types, piece.Element)
// 			}
// 		}
// 	}

// 	return types
// }

// func (f *Floor) NumPairs() int {
// 	gens := f.GeneratorTypes("")

// 	count := 0
// 	for _, gen := range gens {
// 		if f.HasChip(gen) {
// 			count++
// 		}
// 	}

// 	return count
// }

// func (f *Floor) NumUnpairedChips() int {
// 	chips := f.ChipTypes("")

// 	count := 0
// 	for _, chip := range chips {
// 		if !f.HasGenerator(chip) {
// 			count++
// 		}
// 	}

// 	return count
// }

// func (f *Floor) NumUnpairedGenerators() int {
// 	gens := f.GeneratorTypes("")

// 	count := 0
// 	for _, gen := range gens {
// 		if !f.HasChip(gen) {
// 			count++
// 		}
// 	}

// 	return count
// }

// type State struct {
// 	Elevator  int
// 	Floors    []Floor
// 	Elements  []string
// 	MoveCount int
// }

// func (s *State) GenerateMoves(history []State) []MoveStatePair {
// 	movesUp := []MoveStatePair{}
// 	movesDown := []MoveStatePair{}

// 	// minFloor := 0

// 	// for i := 0; i < 4; i++ {
// 	// 	if len(s.Floors[i].Pieces) == 0 {
// 	// 		minFloor++
// 	// 	} else {
// 	// 		break
// 	// 	}
// 	// }

// 	floor := s.Floors[s.Elevator-1]

// 	var move Move
// 	var valid bool
// 	var nextState State

// 	for i, piece := range floor.Pieces {
// 		for j, piece2 := range floor.Pieces {
// 			if i <= j {
// 				continue
// 			}

// 			move = Move{s.Elevator + 1, []Piece{piece, piece2}}
// 			valid, nextState = s.MoveIsValid(move, history)
// 			// fmt.Println(valid)
// 			// fmt.Println(nextState)
// 			if valid {
// 				movesUp = append(movesUp, MoveStatePair{move, nextState})
// 			}
// 		}

// 		// if s.Elevator > minFloor {
// 		move = Move{s.Elevator - 1, []Piece{piece}}
// 		valid, nextState = s.MoveIsValid(move, history)
// 		if valid {
// 			movesDown = append(movesDown, MoveStatePair{move, nextState})
// 		}
// 		// }

// 	}

// 	onePieceUp := len(movesUp) == 0 || true
// 	twoPieceDown := len(movesDown) == 0 || true

// 	for i, piece := range floor.Pieces {
// 		if onePieceUp {
// 			move = Move{s.Elevator + 1, []Piece{piece}}
// 			valid, nextState = s.MoveIsValid(move, history)
// 			if valid {
// 				movesUp = append(movesUp, MoveStatePair{move, nextState})
// 			}
// 		}

// 		if twoPieceDown {
// 			for j, piece2 := range floor.Pieces {
// 				if i <= j {
// 					continue
// 				}

// 				// if s.Elevator > minFloor {
// 				move = Move{s.Elevator - 1, []Piece{piece, piece2}}
// 				valid, nextState = s.MoveIsValid(move, history)
// 				if valid {
// 					movesDown = append(movesDown, MoveStatePair{move, nextState})
// 				}
// 				// }
// 			}
// 		}
// 	}

// 	moves := append(movesUp, movesDown...)
// 	// fmt.Println(moves)

// 	return moves
// }

// func (s *State) MoveIsValid(move Move, history []State) (bool, State) {
// 	// Are we moving within bounds?
// 	if move.TargetFloor > 4 || move.TargetFloor < 1 {
// 		return false, State{}
// 	}

// 	// Are we moving enough things?
// 	if len(move.Pieces) < 1 || len(move.Pieces) > 2 {
// 		return false, State{}
// 	}

// 	currentFloor := s.Floors[s.Elevator-1]
// 	targetFloor := s.Floors[move.TargetFloor-1]

// 	moveGenerators := move.GeneratorTypes()
// 	moveChips := move.ChipTypes()

// 	// Can we move these things together?
// 	if len(moveGenerators) > 0 && len(moveChips) > 0 {
// 		for _, gen := range moveGenerators {
// 			for _, chip := range moveChips {
// 				if gen == chip {
// 					return false, State{}
// 				}
// 			}
// 		}

// 		return false, State{}
// 	}

// 	//Will removing these things fry the current floor?
// 	currentFloorRemainingGenerators := currentFloor.GeneratorTypes(moveGenerators...)
// 	for _, chip := range currentFloor.ChipTypes(moveChips...) {
// 		if !currentFloor.HasGenerator(chip) && len(currentFloorRemainingGenerators) > 0 {
// 			return false, State{}
// 		}
// 	}

// 	// Will we fry something on the target floor?
// 	targetFloorOtherChips := targetFloor.ChipTypes(moveGenerators...)
// 	for _, chip := range targetFloorOtherChips {
// 		if !targetFloor.HasGenerator(chip) {
// 			return false, State{}
// 		}
// 	}

// 	// Will we fry a chip in the elevator on the target floor?
// 	for _, chip := range moveChips {
// 		if !targetFloor.HasGenerator(chip) && len(targetFloor.GeneratorTypes(chip)) > 0 {
// 			return false, State{}
// 		}
// 	}

// 	newState := AdvanceState(*s, move)
// 	if StateInHistory(history, newState) {
// 		return false, State{}
// 	}

// 	return true, newState
// }

// func (s *State) Print() {
// 	var floor Floor
// 	for floorNum := len(s.Floors); floorNum > 0; floorNum-- {
// 		floor = s.Floors[floorNum-1]
// 		fmt.Printf("F%v", floorNum)
// 		if s.Elevator == floorNum {
// 			fmt.Print(" E")
// 		} else {
// 			fmt.Print(" .")
// 		}

// 		for _, elem := range s.Elements {
// 			if floor.HasGenerator(elem) {
// 				fmt.Printf(" %v-G", string([]rune(elem)[0]))
// 			} else {
// 				fmt.Print("  . ")
// 			}

// 			if floor.HasChip(elem) {
// 				fmt.Printf(" %v-C", string([]rune(elem)[0]))
// 			} else {
// 				fmt.Print("  . ")
// 			}
// 		}

// 		fmt.Println()
// 	}
// 	fmt.Printf("\nMoves: %v\n\n", s.MoveCount)
// }

// func (s *State) GoalReached() bool {
// 	for _, floor := range s.Floors[:3] {
// 		if len(floor.Pieces) > 0 {
// 			return false
// 		}
// 	}

// 	return true
// }

// func (s *State) Equal(other State) bool {
// 	if s.Elevator != other.Elevator {
// 		return false
// 	}

// 	for i, sflr := range s.Floors {
// 		oflr := other.Floors[i]

// 		if !sflr.Equal(oflr) {
// 			return false
// 		}
// 	}

// 	return true
// }

// func (s *State) IsoState() [][2]int {
// 	isoState := make([][2]int, len(s.Elements))
// 	for i, elem := range s.Elements {
// 		isoState[i] = [2]int{-1, -1}
// 		for j, floor := range s.Floors {
// 			if floor.HasGenerator(elem) {
// 				isoState[i][0] = j
// 			}

// 			if floor.HasChip(elem) {
// 				isoState[i][1] = j
// 			}
// 		}
// 	}

// 	// fmt.Println(isoState)

// 	return isoState
// }

// func (s *State) Isomorphic(other State) bool {
// 	if s.Elevator != other.Elevator {
// 		return false
// 	}

// 	isoState1 := s.IsoState()
// 	isoState2 := other.IsoState()

// 	for _, s1 := range isoState1 {
// 		existsInOther := false
// 		for _, s2 := range isoState2 {
// 			if s1[0] == s2[0] && s1[1] == s2[1] {
// 				existsInOther = true
// 				break
// 			}
// 		}

// 		if !existsInOther {
// 			// fmt.Printf("%v !~= %v\n", isoState1, isoState2)
// 			return false
// 		}
// 	}

// 	// for _, s1 := range isoState2 {
// 	// 	existsInOther := false

// 	// 	for _, s2 := range isoState1 {
// 	// 		if s1[0] == s2[0] && s1[1] == s2[1] {
// 	// 			existsInOther = true
// 	// 			break
// 	// 		}
// 	// 	}

// 	// 	if !existsInOther {
// 	// 		fmt.Printf("%v !~= %v\n", isoState1, isoState2)
// 	// 		return false
// 	// 	}
// 	// }

// 	// fmt.Printf("%v ~= %v\n", isoState1, isoState2)
// 	return true
// }

// func (s *State) NumPieces() int {
// 	count := 0

// 	for _, floor := range s.Floors {
// 		count += len(floor.Pieces)
// 	}

// 	return count
// }

// func StateInHistory(history []State, ns State) bool {
// 	for _, oldS := range history {
// 		if oldS.Equal(ns) { // was Equal
// 			return true
// 		}
// 	}

// 	return false
// }

// func AdvanceState(oldState State, move Move) State {
// 	// oldState.Print()

// 	newFloors := []Floor{}
// 	// fmt.Println("\n")
// 	// oldState.Print()
// 	// fmt.Printf("Move Applied: %+v\n", move)

// 	for i, floor := range oldState.Floors {
// 		newPieces := []Piece{}
// 		// fmt.Print("Floor Index:", i, " ")
// 		// fmt.Printf("Floor: %+v\n", floor)
// 		switch i + 1 {
// 		case oldState.Elevator:
// 			fmt.Println("Cleaning out old floor")
// 			for _, piece := range floor.Pieces {
// 				if !move.HasPiece(piece) {
// 					newPieces = append(newPieces, piece)
// 				}
// 			}

// 			newFloor := Floor{
// 				Num:    floor.Num,
// 				Pieces: newPieces,
// 			}
// 			fmt.Printf("Result: %+v\n", newFloor)

// 			newFloors = append(newFloors, newFloor)
// 		case move.TargetFloor:
// 			fmt.Println("Creating new floor")

// 			newPieces = append(newPieces, floor.Pieces...)
// 			newPieces = append(newPieces, move.Pieces...)

// 			newFloor := Floor{
// 				Num:    floor.Num,
// 				Pieces: newPieces,
// 			}
// 			fmt.Printf("Result: %+v\n", newFloor)

// 			newFloors = append(newFloors, newFloor)
// 		default:
// 			fmt.Println("Floor unchanged")

// 			newFloor := Floor{
// 				Num:    floor.Num,
// 				Pieces: append(newPieces, floor.Pieces...),
// 			}
// 			fmt.Printf("Result: %+v\n", newFloor)

// 			newFloors = append(newFloors, newFloor)
// 		}
// 	}

// 	newState := State{
// 		Elevator:  move.TargetFloor,
// 		Floors:    newFloors,
// 		Elements:  oldState.Elements,
// 		MoveCount: oldState.MoveCount + 1,
// 	}

// 	if oldState.NumPieces() != newState.NumPieces() {
// 		fmt.Println("ERROR! Lost or duplicate Pieces!")
// 		oldState.Print()
// 		fmt.Printf("Move Applied: %+v\n", move)
// 		newState.Print()
// 		panic("Quitting")
// 	} else {
// 		newState.Print()
// 		fmt.Println("-----------------------------")
// 	}

// 	// newState.Print()

// 	return newState
// }

// type MoveStatePair struct {
// 	Move  Move
// 	State State
// }

// type MoveStateQueue struct {
// 	Items []*MoveStatePair
// }

// func (q *MoveStateQueue) push(move Move, state State) {
// 	msp := &MoveStatePair{move, state}
// 	q.Items = append(q.Items, msp)
// }

// func (q *MoveStateQueue) pop() *MoveStatePair {
// 	if len(q.Items) == 0 {
// 		return nil
// 	}

// 	item := q.Items[0]
// 	q.Items = q.Items[1:]

// 	return item
// }

// func BFS(startState State) {
// 	state := startState
// 	queue := MoveStateQueue{[]*MoveStatePair{}}
// 	statesInQueue := []State{state}

// 	newMoves := state.GenerateMoves(statesInQueue)
// 	fmt.Printf("Adding %v moves to the queue (current length %v)\n", len(newMoves), len(newMoves)+len(queue.Items))
// 	for _, move := range newMoves {
// 		fmt.Println(move)
// 		if !StateInHistory(statesInQueue, move.State) {
// 			queue.push(move.Move, state)
// 			statesInQueue = append(statesInQueue, move.State)
// 		}
// 	}

// 	lastMoveCount := 0

// 	item := queue.pop()
// 	var newState State
// 	for item != nil {
// 		// fmt.Println(item.Move)
// 		newState = AdvanceState(item.State, item.Move)
// 		newState.Print()

// 		if newState.MoveCount > lastMoveCount {
// 			lastMoveCount = newState.MoveCount
// 			fmt.Printf("Considering states %v moves out\n", lastMoveCount)
// 		}

// 		// if StateInHistory(statesInQueue, newState) {
// 		// 	// fmt.Println("Skipping state")
// 		// 	item = queue.pop()
// 		// 	continue
// 		// } else {
// 		// 	// statesInQueue = append(statesInQueue, newState)
// 		// }

// 		if newState.GoalReached() {
// 			newState.Print()
// 			break
// 		}

// 		newMoves := newState.GenerateMoves(statesInQueue)
// 		fmt.Printf("Adding %v moves to the queue (current length %v)\n", len(newMoves), len(newMoves)+len(queue.Items))
// 		for _, move := range newMoves {
// 			if !StateInHistory(statesInQueue, move.State) {
// 				queue.push(move.Move, newState)
// 				statesInQueue = append(statesInQueue, move.State)
// 			}
// 		}

// 		item = queue.pop()
// 	}
// }

// func ParseDataString(dat string) ([]Floor, []string) {
// 	floors := []Floor{}
// 	typesMap := map[string]bool{}

// 	generatorRegex := regexp.MustCompile("a ([^\\s]+) generator")
// 	chipRegex := regexp.MustCompile("a ([^\\s]+)-compatible microchip")

// 	for num, line := range strings.Split(dat, "\n") {
// 		pieces := []Piece{}
// 		generatorMatches := generatorRegex.FindAllStringSubmatch(line, -1)
// 		for _, match := range generatorMatches {
// 			pieces = append(pieces, Piece{
// 				Element: match[1],
// 				Type:    "generator",
// 			})
// 			typesMap[match[1]] = true
// 		}

// 		chipMatches := chipRegex.FindAllStringSubmatch(line, -1)
// 		for _, match := range chipMatches {
// 			pieces = append(pieces, Piece{
// 				Element: match[1],
// 				Type:    "chip",
// 			})
// 			typesMap[match[1]] = true
// 		}

// 		floors = append(floors, Floor{
// 			Num:    num + 1,
// 			Pieces: pieces,
// 		})
// 	}

// 	types := make([]string, len(typesMap))
// 	i := -1

// 	for key := range typesMap {
// 		i++
// 		types[i] = key
// 	}

// 	sort.Strings(types)

// 	return floors, types
// }

// func LoadData(filename string) ([]Floor, []string) {
// 	dat := lib.ReadFileData(filename)

// 	return ParseDataString(strings.Trim(string(dat), "\n"))
// }

// func RunPartA(filename string) {
// 	floors, allTypes := LoadData(filename)

// 	state := State{
// 		Elevator:  1,
// 		Floors:    floors,
// 		Elements:  allTypes,
// 		MoveCount: 0,
// 	}

// 	state.Print()

// 	BFS(state)
// }

// func RunPartB(filename string) {

// }
