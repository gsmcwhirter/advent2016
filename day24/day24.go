package day24

import (
	"fmt"
	"sort"
	"strings"

	"github.com/gsmcwhirter/advent2016/lib"
)

// type GraphNode struct {
// 	X         int
// 	Y         int
// 	Neighbors []*GraphNode
// 	Visited   bool
// 	Name      rune
// }

// func NewGraphNode(x, y int, name rune) *GraphNode {
// 	gn := &GraphNode{
// 		X:         x,
// 		Y:         y,
// 		Neighbors: []*GraphNode{},
// 		Visited:   false,
// 		Name:      name,
// 	}

// 	return gn
// }

// type SpecialNodes struct {
// 	Names map[rune]*GraphNode
// }

type Keys []rune
type ShortestDistPairs map[rune]map[rune]int
type Location []int

type GridSquare struct {
	IsWall      bool
	IsSpecial   bool
	SpecialName rune
}

type Grid struct {
	Rows    int
	Cols    int
	Data    [][]GridSquare
	CurrRow int
	CurrCol int
}

func (g *Grid) Print() {
	for r := 0; r < g.Rows; r++ {
		for c := 0; c < g.Cols; c++ {
			fmt.Print(string(g.Data[r][c].SpecialName))
		}
		fmt.Println()
	}
}

func (g *Grid) GenerateMoves(from Location, visited [][]bool) []Location {
	newMoves := make([]Location, 0, 4)

	if !g.Data[from[0]-1][from[1]].IsWall && !visited[from[0]-1][from[1]] {
		newMoves = append(newMoves, Location([]int{from[0] - 1, from[1], from[2] + 1}))
	}

	if !g.Data[from[0]+1][from[1]].IsWall && !visited[from[0]+1][from[1]] {
		newMoves = append(newMoves, Location([]int{from[0] + 1, from[1], from[2] + 1}))
	}

	if !g.Data[from[0]][from[1]-1].IsWall && !visited[from[0]][from[1]-1] {
		newMoves = append(newMoves, Location([]int{from[0], from[1] - 1, from[2] + 1}))
	}

	if !g.Data[from[0]][from[1]+1].IsWall && !visited[from[0]][from[1]+1] {
		newMoves = append(newMoves, Location([]int{from[0], from[1] + 1, from[2] + 1}))
	}

	return newMoves
}

func (g *Grid) FindShortestDistance(loc1, loc2 Location) int {
	queue := LocationQueue{[]*Location{}}
	visited := make([][]bool, g.Rows)
	for r := 0; r < g.Rows; r++ {
		visited[r] = make([]bool, g.Cols)
	}

	start := Location(make([]int, 3))
	start[0] = loc1[0]
	start[1] = loc1[1]
	start[2] = 0

	queue.Push(start)
	visited[loc1[0]][loc1[1]] = true

	loc := queue.Pop()
	for loc != nil {
		if (*loc)[0] == loc2[0] && (*loc)[1] == loc2[1] {
			return (*loc)[2]
		}

		newMoves := g.GenerateMoves(*loc, visited)
		for _, move := range newMoves {
			queue.Push(move)
			visited[move[0]][move[1]] = true
		}

		loc = queue.Pop()
	}

	return -1
}

func (g *Grid) GenerateShortestDistancePairs(specialLocations map[rune][]int) (Keys, ShortestDistPairs) {
	shortestDistPairs := ShortestDistPairs(map[rune]map[rune]int{})

	keys := Keys(make([]rune, len(specialLocations)))
	i := 0
	for key := range specialLocations {
		keys[i] = key
		shortestDistPairs[key] = map[rune]int{}
		i++
	}

	for ki, key1 := range keys {
		for _, key2 := range keys[ki+1:] {
			dist := g.FindShortestDistance(specialLocations[key1], specialLocations[key2])
			shortestDistPairs[key1][key2] = dist
			shortestDistPairs[key2][key1] = dist
		}
	}

	return keys, shortestDistPairs
}

type LocationQueue struct {
	Items []*Location
}

func (q *LocationQueue) Push(loc Location) {
	newState := Location(append([]int{}, loc...))
	q.Items = append(q.Items, &newState)
}

func (q *LocationQueue) Pop() *Location {
	if len(q.Items) == 0 {
		return nil
	}

	item := q.Items[0]
	q.Items = q.Items[1:]

	return item
}

type KeyHistory struct {
	Key          rune
	Visited      map[rune]bool
	VisitedOrder []rune
	TotalSteps   int
}

func (kh *KeyHistory) Copy() KeyHistory {
	newKH := KeyHistory{
		Key:          kh.Key,
		Visited:      map[rune]bool{},
		VisitedOrder: append([]rune{}, kh.VisitedOrder...),
		TotalSteps:   kh.TotalSteps,
	}

	for key, val := range kh.Visited {
		newKH.Visited[key] = val
	}

	return newKH
}

type SortableKeyQueue struct {
	Items []*KeyHistory
}

func (q *SortableKeyQueue) Push(loc KeyHistory) {
	newKH := loc.Copy()
	q.Items = append(q.Items, &newKH)
}

func (q *SortableKeyQueue) Pop() *KeyHistory {
	if len(q.Items) == 0 {
		return nil
	}

	item := q.Items[0]
	q.Items = q.Items[1:]

	return item
}

func (q SortableKeyQueue) Len() int {
	return len(q.Items)
}

func (q SortableKeyQueue) Less(i, j int) bool {
	if q.Items[i].TotalSteps < q.Items[j].TotalSteps {
		return true
	}

	return false
}

