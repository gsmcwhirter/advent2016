package day11

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gsmcwhirter/advent2016/lib"
)

type PieceType int

const (
	GENERATOR PieceType = iota
	CHIP
)

type Element int

const (
	ELEMENT_A Element = iota
	ELEMENT_B
	ELEMENT_C
	ELEMENT_D
	ELEMENT_E
)

type Piece struct {
	Element Element
	Type    PieceType
}

func (p *Piece) Equals(other Piece) bool {
	return p.Element == other.Element && p.Type == other.Type
}

func (p *Piece) ToString() string {
	elemString := " "
	typeString := " "

	switch p.Element {
	case ELEMENT_A:
		elemString = "A"
	case ELEMENT_B:
		elemString = "B"
	case ELEMENT_C:
		elemString = "C"
	case ELEMENT_D:
		elemString = "D"
	}

	switch p.Type {
	case GENERATOR:
		typeString = "G"
	case CHIP:
		typeString = "C"
	}

	return fmt.Sprintf("%s-%s", elemString, typeString)
}

func (p *Piece) PairPiece() Piece {
	if p.Type == CHIP {
		return Piece{
			p.Element,
			GENERATOR,
		}
	}

	return Piece{
		p.Element,
		CHIP,
	}
}

type PieceSet struct {
	PiecesByElement map[Element]map[PieceType]Piece
	PiecesByType    map[PieceType]map[Element]Piece
	Size            int
}

func NewPieceSet() PieceSet {
	ps := PieceSet{
		PiecesByElement: map[Element]map[PieceType]Piece{},
		PiecesByType:    map[PieceType]map[Element]Piece{},
		Size:            0,
	}

	ps.PiecesByType[CHIP] = map[Element]Piece{}
	ps.PiecesByType[GENERATOR] = map[Element]Piece{}

	ps.PiecesByElement[ELEMENT_A] = map[PieceType]Piece{}
	ps.PiecesByElement[ELEMENT_B] = map[PieceType]Piece{}
	ps.PiecesByElement[ELEMENT_C] = map[PieceType]Piece{}
	ps.PiecesByElement[ELEMENT_D] = map[PieceType]Piece{}
	ps.PiecesByElement[ELEMENT_E] = map[PieceType]Piece{}

	return ps
}

func (ps *PieceSet) Add(piece Piece) {
	_, exists := ps.PiecesByElement[piece.Element][piece.Type]
	if !exists {
		ps.PiecesByElement[piece.Element][piece.Type] = piece
		ps.Size++
	}

	ps.PiecesByType[piece.Type][piece.Element] = piece
}

func (ps *PieceSet) Contains(piece Piece) bool {
	_, exists := ps.PiecesByElement[piece.Element][piece.Type]
	if exists {
		return true
	}

	return false
}

func (ps *PieceSet) Remove(piece Piece) {
	if ps.Contains(piece) {
		delete(ps.PiecesByElement[piece.Element], piece.Type)
		delete(ps.PiecesByType[piece.Type], piece.Element)
		ps.Size--
	}
}

func (ps *PieceSet) PieceList() []Piece {
	plist := make([]Piece, ps.Size)
	i := 0

	for element := range ps.PiecesByElement {
		for ptype := range ps.PiecesByElement[element] {
			plist[i] = ps.PiecesByElement[element][ptype]
			i++
		}
	}

	return plist
}

func (ps *PieceSet) Equals(other PieceSet) bool {
	if ps.Size != other.Size {
		return false
	}

	for _, piece := range ps.PieceList() {
		if !other.Contains(piece) {
			return false
		}
	}

	return true
}

func (ps *PieceSet) Copy() PieceSet {
	newPS := NewPieceSet()
	newPS.Size = ps.Size
	for element := range ps.PiecesByElement {
		newPS.PiecesByElement[element] = map[PieceType]Piece{}

		for ptype := range ps.PiecesByElement[element] {
			newPS.PiecesByElement[element][ptype] = ps.PiecesByElement[element][ptype]
		}
	}

	for ptype := range ps.PiecesByType {
		newPS.PiecesByType[ptype] = map[Element]Piece{}

		for element := range ps.PiecesByType[ptype] {
			newPS.PiecesByType[ptype][element] = ps.PiecesByType[ptype][element]
		}
	}

	return newPS
}

