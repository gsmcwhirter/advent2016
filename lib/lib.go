package lib

import (
	"io/ioutil"
	"math"
)

//Check causes a panic on errors
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFileData(filename string) string {
	dat, err := ioutil.ReadFile(filename)
	Check(err)

	return string(dat)
}

func IntMax(a int, b int) int {
	if a >= b {
		return a
	}

	return b
}

func IntMin(a int, b int) int {
	if a <= b {
		return a
	}

	return b
}

func IntAbs(val int) int {
	return int(math.Abs(float64(val)))
}

func IntClamp(val int, min int, max int) int {
	if val < min {
		return min
	}

	if val > max {
		return max
	}

	return val
}

type RuneCount struct {
	Rune  rune
	Count int
}

func NewRuneCounts(counts map[rune]int) []RuneCount {
	runeCounts := []RuneCount{}
	for key, val := range counts {
		runeCounts = append(runeCounts, RuneCount{key, val})
	}

	return runeCounts
}

type RuneCountSorter struct {
	Counts []RuneCount
}

func (rcs RuneCountSorter) Len() int {
	return len(rcs.Counts)
}

func (rcs RuneCountSorter) Less(i, j int) bool {
	rcA := rcs.Counts[i]
	rcB := rcs.Counts[j]

	if rcA.Count > rcB.Count {
		return true
	}

	if rcA.Count < rcB.Count {
		return false
	}

	if rcA.Rune < rcB.Rune {
		return true
	}

	return false
}

func (rcs RuneCountSorter) Swap(i, j int) {
	var tmp RuneCount

	tmp = rcs.Counts[i]
	rcs.Counts[i] = rcs.Counts[j]
	rcs.Counts[j] = tmp
}
