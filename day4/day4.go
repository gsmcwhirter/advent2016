package day4

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/gsmcwhirter/advent2016/lib"
)

type Room struct {
	Name     []string
	Id       int
	Checksum []rune
}

func (room *Room) isValid() bool {
	counts := map[rune]int{}
	for _, part := range room.Name {
		runes := []rune(part)
		for _, r := range runes {
			_, exists := counts[r]
			if !exists {
				counts[r] = 0
			}

			counts[r]++
		}
	}

	runeCounts := lib.NewRuneCounts(counts)
	rcs := lib.RuneCountSorter{Counts: runeCounts}
	sort.Sort(rcs)

	first5Runes := make([]rune, 5)
	for i, rc := range rcs.Counts[:5] {
		first5Runes[i] = rc.Rune
	}

	for i := 0; i < 5; i++ {
		if first5Runes[i] != room.Checksum[i] {
			return false
		}
	}

	return true
}

func (room *Room) DecryptedName() string {
	base := int([]byte("a")[0])
	space := []byte(" ")[0]

	nameBytes := []byte{}
	for _, part := range room.Name {
		for _, b := range []byte(part) {
			nameBytes = append(nameBytes, byte((int(b)-base+room.Id)%26+base))
		}
		nameBytes = append(nameBytes, space)
	}

	return strings.Trim(string(nameBytes), " ")
}

func NewRoom(line string) Room {
	parts := strings.Split(line, "-")
	idChecksumParts := strings.Split(strings.Trim(parts[len(parts)-1], "]"), "[")
	roomID, _ := strconv.Atoi(idChecksumParts[0])
	room := Room{parts[:len(parts)-1], roomID, []rune(idChecksumParts[1])}

	return room
}

func ParseDataString(dat string) []Room {
	rooms := []Room{}
	for _, line := range strings.Split(dat, "\n") {
		rooms = append(rooms, NewRoom(line))
	}
	return rooms
}

//LoadData stuff
func LoadData(filename string) []Room {
	dat := lib.ReadFileData(filename)

	return ParseDataString(string(dat))
}

func RunPartA(filename string) {
	rooms := LoadData(filename)

	validIDSum := 0
	for _, room := range rooms {
		if room.isValid() {
			validIDSum += room.Id
		}
	}

	fmt.Println(validIDSum)
}

func RunPartB(filename string) {
	rooms := LoadData(filename)

	for _, room := range rooms {
		if room.isValid() {
			fmt.Printf("%i: %s\n", room.Id, room.DecryptedName())
		}
	}
}