func (ps *PieceSet) IsLegal() bool {
	for _, piece := range ps.PiecesByType[CHIP] {
		chipOkay := false
		if ps.Contains(piece.PairPiece()) {
			chipOkay = true
		} else {
			if len(ps.PiecesByType[GENERATOR]) == 0 {
				chipOkay = true
			}
		}

		if !chipOkay {
			return false
		}
	}

	return true
}

func (ps *PieceSet) ToString() string {
	str := ""
	for element := range ps.PiecesByElement {
		for _, piece := range ps.PiecesByElement[element] {
			str = fmt.Sprintf("%v, %v", str, piece.ToString())
		}
	}

	return str
}

type Move struct {
	CurrentFloor int
	TargetFloor  int
	Pieces       PieceSet
}

func (m *Move) Equals(other Move) bool {
	return m.CurrentFloor == other.CurrentFloor && m.TargetFloor == other.TargetFloor && m.Pieces.Equals(other.Pieces)
}

func (m *Move) ToString() string {
	return fmt.Sprintf("%v -> %v%v", m.CurrentFloor, m.TargetFloor, m.Pieces.ToString())
}

type PairLocation struct {
	GeneratorLoc int
	ChipLoc      int
}

func (pl *PairLocation) Equals(other PairLocation) bool {
	return pl.GeneratorLoc == other.GeneratorLoc && pl.ChipLoc == other.ChipLoc
}

type State struct {
	Floors        int
	Elevator      int
	Elements      []Element
	MovesExecuted []Move
	Locations     map[Element]PairLocation
}

func NewState(floors int) State {
	return State{
		floors,
		0,
		[]Element{
			ELEMENT_A,
			ELEMENT_B,
			ELEMENT_C,
			ELEMENT_D,
			ELEMENT_E,
		},
		[]Move{},
		map[Element]PairLocation{},
	}
}

func (s *State) PairLocationList() []PairLocation {
	plList := make([]PairLocation, len(s.Locations))
	i := 0

	for element := range s.Locations {
		plList[i] = s.Locations[element]
		i++
	}

	return plList
}

func (s *State) Isomorphic(other State) bool {
	if s.Elevator != other.Elevator {
		return false
	}

	// assumes len(PairLocationList) are identical

	list1 := s.PairLocationList()
	list2 := other.PairLocationList()
	l1ItemsFound := map[int]bool{}
	l2ItemsFound := map[int]bool{}

	for _, pl1 := range list1 {
		for i, pl2 := range list2 {
			if pl1.Equals(pl2) && !l2ItemsFound[i] {
				l2ItemsFound[i] = true
				break
			}
		}
	}

	for i := 0; i < len(list2); i++ {
		if !l2ItemsFound[i] {
			return false
		}
	}

	for _, pl2 := range list2 {
		for i, pl1 := range list1 {
			if pl2.Equals(pl1) && !l1ItemsFound[i] {
				l1ItemsFound[i] = true
				break
			}
		}
	}

	for i := 0; i < len(list1); i++ {
		if !l1ItemsFound[i] {
			return false
		}
	}

	return true
}

func (s *State) PiecesOnFloor(floor int) []Piece {
	plist := []Piece{}

	for element, pl := range s.Locations {
		if pl.GeneratorLoc == floor {
			plist = append(plist, Piece{element, GENERATOR})
		}

		if pl.ChipLoc == floor {
			plist = append(plist, Piece{element, CHIP})
		}
	}

	return plist
}

func (s *State) IsFloorEmpty(floor int) bool {
	list := s.PairLocationList()
	for _, pl := range list {
		if pl.ChipLoc == floor || pl.GeneratorLoc == floor {
			return false
		}
	}

	return true
}

func (s *State) Copy() State {
	newS := NewState(s.Floors)
	newS.Elevator = s.Elevator
	newS.MovesExecuted = append([]Move{}, s.MovesExecuted...)

	for element := range s.Locations {
		newS.Locations[element] = s.Locations[element]
	}

	return newS
}

func (s *State) IsLegal() bool {
	for floor := 0; floor < s.Floors; floor++ {
		floorPieces := NewPieceSet()
		for element, pl := range s.Locations {
			if pl.ChipLoc == floor {
				floorPieces.Add(Piece{element, CHIP})
			}

			if pl.GeneratorLoc == floor {
				floorPieces.Add(Piece{element, GENERATOR})
			}
		}

		if !floorPieces.IsLegal() {
			return false
		}
	}

	return true
}