func (q SortableKeyQueue) Swap(i, j int) {
	tmp := q.Items[i]
	q.Items[i] = q.Items[j]
	q.Items[j] = tmp
}

func FindShortestRoute(start rune, keys Keys, shortestDistPairs ShortestDistPairs) *KeyHistory {
	state := KeyHistory{
		Key:          start,
		Visited:      map[rune]bool{start: true},
		VisitedOrder: []rune{start},
		TotalSteps:   0,
	}

	queue := SortableKeyQueue{[]*KeyHistory{}}
	queue.Push(state)

	var done bool
	var nextState KeyHistory
	curr := queue.Pop()

	for curr != nil {
		fmt.Printf("Visiting %v (path %v)\n", string(curr.Key), string(curr.VisitedOrder))
		done = true
		for _, key := range keys {
			if !curr.Visited[key] {
				done = false
				nextState = curr.Copy()

				nextState.Key = key
				nextState.Visited[key] = true
				nextState.TotalSteps += shortestDistPairs[curr.Key][key]
				nextState.VisitedOrder = append(nextState.VisitedOrder, key)

				queue.Push(nextState)
			}
		}

		if done {
			return curr
		}

		sort.Sort(queue)

		curr = queue.Pop()
	}

	return nil
}

func FindShortestRouteEndAtTarget(start rune, target rune, keys Keys, shortestDistPairs ShortestDistPairs) *KeyHistory {
	state := KeyHistory{
		Key:          start,
		Visited:      map[rune]bool{start: true},
		VisitedOrder: []rune{start},
		TotalSteps:   0,
	}

	queue := SortableKeyQueue{[]*KeyHistory{}}
	queue.Push(state)

	var done bool
	var nextState KeyHistory
	curr := queue.Pop()

	for curr != nil {
		fmt.Printf("Visiting %v (path %v)\n", string(curr.Key), string(curr.VisitedOrder))
		done = true
		for _, key := range keys {
			if !curr.Visited[key] {
				done = false
				nextState = curr.Copy()

				nextState.Key = key
				nextState.Visited[key] = true
				nextState.TotalSteps += shortestDistPairs[curr.Key][key]
				nextState.VisitedOrder = append(nextState.VisitedOrder, key)

				queue.Push(nextState)
			}
		}

		if done && curr.Key == target {
			return curr
		}

		if done {
			curr.Visited[target] = false
			queue.Push(*curr)
		}

		sort.Sort(queue)

		curr = queue.Pop()
	}

	return nil
}

func ParseDataString(dat string) (Grid, map[rune][]int) {
	lines := strings.Split(dat, "\n")

	rows := len(lines)
	cols := len(lines[0])

	specialSpots := map[rune][]int{}

	grid := Grid{
		Rows:    rows,
		Cols:    cols,
		Data:    make([][]GridSquare, rows),
		CurrRow: -1,
		CurrCol: -1,
	}

	for i, line := range lines {
		grid.Data[i] = make([]GridSquare, cols)
		for j, spot := range []rune(line) {
			switch spot {
			case '#':
				grid.Data[i][j] = GridSquare{
					IsWall:      true,
					IsSpecial:   false,
					SpecialName: spot,
				}
			case '.':
				grid.Data[i][j] = GridSquare{
					IsWall:      false,
					IsSpecial:   false,
					SpecialName: spot,
				}
			default:
				grid.Data[i][j] = GridSquare{
					IsWall:      false,
					IsSpecial:   true,
					SpecialName: spot,
				}

				specialSpots[spot] = []int{i, j}
			}
		}
	}

	return grid, specialSpots
}

func LoadData(filename string) (Grid, map[rune][]int) {
	dat := lib.ReadFileData(filename)

	return ParseDataString(strings.Trim(string(dat), "\n"))
}

func RunPartA(filename string) {
	grid, specialLocations := LoadData(filename) //

	grid.Print()
	for name, loc := range specialLocations {
		fmt.Printf("%v @ %v, %v\n", string(name), loc[0], loc[1])
	}

	fmt.Println()
	keys, shortestDistPairs := grid.GenerateShortestDistancePairs(specialLocations)

	for ki, key1 := range keys {
		for _, key2 := range keys[ki+1:] {
			fmt.Printf("%v -- %v = %v\n", string(key1), string(key2), shortestDistPairs[key1][key2])
		}
	}

	shortest := FindShortestRoute('0', keys, shortestDistPairs)

	if shortest != nil {
		fmt.Printf("done in %v steps following %v!\n", shortest.TotalSteps, string(shortest.VisitedOrder))
	} else {
		fmt.Println("No path found???")
	}
}

func RunPartB(filename string) {
	grid, specialLocations := LoadData(filename) //

	grid.Print()
	for name, loc := range specialLocations {
		fmt.Printf("%v @ %v, %v\n", string(name), loc[0], loc[1])
	}

	fmt.Println()
	keys, shortestDistPairs := grid.GenerateShortestDistancePairs(specialLocations)

	for ki, key1 := range keys {
		for _, key2 := range keys[ki+1:] {
			fmt.Printf("%v -- %v = %v\n", string(key1), string(key2), shortestDistPairs[key1][key2])
		}
	}

	shortest := FindShortestRouteEndAtTarget('0', '0', keys, shortestDistPairs)

	if shortest != nil {
		fmt.Printf("done in %v steps following %v!\n", shortest.TotalSteps, string(shortest.VisitedOrder))
	} else {
		fmt.Println("No path found???")
	}
}
