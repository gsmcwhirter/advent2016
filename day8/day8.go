package day8

import (
	"fmt"
	"strings"

	"strconv"

	"github.com/gsmcwhirter/advent2016/lib"
)

type GridElement struct {
	Id    int
	On    bool
	Up    *GridElement
	Down  *GridElement
	Left  *GridElement
	Right *GridElement
}

func (g *GridElement) GoRight(num int) *GridElement {
	curr := g
	for i := 0; i < num; i++ {
		curr = curr.Right
	}

	return curr
}

func (g *GridElement) GoLeft(num int) *GridElement {
	curr := g
	for i := 0; i < num; i++ {
		curr = curr.Left
	}

	return curr
}

func (g *GridElement) GoUp(num int) *GridElement {
	curr := g
	for i := 0; i < num; i++ {
		curr = curr.Up
	}

	return curr
}

func (g *GridElement) GoDown(num int) *GridElement {
	curr := g
	for i := 0; i < num; i++ {
		curr = curr.Down
	}

	return curr
}

type Torus struct {
	Rows   int
	Cols   int
	Origin *GridElement
}

func NewTorus(rows, cols int) Torus {
	nextId := 0
	origin := &GridElement{Id: 0}
	nextId++

	var prevRowHead *GridElement
	var rowHead *GridElement
	var last *GridElement
	for r := 0; r < rows; r++ {
		//Column "0"

		if r == 0 {
			rowHead = origin
		} else {
			prevRowHead = rowHead
			rowHead = &GridElement{Id: nextId, Up: prevRowHead}
			nextId++

			prevRowHead.Down = rowHead

		}
		// fmt.Print("prev row")
		// fmt.Println(prevRowHead)
		// fmt.Print("New row ")
		// fmt.Println(rowHead)

		last = rowHead

		//Rest of the row columns
		for c := 1; c < cols; c++ {
			last.Right = &GridElement{Id: nextId, Left: last}
			nextId++

			last = last.Right

			if r > 0 {
				above := rowHead.Up.GoRight(c)
				above.Down = last
				last.Up = above
			}
		}

		// Connect around the row
		last.Right = rowHead
		rowHead.Left = last

		// fmt.Print("done row")
		// fmt.Println(rowHead)
	}

	fmt.Println("rows done")

	for c := 0; c < cols; c++ {
		origin.GoRight(c).Up = rowHead.GoRight(c)
		rowHead.GoRight(c).Down = origin.GoRight(c)
	}

	fmt.Println("wraparound done")

	// for r := 0; r < rows; r++ {
	// 	for c := 0; c < cols; c++ {
	// 		x := origin.GoDown(r).GoRight(c)
	// 		fmt.Printf("%v -- u: %v, d: %v, l: %v, r: %v\n", x.Id, x.Up.Id, x.Down.Id, x.Left.Id, x.Right.Id)
	// 	}
	// }

	return Torus{
		Rows:   rows,
		Cols:   cols,
		Origin: origin,
	}
}

func (t *Torus) TurnOn(row, col int) {
	t.Origin.GoDown(row).GoRight(col).On = true
}

func (t *Torus) TurnOff(row, col int) {
	t.Origin.GoDown(row).GoRight(col).On = false
}

func (t *Torus) Rect(width, height int) {
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			t.TurnOn(r, c)
		}
	}
}

func (t *Torus) Print() {
	for r := 0; r < t.Rows; r++ {
		for c := 0; c < t.Cols; c++ {
			x := t.Origin.GoDown(r).GoRight(c)
			if x.On {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (t *Torus) RotateRow(row, shift int) {
	rowHead := t.Origin.GoDown(row)
	above := rowHead.Up
	below := rowHead.Down

	for c := 0; c < t.Cols; c++ {
		newAbove := above.GoRight(c + shift)
		newBelow := below.GoRight(c + shift)
		curr := rowHead.GoRight(c)

		curr.Up = newAbove
		newAbove.Down = curr

		curr.Down = newBelow
		newBelow.Up = curr
	}

	if row == 0 {
		t.Origin = rowHead.GoLeft(shift)
	}
}

func (t *Torus) RotateCol(col, shift int) {
	colHead := t.Origin.GoRight(col)
	left := colHead.Left
	right := colHead.Right

	for r := 0; r < t.Rows; r++ {
		newLeft := left.GoDown(r + shift)
		newRight := right.GoDown(r + shift)
		curr := colHead.GoDown(r)

		curr.Left = newLeft
		newLeft.Right = curr

		curr.Right = newRight
		newRight.Left = curr
	}

	if col == 0 {
		t.Origin = colHead.GoUp(shift)
	}
}

func (t *Torus) ApplyCommand(cmd Command) {
	switch cmd.Command {
	case "rect":
		t.Rect(cmd.Value1, cmd.Value2)
	case "rotate":
		if cmd.Command2 == "row" {
			t.RotateRow(cmd.Value1, cmd.Value2)
		} else {
			t.RotateCol(cmd.Value1, cmd.Value2)
		}
	}
}

func (t *Torus) NumberOn() int {
	count := 0

	curr := t.Origin.Left.Up

	for r := 0; r < t.Rows; r++ {
		curr = curr.Down
		for c := 0; c < t.Cols; c++ {
			curr = curr.Right
			if curr.On {
				count++
			}
		}
	}

	return count
}

type Command struct {
	Command  string // rect or rotate
	Command2 string // row or column
	Value1   int    // row/col or number of columns
	Value2   int    // rotation or number of rows
}

func ParseDataString(dat string) []Command {
	commands := []Command{}
	for _, line := range strings.Split(dat, "\n") {
		parts := strings.Split(line, " ")
		cmd := parts[0]
		switch cmd {
		case "rect":
			rc := strings.Split(parts[1], "x")
			v1, _ := strconv.Atoi(rc[0])
			v2, _ := strconv.Atoi(rc[1])

			commands = append(commands, Command{
				Command:  cmd,
				Command2: "",
				Value1:   v1,
				Value2:   v2,
			})
		case "rotate":
			v1, _ := strconv.Atoi(strings.Split(parts[2], "=")[1])
			v2, _ := strconv.Atoi(parts[4])

			commands = append(commands, Command{
				Command:  cmd,
				Command2: parts[1],
				Value1:   v1,
				Value2:   v2,
			})
		}
	}

	return commands
}

func LoadData(filename string) []Command {
	dat := lib.ReadFileData(filename)

	return ParseDataString(strings.Trim(string(dat), "\n"))
}

func RunPartA(filename string) {
	t := NewTorus(6, 50)
	commands := LoadData(filename)

	fmt.Println("starting applying commands")
	for _, cmd := range commands {
		t.ApplyCommand(cmd)
	}

	fmt.Println(t.NumberOn())
}

func RunPartB(filename string) {
	t := NewTorus(6, 50)
	commands := LoadData(filename)

	fmt.Println("starting applying commands")
	for _, cmd := range commands {
		t.ApplyCommand(cmd)
	}

	t.Print()
}