func (s *State) MinFloor() int {
	minFloor := 0
	for f := 0; f < s.Elevator; f++ {
		if s.IsFloorEmpty(f) {
			minFloor++
		} else {
			break
		}
	}

	return minFloor
}

func (s *State) IsComplete() bool {
	return s.Elevator == s.Floors-1 && s.MinFloor() == s.Elevator
}

func (s *State) GenerateMoves(history []State) []Move {
	movesUp := []Move{}
	movesDown := []Move{}

	minFloor := s.MinFloor()

	upFloor := s.Elevator + 1
	downFloor := s.Elevator - 1

	piecesOnCurrentFloor := s.PiecesOnFloor(s.Elevator)

	var move Move
	var ps PieceSet

	for i, piece1 := range piecesOnCurrentFloor {
		for j, piece2 := range piecesOnCurrentFloor {
			if j <= i {
				continue
			}

			if upFloor < s.Floors {
				ps = NewPieceSet()
				ps.Add(piece1)
				ps.Add(piece2)

				move = Move{
					CurrentFloor: s.Elevator,
					TargetFloor:  upFloor,
					Pieces:       ps,
				}

				// fmt.Printf("Considering move %v\n", move.ToString())
				if s.MoveIsLegal(move, history) {
					movesUp = append(movesUp, move)
				}
			}
		}

		if downFloor >= minFloor {
			ps = NewPieceSet()
			ps.Add(piece1)

			move = Move{
				CurrentFloor: s.Elevator,
				TargetFloor:  downFloor,
				Pieces:       ps,
			}

			// fmt.Printf("Considering move %v\n", move.ToString())
			if s.MoveIsLegal(move, history) {
				movesDown = append(movesDown, move)
			}
		}
	}

	onePieceUp := len(movesUp) == 0
	twoPieceDown := len(movesDown) == 0

	if onePieceUp || twoPieceDown {
		for i, piece1 := range piecesOnCurrentFloor {
			if twoPieceDown {
				for j, piece2 := range piecesOnCurrentFloor {
					if j <= i {
						continue
					}

					if downFloor >= minFloor {
						ps = NewPieceSet()
						ps.Add(piece1)
						ps.Add(piece2)

						move = Move{
							CurrentFloor: s.Elevator,
							TargetFloor:  downFloor,
							Pieces:       ps,
						}

						// fmt.Printf("Considering move %v\n", move.ToString())
						if s.MoveIsLegal(move, history) {
							movesUp = append(movesUp, move)
						}
					}
				}
			}

			if onePieceUp {
				if upFloor < s.Floors {
					ps = NewPieceSet()
					ps.Add(piece1)

					move = Move{
						CurrentFloor: s.Elevator,
						TargetFloor:  upFloor,
						Pieces:       ps,
					}

					// fmt.Printf("Considering move %v\n", move.ToString())
					if s.MoveIsLegal(move, history) {
						movesDown = append(movesDown, move)
					}
				}
			}
		}
	}

	return append(movesUp, movesDown...)
}

func (s *State) MoveIsLegal(move Move, history []State) bool {
	newS := s.Copy()
	newS.ApplyMove(move)

	return newS.IsLegal() && !StateInHistory(history, newS)
}

func (s *State) ApplyMove(move Move) {
	// fmt.Printf("Moving elevator from %v to %v\n", s.Elevator, move.TargetFloor)
	s.Elevator = move.TargetFloor
	s.MovesExecuted = append(s.MovesExecuted, move)

	for _, piece := range move.Pieces.PieceList() {
		if piece.Type == CHIP {
			s.Locations[piece.Element] = PairLocation{
				ChipLoc:      move.TargetFloor,
				GeneratorLoc: s.Locations[piece.Element].GeneratorLoc,
			}
			continue
		}

		if piece.Type == GENERATOR {
			s.Locations[piece.Element] = PairLocation{
				ChipLoc:      s.Locations[piece.Element].ChipLoc,
				GeneratorLoc: move.TargetFloor,
			}
		}
	}
}

func (s *State) NumMoves() int {
	return len(s.MovesExecuted)
}

