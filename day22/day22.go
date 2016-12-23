package day22

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type FileSystem struct {
	Name string
	X    int
	Y    int
	Size int
	Used int
	Free int
}

func (fs *FileSystem) ToString() string {
	return fmt.Sprintf("%v/%v", fs.Used, fs.Size)
}

type SortData struct {
	Free  int
	Index int
}

func NewFileSystem(name string, size string, used string, free string) FileSystem {
	sizeInt, _ := strconv.Atoi(strings.Trim(size, "T"))
	usedInt, _ := strconv.Atoi(strings.Trim(used, "T"))
	freeInt, _ := strconv.Atoi(strings.Trim(free, "T"))

	nameRegex := regexp.MustCompile("-x(\\d+)-y(\\d+)")
	match := nameRegex.FindStringSubmatch(name)
	xInt, _ := strconv.Atoi(match[1])
	yInt, _ := strconv.Atoi(match[2])

	return FileSystem{
		Name: name,
		X:    xInt,
		Y:    yInt,
		Size: sizeInt,
		Used: usedInt,
		Free: freeInt,
	}
}

func (fs *FileSystem) Equals(other FileSystem) bool {
	return fs.Name == other.Name //unique identifier
	// return false
}

func (fs *FileSystem) CanMoveDataTo(other FileSystem, targetName string) bool {
	// if fs.Name == targetName && other.Used == 0 && other.Free >= fs.Used {
	// 	return true
	// }

	// return other.Free >= fs.Used
	return other.Free >= fs.Used
}

func (fs FileSystem) MoveDataTo(other FileSystem) (FileSystem, FileSystem) {
	other.Used += fs.Used
	other.Free -= fs.Used

	tmp := other.Name
	other.Name = fs.Name
	// other.Name += fs.Name
	fs.Name = tmp

	fs.Free = fs.Size
	fs.Used = 0

	return fs, other
}

func ParseDataString(dat string) []FileSystem {
	lines := strings.Split(dat, "\n")
	fs := make([]FileSystem, len(lines)-2)
	for i, line := range lines[2:] {
		parts := strings.Split(line, " ")
		fs[i] = NewFileSystem(parts[0], parts[1], parts[2], parts[3])
	}

	return fs
}

func LoadData(filename string) []FileSystem {
	// cat data/day22.txt | sort -k4n | tr -s ' ' | cut -f 1,3,4 -d ' '
	cat := exec.Command("cat", filename)
	sort := exec.Command("sort", "-k4n")
	tr := exec.Command("tr", "-s", " ")
	cut := exec.Command("cut", "-f", "1,2,3,4", "-d", " ")

	r1, w1 := io.Pipe()
	cat.Stdout = w1
	sort.Stdin = r1

	r2, w2 := io.Pipe()
	sort.Stdout = w2
	tr.Stdin = r2

	r3, w3 := io.Pipe()
	tr.Stdout = w3
	cut.Stdin = r3

	var outBuffer bytes.Buffer
	cut.Stdout = &outBuffer

	cat.Start()

	sort.Start()
	cat.Wait()
	w1.Close()

	tr.Start()
	sort.Wait()
	w2.Close()

	cut.Start()
	tr.Wait()
	w3.Close()

	cut.Wait()

	dat := outBuffer.String()

	return ParseDataString(strings.Trim(dat, "\n"))
}

func RunPartA(filename string) {
	fsList := LoadData(filename)
	helpers := []SortData{}

	lastValue := -1

	for i, fs := range fsList {
		if fs.Free > lastValue {
			helpers = append(helpers, SortData{
				Free:  fs.Free,
				Index: i,
			})
		}
	}

	pairCount := 0
	for _, fs := range fsList {
		if fs.Used == 0 {
			continue
		}

		for _, data := range helpers {
			if fs.Used <= data.Free {
				pairCount += len(fsList) - data.Index
				break
			}
		}
	}

	fmt.Println(pairCount)
}

type Grid struct {
	Rows     int
	Cols     int
	Data     [][]FileSystem
	NumMoves int
}

func NewGrid(rows, cols int, fsList []FileSystem) Grid {
	grid := Grid{
		Rows:     rows,
		Cols:     cols,
		Data:     make([][]FileSystem, rows),
		NumMoves: 0,
	}

	for i := 0; i < rows; i++ {
		grid.Data[i] = make([]FileSystem, cols)
	}

	for _, fs := range fsList {
		grid.Data[fs.Y][fs.X] = fs
	}

	return grid
}

