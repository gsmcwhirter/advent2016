package day10

import (
	"fmt"
	"strings"

	"regexp"

	"strconv"

	"github.com/gsmcwhirter/advent2016/lib"
)

// type Chip struct {
// 	Value int
// }

type Bot struct {
	Id         string
	IsBot      bool
	Chips      []int
	LowOutput  *Bot
	HighOutput *Bot
}

func (b *Bot) Print() {
	typeWord := "Bot"
	if !b.IsBot {
		typeWord = "Output"
	}

	fmt.Printf("%s holding chips %v; low -> %v, high -> %v\n", typeWord, b.Chips, b.LowOutput, b.HighOutput)
}

func (b *Bot) HasTarget(cmpTarget1, cmpTarget2 int) bool {
	if len(b.Chips) < 2 {
		return false
	}

	high := lib.IntMax(b.Chips[0], b.Chips[1])
	low := lib.IntMin(b.Chips[0], b.Chips[1])

	if (high == cmpTarget1 && low == cmpTarget2) || (high == cmpTarget2 && low == cmpTarget1) {
		return true
	}

	return false
}

func (b *Bot) Process() bool {
	if len(b.Chips) < 2 {
		return false
	}

	if len(b.Chips) > 2 {
		fmt.Printf("Bot %v has too many chips?\n", b)
	}

	high := lib.IntMax(b.Chips[0], b.Chips[1])
	low := lib.IntMin(b.Chips[0], b.Chips[1])

	b.HighOutput.AddChip(high)
	b.LowOutput.AddChip(low)

	b.Chips = b.Chips[2:]

	return true
}

func (b *Bot) AddChip(val int) {
	b.Chips = append(b.Chips, val)
}

type Factory struct {
	Bots map[string]*Bot
	// Inputs  map[string]*Bot
	Outputs map[string]*Bot
}

func (f *Factory) GetBot(id string) *Bot {
	_, exists := f.Bots[id]

	if !exists {
		f.Bots[id] = &Bot{
			Id:         id,
			IsBot:      true,
			Chips:      []int{},
			LowOutput:  nil,
			HighOutput: nil,
		}
	}

	return f.Bots[id]
}

func (f *Factory) GetOutput(id string) *Bot {
	_, exists := f.Outputs[id]

	if !exists {
		f.Outputs[id] = &Bot{
			Id:         id,
			IsBot:      false,
			Chips:      []int{},
			LowOutput:  nil,
			HighOutput: nil,
		}
	}

	return f.Outputs[id]
}

func (f *Factory) parseLine(line string) {
	valueRegex := regexp.MustCompile("^value (\\d+) goes to bot (\\d+)$")
	ruleRegex := regexp.MustCompile("^bot (\\d+) gives low to (bot|output) (\\d+) and high to (bot|output) (\\d+)$")

	valueMatches := valueRegex.FindAllStringSubmatch(line, -1)

	if valueMatches != nil {
		bot := f.GetBot(valueMatches[0][2])

		chipValue, _ := strconv.Atoi(valueMatches[0][1])
		bot.AddChip(chipValue)

		return
	}

	ruleMatches := ruleRegex.FindAllStringSubmatch(line, -1)

	if ruleMatches != nil {
		bot := f.GetBot(ruleMatches[0][1])

		var lowTarget *Bot
		var highTarget *Bot

		switch ruleMatches[0][2] {
		case "bot":
			lowTarget = f.GetBot(ruleMatches[0][3])
		case "output":
			lowTarget = f.GetOutput(ruleMatches[0][3])
		}

		bot.LowOutput = lowTarget

		switch ruleMatches[0][4] {
		case "bot":
			highTarget = f.GetBot(ruleMatches[0][5])
		case "output":
			highTarget = f.GetOutput(ruleMatches[0][5])
		}

		bot.HighOutput = highTarget

		return
	}

	fmt.Printf("Error parsing line '%s'\n", line)
}

func ParseDataString(dat string) Factory {
	factory := Factory{
		map[string]*Bot{},
		// map[string]Bot{},
		map[string]*Bot{},
	}

	for _, line := range strings.Split(dat, "\n") {
		factory.parseLine(line)
	}

	return factory
}

func LoadData(filename string) Factory {
	dat := lib.ReadFileData(filename)

	return ParseDataString(strings.Trim(string(dat), "\n"))
}

func RunPartA(filename string) {
	factory := LoadData(filename)

	found := false
	moves := 1
	for !found && moves > 0 {
		moves = 0
		for _, bot := range factory.Bots {
			found = bot.HasTarget(17, 61)

			if found {
				fmt.Println(bot.Id)
				break
			}

			if bot.Process() {
				moves++
			}
		}
	}
}

func RunPartB(filename string) {
	factory := LoadData(filename)

	moves := 1
	for moves > 0 {
		moves = 0

		out0Chips := factory.GetOutput("0").Chips
		out1Chips := factory.GetOutput("1").Chips
		out2Chips := factory.GetOutput("2").Chips
		if len(out0Chips) > 0 && len(out1Chips) > 0 && len(out2Chips) > 0 {
			fmt.Println(out0Chips[0] * out1Chips[0] * out2Chips[0])
			break
		}

		for _, bot := range factory.Bots {
			if bot.Process() {
				moves++
			}
		}
	}
}
