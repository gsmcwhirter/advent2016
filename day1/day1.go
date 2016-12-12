package day1

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/gsmcwhirter/advent2016/lib"
)

//ParseDataString splits on spaces
func ParseDataString(datString string) []string {
	return strings.Split(datString, ", ")
}

//Facing does stuff
func Facing(index int, direction int) string {
	switch pair := [2]int{index, direction}; pair {
	case [2]int{0, 1}:
		return "N"
	case [2]int{0, -1}:
		return "S"
	case [2]int{1, 1}:
		return "E"
	case [2]int{1, -1}:
		return "W"
	}

	return "unknown"
}

//Distance does stuff
func Distance(travelled [2]int) int {
	return int(math.Abs(float64(travelled[0])) + math.Abs(float64(travelled[1])))
}

//PrintCurrentState prints the current state
func PrintCurrentState(index int, direction int, location [2]int) {
	fmt.Print("current facing: ")
	fmt.Print(Facing(index, direction))
	fmt.Print(", current distance: ")
	fmt.Print(Distance(location))
	fmt.Print(", location: ")
	fmt.Print(location)
	fmt.Print("\n")
}

//NextLocation steps to the next location
func NextLocation(step string, index int, direction int, location [2]int) (int, int, [2]int, [][2]int) {
	PrintCurrentState(index, direction, location)
	fmt.Print("step: ")
	fmt.Print(step)
	fmt.Print("\n")

	visited := [][2]int{}

	stepRunes := []rune(step)
	dist, _ := strconv.Atoi(string(stepRunes[1:]))

	switch string(stepRunes[0]) {
	case "L":
		if index == 0 {
			direction *= -1
		}
	case "R":
		if index == 1 {
			direction *= -1
		}
	}

	index = 1 - index

	for i := 0; i < dist; i++ {
		location[index] += direction
		visited = append(visited, location)
	}

	return index, direction, location, visited
}

//LoadData stuff
func LoadData(filename string) []string {
	dat := lib.ReadFileData(filename)

	return ParseDataString(string(dat))
}

//RunPartA is a "main"
func RunPartA(filename string) {
	steps := LoadData(filename)
	index := 0
	direction := 1           // 1 or -1 indicating N/E or S/W
	location := [2]int{0, 0} // N/S, E/W

	for _, step := range steps {
		index, direction, location, _ = NextLocation(step, index, direction, location)
	}

	PrintCurrentState(index, direction, location)
}

//RunPartB is a "main"
func RunPartB(filename string) {
	steps := LoadData(filename)
	index := 0
	direction := 1           // 1 or -1 indicating N/E or S/W
	location := [2]int{0, 0} // N/S, E/W

	locsVisited := [][2]int{}
	visited := map[[2]int]bool{}
	visited[location] = true

	for _, step := range steps {
		index, direction, location, locsVisited = NextLocation(step, index, direction, location)
		for _, loc := range locsVisited {
			_, present := visited[loc]
			if present {
				fmt.Print("HQ: ")
				fmt.Print(loc)
				fmt.Print("\n")
				PrintCurrentState(index, direction, loc)
				return
			}

			visited[loc] = true
		}
	}
}
