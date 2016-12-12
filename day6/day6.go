package day6

import (
	"fmt"
	"sort"
	"strings"

	"github.com/gsmcwhirter/advent2016/lib"
)

type Message struct {
	Chars []rune
}

func ParseDataString(dat string) []Message {
	messages := []Message{}
	for _, line := range strings.Split(dat, "\n") {
		messages = append(messages, Message{[]rune(line)})
	}

	return messages
}

//LoadData stuff
func LoadData(filename string) []Message {
	dat := lib.ReadFileData(filename)

	return ParseDataString(strings.Trim(string(dat), "\n"))
}

func RunPartA(filename string) {
	messages := LoadData(filename)

	charCounts := make([]map[rune]int, len(messages[0].Chars))

	for i := 0; i < len(charCounts); i++ {
		charCounts[i] = map[rune]int{}
	}

	for _, msg := range messages {
		for i, char := range msg.Chars {
			_, exists := charCounts[i][char]
			if !exists {
				charCounts[i][char] = 0
			}

			charCounts[i][char]++
		}
	}

	for _, counts := range charCounts {
		runeCounts := lib.NewRuneCounts(counts)
		rcs := lib.RuneCountSorter{Counts: runeCounts}
		sort.Sort(rcs)
		fmt.Print(string(rcs.Counts[0].Rune))
	}
	fmt.Println()

}

func RunPartB(filename string) {
	messages := LoadData(filename)

	charCounts := make([]map[rune]int, len(messages[0].Chars))

	for i := 0; i < len(charCounts); i++ {
		charCounts[i] = map[rune]int{}
	}

	for _, msg := range messages {
		for i, char := range msg.Chars {
			_, exists := charCounts[i][char]
			if !exists {
				charCounts[i][char] = 0
			}

			charCounts[i][char]++
		}
	}

	for _, counts := range charCounts {
		runeCounts := lib.NewRuneCounts(counts)
		rcs := lib.RuneCountSorter{Counts: runeCounts}
		sort.Sort(rcs)
		fmt.Print(string(rcs.Counts[len(rcs.Counts)-1].Rune))
	}
	fmt.Println()
}