func (s *State) Print() {
	var piece Piece
	for floorNum := s.Floors - 1; floorNum >= 0; floorNum-- {

		fmt.Printf("F%v", floorNum+1)
		if s.Elevator == floorNum {
			fmt.Print(" E")
		} else {
			fmt.Print(" .")
		}

		for _, element := range s.Elements {
			if s.Locations[element].GeneratorLoc == floorNum {
				piece = Piece{element, GENERATOR}
				fmt.Printf(" %v", piece.ToString())
			} else {
				fmt.Print("  . ")
			}

			if s.Locations[element].ChipLoc == floorNum {
				piece = Piece{element, CHIP}
				fmt.Printf(" %v", piece.ToString())
			} else {
				fmt.Print("  . ")
			}
		}

		fmt.Println()
	}
	fmt.Printf("\nMoves: %v\n\n", len(s.MovesExecuted))
}

func StateInHistory(history []State, state State) bool {
	for _, oldState := range history {
		if state.Isomorphic(oldState) {
			return true
		}
	}

	return false
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

func BFS(startState State) {
	queue := StateQueue{[]*State{}}
	history := []State{startState}

	newMoves := startState.GenerateMoves(history)
	var nextState State

	fmt.Printf("Adding %v moves to the queue (current length %v)\n", len(newMoves), len(newMoves)+len(queue.Items))
	for _, move := range newMoves {
		fmt.Println(move)
		nextState = startState.Copy()
		nextState.ApplyMove(move)

		queue.Push(nextState)
		// history = append(history, nextState)
	}

	lastMoveCount := 0

	newState := queue.Pop()
	for newState != nil {
		// newState.Print()

		history = append(history, nextState)

		if newState.NumMoves() > lastMoveCount {
			lastMoveCount = newState.NumMoves()
			fmt.Printf("Considering states %v moves out\n", lastMoveCount)
		}

		if newState.IsComplete() {
			newState.Print()
			break
		}

		newMoves := newState.GenerateMoves(history)
		fmt.Printf("Adding %v moves to the queue (current length %v)\n", len(newMoves), len(newMoves)+len(queue.Items))
		for _, move := range newMoves {
			nextState = newState.Copy()
			nextState.ApplyMove(move)

			// if !StateInHistory(history, nextState) {
			queue.Push(nextState)
			// history = append(history, nextState)
			// }
		}

		newState = queue.Pop()
	}
}

func ParseDataString(dat string) State {

	state := NewState(4)

	elementsMap := map[string]int{}
	elements := []Element{ELEMENT_A, ELEMENT_B, ELEMENT_C, ELEMENT_D, ELEMENT_E}
	nextElement := 0

	generatorRegex := regexp.MustCompile("a ([^\\s]+) generator")
	chipRegex := regexp.MustCompile("a ([^\\s]+)-compatible microchip")

	for num, line := range strings.Split(dat, "\n") {
		generatorMatches := generatorRegex.FindAllStringSubmatch(line, -1)
		for _, match := range generatorMatches {

			elementIdx, exists := elementsMap[match[1]]
			if !exists {
				if nextElement > 3 {
					fmt.Printf("Error: %+v, %v\n", elementsMap, match[1])
				}
				elementIdx = nextElement
				elementsMap[match[1]] = elementIdx
				nextElement++
			}
			element := elements[elementIdx]

			_, plExists := state.Locations[element]
			if !plExists {
				state.Locations[element] = PairLocation{
					ChipLoc:      -1,
					GeneratorLoc: num,
				}
			} else {
				state.Locations[element] = PairLocation{
					ChipLoc:      state.Locations[element].ChipLoc,
					GeneratorLoc: num,
				}
			}
		}

		chipMatches := chipRegex.FindAllStringSubmatch(line, -1)
		for _, match := range chipMatches {

			elementIdx, exists := elementsMap[match[1]]
			if !exists {
				if nextElement > 3 {
					fmt.Printf("Error: %+v, %v\n", elementsMap, match[1])
				}
				elementIdx = nextElement
				elementsMap[match[1]] = elementIdx
				nextElement++
			}
			element := elements[elementIdx]

			_, plExists := state.Locations[element]
			if !plExists {
				state.Locations[element] = PairLocation{
					GeneratorLoc: -1,
					ChipLoc:      num,
				}
			} else {
				state.Locations[element] = PairLocation{
					GeneratorLoc: state.Locations[element].GeneratorLoc,
					ChipLoc:      num,
				}
			}
		}
	}

	return state
}

func LoadData(filename string) State {
	dat := lib.ReadFileData(filename)

	return ParseDataString(strings.Trim(string(dat), "\n"))
}

func RunPartA(filename string) {
	startState := LoadData(filename)

	// startState.Print()

	BFS(startState)
}

func RunPartB(filename string) {

}
