package day10

import (
	"fmt"
	"strings"

	"regexp"

	"github.com/gsmcwhirter/advent2016/lib"
)

type Chip struct {
	Value string
}

type Bot struct {
	IsBot      bool
	Chips      []Chip
	LowOutput  *Bot
	HighOutput *Bot
}

type Factory struct {
	Bots map[string]Bot
	// Inputs  map[string]Bot
	Outputs map[string]Bot
}

func (f *Factory) GetBot(id string) *Bot {
	_, exists := f.Bots[id]

	if !exists {
		f.Bots[id] = Bot{
			IsBot:      true,
			Chips:      []Chip{},
			LowOutput:  nil,
			HighOutput: nil,
		}
	}

	bot := f.Bots[id]
	return &bot
}

func (f *Factory) GetOutput(id string) *Bot {
	_, exists := f.Outputs[id]

	if !exists {
		f.Outputs[id] = Bot{
			IsBot:      false,
			Chips:      []Chip{},
			LowOutput:  nil,
			HighOutput: nil,
		}
	}

	output := f.Outputs[id]
	return &output
}

func (f *Factory) parseLine(line string) {
	valueRegex := regexp.MustCompile("value (\\d+) goes to (bot|output) (\\d+)")
	ruleRegex := regexp.MustCompile("bot (\\d+) gives low to (bot|output) (\\d+) and high to (bot|output) (\\d+)")

	valueMatches := valueRegex.FindAllStringSubmatch(line, -1)

	if valueMatches != nil {
		bot := f.GetBot(valueMatches[0][3])

		bot.Chips = append(bot.Chips, Chip{})
		return
	}

	ruleMatches := ruleRegex.FindAllStringSubmatch(line, -1)

	if ruleMatches != nil {

		return
	}

	fmt.Printf("Error parsing line '%s'\n", line)
}

func ParseDataString(dat string) Factory {
	factory := Factory{
		map[string]Bot{},
		// map[string]Bot{},
		map[string]Bot{},
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

}

func RunPartB(filename string) {

}