func (g *Grid) Equals(other Grid) bool {
	//assume number of rows and columns are the same

	for r := 0; r < g.Rows; r++ {
		for c := 0; c < g.Cols; c++ {
			if !g.Data[r][c].Equals(other.Data[r][c]) {
				return false
			}
		}
	}

	return true
}

func (g *Grid) Copy() Grid {
	newGrid := NewGrid(g.Rows, g.Cols, []FileSystem{})
	newGrid.NumMoves = g.NumMoves

	for r := 0; r < g.Rows; r++ {
		for c := 0; c < g.Cols; c++ {
			newGrid.Data[r][c] = g.Data[r][c]
		}
	}

	return newGrid
}

func (g *Grid) IsComplete(targetName string) bool {
	return g.Data[0][0].Name == targetName
}

func (g *Grid) FindEmpty() (int, int) {
	for r := 0; r < g.Rows; r++ {
		for c := 0; c < g.Cols; c++ {
			if g.Data[r][c].Used == 0 {
				return r, c
			}
		}
	}

	return -1, -1
}

func (g *Grid) IsImmovable(r, c int, targetName string) bool {
	// return g.Data[r][c].Free < 13
	var adj FileSystem
	curr := g.Data[r][c]
	if r > 0 {
		// Check swapping up
		adj = g.Data[r-1][c]
		if curr.CanMoveDataTo(adj, targetName) {
			return false
		}
	}

	if r < g.Rows-1 {
		// Check swapping down
		adj = g.Data[r+1][c]
		if curr.CanMoveDataTo(adj, targetName) {
			return false
		}
	}

	if c > 0 {
		// Check swapping left
		adj = g.Data[r][c-1]
		if curr.CanMoveDataTo(adj, targetName) {
			return false
		}
	}

	if c < g.Cols-1 {
		// Check swapping right
		adj = g.Data[r][c+1]
		if curr.CanMoveDataTo(adj, targetName) {
			return false
		}
	}

	return true
}

func (g *Grid) Print(targetName string) {
	for r := 0; r < g.Rows; r++ {
		for c := 0; c < g.Cols; c++ {
			// fmt.Printf("%v/%v   ", g.Data[r][c].Used, g.Data[r][c].Size)
			switch {
			case g.Data[r][c].Name == targetName:
				fmt.Print(" G ")
			case g.Data[r][c].Used == 0:
				fmt.Print(" ? ")
			case g.IsImmovable(r, c, targetName):
				fmt.Print(" # ")
			default:
				fmt.Print(" . ")
			}
		}

		fmt.Println()
	}
}

func (g *Grid) GenerateMoves(history []Grid, targetName string) []Grid {
	newGrids := []Grid{}

	var curr FileSystem
	var adj FileSystem
	var newCurr FileSystem
	var newAdj FileSystem
	var newGrid Grid

	for r := 0; r < g.Rows; r++ {
		for c := 0; c < g.Rows; c++ {
			newGrid = g.Copy()
			curr = newGrid.Data[r][c]

			if r > 0 {
				// Check swapping up
				adj = newGrid.Data[r-1][c]
				if curr.CanMoveDataTo(adj, targetName) {
					newCurr, newAdj = curr.MoveDataTo(adj)
					newGrid.Data[r][c] = newCurr
					newGrid.Data[r-1][c] = newAdj
					newGrid.NumMoves++
					if !GridInHistory(history, newGrid) {
						newGrids = append(newGrids, newGrid)
					}
					newGrid = g.Copy()
				}
			}

			if r < g.Rows-1 {
				// Check swapping down
				adj = newGrid.Data[r+1][c]
				if curr.CanMoveDataTo(adj, targetName) {
					newCurr, newAdj = curr.MoveDataTo(adj)
					newGrid.Data[r][c] = newCurr
					newGrid.Data[r+1][c] = newAdj
					newGrid.NumMoves++
					if !GridInHistory(history, newGrid) {
						newGrids = append(newGrids, newGrid)
					}
					newGrid = g.Copy()
				}
			}

			if c > 0 {
				// Check swapping left
				adj = newGrid.Data[r][c-1]
				if curr.CanMoveDataTo(adj, targetName) {
					newCurr, newAdj = curr.MoveDataTo(adj)
					newGrid.Data[r][c] = newCurr
					newGrid.Data[r][c-1] = newAdj
					newGrid.NumMoves++
					if !GridInHistory(history, newGrid) {
						newGrids = append(newGrids, newGrid)
					}
					newGrid = g.Copy()
				}
			}

			if c < g.Cols-1 {
				// Check swapping right
				adj = newGrid.Data[r][c+1]
				if curr.CanMoveDataTo(adj, targetName) {
					newCurr, newAdj = curr.MoveDataTo(adj)
					newGrid.Data[r][c] = newCurr
					newGrid.Data[r][c+1] = newAdj
					newGrid.NumMoves++
					if !GridInHistory(history, newGrid) {
						newGrids = append(newGrids, newGrid)
					}
					newGrid = g.Copy()
				}
			}
		}
	}

	return newGrids
}

