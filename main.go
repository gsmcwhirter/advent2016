package main

import (
	"os"

	"github.com/gsmcwhirter/advent2016/day1"
	"github.com/gsmcwhirter/advent2016/day10"
	"github.com/gsmcwhirter/advent2016/day11"
	"github.com/gsmcwhirter/advent2016/day12"
	"github.com/gsmcwhirter/advent2016/day13"
	"github.com/gsmcwhirter/advent2016/day14"
	"github.com/gsmcwhirter/advent2016/day15"
	"github.com/gsmcwhirter/advent2016/day16"
	"github.com/gsmcwhirter/advent2016/day17"
	"github.com/gsmcwhirter/advent2016/day18"
	"github.com/gsmcwhirter/advent2016/day19"
	"github.com/gsmcwhirter/advent2016/day2"
	"github.com/gsmcwhirter/advent2016/day20"
	"github.com/gsmcwhirter/advent2016/day21"
	"github.com/gsmcwhirter/advent2016/day22"
	"github.com/gsmcwhirter/advent2016/day23"
	"github.com/gsmcwhirter/advent2016/day24"
	"github.com/gsmcwhirter/advent2016/day25"
	"github.com/gsmcwhirter/advent2016/day3"
	"github.com/gsmcwhirter/advent2016/day4"
	"github.com/gsmcwhirter/advent2016/day5"
	"github.com/gsmcwhirter/advent2016/day6"
	"github.com/gsmcwhirter/advent2016/day7"
	"github.com/gsmcwhirter/advent2016/day8"
	"github.com/gsmcwhirter/advent2016/day9"
)

func main() {
	day := string(os.Args[1])
	filename := os.Args[2]

	switch day {
	case "1a":
		day1.RunPartA(filename)
	case "1b":
		day1.RunPartB(filename)
	case "2a":
		day2.RunPartA(filename)
	case "2b":
		day2.RunPartB(filename)
	case "3a":
		day3.RunPartA(filename)
	case "3b":
		day3.RunPartB(filename)
	case "4a":
		day4.RunPartA(filename)
	case "4b":
		day4.RunPartB(filename)
	case "5a":
		day5.RunPartA(filename)
	case "5b":
		day5.RunPartB(filename)
	case "6a":
		day6.RunPartA(filename)
	case "6b":
		day6.RunPartB(filename)
	case "7a":
		day7.RunPartA(filename)
	case "7b":
		day7.RunPartB(filename)
	case "8a":
		day8.RunPartA(filename)
	case "8b":
		day8.RunPartB(filename)
	case "9a":
		day9.RunPartA(filename)
	case "9b":
		day9.RunPartB(filename)
	case "10a":
		day10.RunPartA(filename)
	case "10b":
		day10.RunPartB(filename)
	case "11a":
		day11.RunPartA(filename)
	case "11b":
		day11.RunPartB(filename)
	case "12a":
		day12.RunPartA(filename)
	case "12b":
		day12.RunPartB(filename)
	case "13a":
		day13.RunPartA(filename)
	case "13b":
		day13.RunPartB(filename)
	case "14a":
		day14.RunPartA(filename)
	case "14b":
		day14.RunPartB(filename)
	case "15a":
		day15.RunPartA(filename)
	case "15b":
		day15.RunPartB(filename)
	case "16a":
		day16.RunPartA(filename)
	case "16b":
		day16.RunPartB(filename)
	case "17a":
		day17.RunPartA(filename)
	case "17b":
		day17.RunPartB(filename)
	case "18a":
		day18.RunPartA(filename)
	case "18b":
		day18.RunPartB(filename)
	case "19a":
		day19.RunPartA(filename)
	case "19b":
		day19.RunPartB(filename)
	case "20a":
		day20.RunPartA(filename)
	case "20b":
		day20.RunPartB(filename)
	case "21a":
		day21.RunPartA(filename)
	case "21b":
		day21.RunPartB(filename)
	case "22a":
		day22.RunPartA(filename)
	case "22b":
		day22.RunPartB(filename)
	case "23a":
		day23.RunPartA(filename)
	case "23b":
		day23.RunPartB(filename)
	case "24a":
		day24.RunPartA(filename)
	case "24b":
		day24.RunPartB(filename)
	case "25a":
		day25.RunPartA(filename)
	case "25b":
		day25.RunPartB(filename)
	}

}
