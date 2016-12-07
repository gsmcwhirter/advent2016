package main

import (
	"os"

	"github.com/gsmcwhirter/adventofcode/day1"
	"github.com/gsmcwhirter/adventofcode/day2"
	"github.com/gsmcwhirter/adventofcode/day3"
	"github.com/gsmcwhirter/adventofcode/day4"
	"github.com/gsmcwhirter/adventofcode/day5"
	"github.com/gsmcwhirter/adventofcode/day6"
	"github.com/gsmcwhirter/adventofcode/day7"
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
	}

}
