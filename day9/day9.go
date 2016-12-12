package day9

import (
	"fmt"
	"strings"

	"regexp"

	"strconv"

	"github.com/gsmcwhirter/advent2016/lib"
)

func DecompressSize(inStr string, recursive bool) int {
	if inStr == "" {
		return 0
	}

	size := 0
	inRunes := []rune(inStr)

	// fmt.Printf("Decompressing %v runes\n", len(inRunes))

	markerRegex := regexp.MustCompile("\\((\\d+)x(\\d+)\\)")
	// This is a 2d array of locations
	// Each element is 2xSubgroups in length
	// The elements of each element are paired start/stop indices
	maybeMarkers := markerRegex.FindAllStringSubmatchIndex(inStr, -1)

	if maybeMarkers == nil {
		return len(inStr)
	}

	inIndex := 0
	markerIndex := 0
	var markerStart int
	var markerStop int
	var markerChars int
	var markerDups int
	for inIndex < len(inRunes) {
		if markerIndex < len(maybeMarkers) {
			markerStart = maybeMarkers[markerIndex][0]
			markerStop = maybeMarkers[markerIndex][1]
			markerChars, _ = strconv.Atoi(string(inRunes[maybeMarkers[markerIndex][2]:maybeMarkers[markerIndex][3]]))
			markerDups, _ = strconv.Atoi(string(inRunes[maybeMarkers[markerIndex][4]:maybeMarkers[markerIndex][5]]))

			// fmt.Printf("marker info: %v %v %v %v\n", markerStart, markerStop, markerChars, markerDups)
		} else {
			markerStart = len(inRunes)
			markerStop = len(inRunes)
			markerChars = 0
			markerDups = 0

			// fmt.Println("no marker info")
		}

		// Proceed up to the next marker
		// outRunes = append(outRunes, inRunes[inIndex:markerStart]...)
		size += markerStart - inIndex

		// fmt.Printf("non marker: %v %v %v %v\n", inIndex, markerStart, inRunes[inIndex:markerStart], outRunes)

		// Handle the next marker
		// var toDup []rune
		var dupSize int
		if recursive {
			// toDup = []rune(Decompress(string(inRunes[markerStop:markerStop+markerChars]), true))
			dupSize = DecompressSize(string(inRunes[markerStop:markerStop+markerChars]), true)
		} else {
			// toDup = inRunes[markerStop : markerStop+markerChars]
			dupSize = markerChars
		}

		size += dupSize * markerDups

		// for i := 0; i < markerDups; i++ {
		// 	outRunes = append(outRunes, toDup...)
		// 	// fmt.Printf("marker: %v %v %v %v %v %v\n", markerStart, markerStop, markerChars, markerDups, inRunes[markerStop:markerStop+markerChars], outRunes)
		// }

		// Reset the general position
		inIndex = markerStop + markerChars

		// Reset the next marker
		nextMarkerIndex := len(maybeMarkers)
		for i := markerIndex; i < len(maybeMarkers); i++ {
			if maybeMarkers[i][0] >= inIndex {
				nextMarkerIndex = i
				break
			}
		}

		markerIndex = nextMarkerIndex
	}

	// return string(outRunes)
	return size
}

func LoadData(filename string) string {
	dat := lib.ReadFileData(filename)

	return strings.Trim(string(dat), "\n")
}

func RunPartA(filename string) {
	input := LoadData(filename)
	fmt.Println(DecompressSize(input, false))
}

func RunPartB(filename string) {
	input := LoadData(filename)
	fmt.Println(DecompressSize(input, true))
}