func GridInHistory(history []Grid, grid Grid) bool {
	for _, hGrid := range history {
		if hGrid.Equals(grid) {
			return true
		}
	}

	return false
}

type GridQueue struct {
	Items []*Grid
}

func (q *GridQueue) Push(grid Grid) {
	newGrid := grid.Copy()
	q.Items = append(q.Items, &newGrid)
}

func (q *GridQueue) Pop() *Grid {
	if len(q.Items) == 0 {
		return nil
	}

	item := q.Items[0]
	q.Items = q.Items[1:]

	return item
}

func BFS(startGrid Grid, targetName string) {
	queue := GridQueue{[]*Grid{&startGrid}}
	history := []Grid{startGrid}

	lastMoveCount := 0

	grid := queue.Pop()
	for grid != nil {
		// grid.Print()
		// history = append(history, *grid)

		if grid.NumMoves > lastMoveCount {
			lastMoveCount = grid.NumMoves
			fmt.Printf("Considering states %v moves out\n", lastMoveCount)
		}

		if grid.IsComplete(targetName) {
			// newState.Print()
			fmt.Println(grid.NumMoves)
			break
		}

		newMoves := grid.GenerateMoves(history, targetName)
		fmt.Printf("Adding %v moves to the queue (current length %v)\n", len(newMoves), len(newMoves)+len(queue.Items))
		for _, move := range newMoves {
			queue.Push(move)
			history = append(history, move)
		}

		grid = queue.Pop()
	}
}

func RunPartB(filename string) {
	fsList := LoadData(filename)

	maxX := -1
	maxY := -1
	for _, fs := range fsList {
		if fs.X > maxX {
			maxX = fs.X
		}

		if fs.Y > maxY {
			maxY = fs.Y
		}
	}
	targetName := fmt.Sprintf("/dev/grid/node-x%v-y%v", maxX, 0)
	grid := NewGrid(maxY+1, maxX+1, fsList)
	reader := bufio.NewReader(os.Stdin)
	row, col := grid.FindEmpty()
	var curr FileSystem
	var adj FileSystem
	for {
		curr = grid.Data[row][col]
		fmt.Printf("Currently at %v, %v (%v/%v)\n", row, col, curr.Used, curr.Size)
		grid.Print(targetName)

		fmt.Print("Enter move: ")
		moveRaw, _ := reader.ReadString('\n')

		move := []rune(moveRaw)[0]

		switch move {
		case 'l':
			fmt.Println("Moving left.")
			adj = grid.Data[row][col-1]
		case 'r':
			fmt.Println("Moving right.")
			adj = grid.Data[row][col+1]
		case 'u':
			fmt.Println("Moving up.")
			adj = grid.Data[row-1][col]
		case 'd':
			fmt.Println("Moving down.")
			adj = grid.Data[row+1][col]
		}

		if adj.CanMoveDataTo(curr, targetName) {
			fmt.Printf("Moving data from %v to %v\n", adj.ToString(), curr.ToString())
			newAdj, newCurr := adj.MoveDataTo(curr)
			grid.Data[row][col] = newCurr
			switch move {
			case 'l':
				grid.Data[row][col-1] = newAdj
				col--
			case 'r':
				grid.Data[row][col+1] = newAdj
				col++
			case 'u':
				grid.Data[row-1][col] = newAdj
				row--
			case 'd':
				grid.Data[row+1][col] = newAdj
				row++
			}
		} else {
			fmt.Println("Can't move data there.")
		}

	}
	// BFS(grid, targetName)
}
