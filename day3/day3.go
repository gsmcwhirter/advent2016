package day3

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/gsmcwhirter/advent2016/lib"
)

type Triangle struct {
	side1 int
	side2 int
	side3 int
}

func NewTriangle(sides [3]int) Triangle {
	sort.Ints(sides[:])
	triangle := Triangle{sides[0], sides[1], sides[2]}
	return triangle
}

func (triangle *Triangle) isValid() bool {
	return triangle.side1+triangle.side2 > triangle.side3
}

func ParseDataString(dat string) []Triangle {
	triangles := []Triangle{}
	for _, line := range strings.Split(dat, "\n") {
		line = strings.Trim(line, " ")
		ints := [3]int{0, 0, 0}
		i := 0
		for _, intin := range strings.Split(line, " ") {
			if intin == "" {
				continue
			}
			parsedInt, _ := strconv.Atoi(intin)
			ints[i] = parsedInt
			i++
		}
		triangles = append(triangles, NewTriangle(ints))
	}

	return triangles
}

func ParseDataString2(dat string) []Triangle {
	triangles := []Triangle{}
	cols := [3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

	for lineNum, line := range strings.Split(dat, "\n") {
		if lineNum%3 == 0 {
			cols = [3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
		}

		fmt.Println(lineNum)
		sideNum := lineNum % 3
		fmt.Println(sideNum)
		triNum := 0
		for _, intin := range strings.Split(strings.Trim(line, " "), " ") {
			if intin == "" {
				continue
			}
			parsedInt, _ := strconv.Atoi(intin)
			cols[triNum][sideNum] = parsedInt

			fmt.Println(triNum)
			fmt.Println(cols)

			triNum++
		}

		if sideNum == 2 {
			triangles = append(triangles, NewTriangle(cols[0]))
			triangles = append(triangles, NewTriangle(cols[1]))
			triangles = append(triangles, NewTriangle(cols[2]))
		}

	}

	return triangles
}

//LoadData stuff
func LoadData(filename string) []Triangle {
	dat := lib.ReadFileData(filename)

	return ParseDataString(string(dat))
}

func LoadData2(filename string) []Triangle {
	dat := lib.ReadFileData(filename)

	return ParseDataString2(string(dat))
}

func RunPartA(filename string) {
	triangles := LoadData(filename)

	count := 0
	for _, triangle := range triangles {
		if triangle.isValid() {
			count++
		}
	}

	fmt.Println(count)
}

func RunPartB(filename string) {
	triangles := LoadData2(filename)

	count := 0
	for _, triangle := range triangles {
		if triangle.isValid() {
			count++
		}
	}

	fmt.Println(count)
}
