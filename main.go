package main

import (
	"os"

	"github.com/gsmcwhirter/advent2016/day1"
	"github.com/gsmcwhirter/advent2016/day10"
	"github.com/gsmcwhirter/advent2016/day11"
	"github.com/gsmcwhirter/advent2016/day2"
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
	}

}
