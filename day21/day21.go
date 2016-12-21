package day21

import (
	"fmt"
	"strings"

	"regexp"

	"strconv"

	"github.com/gsmcwhirter/advent2016/lib"
)

type ScramblerState struct {
	Runes      []rune
	StartIndex int
	Length     int
}

func NewScrambler(word string) ScramblerState {
	runes := []rune(word)
	return ScramblerState{
		Runes:      runes,
		StartIndex: 0,
		Length:     len(runes),
	}
}

func (s *ScramblerState) RealIndex(index int) int {
	return lib.IntMod(index+s.StartIndex, s.Length)
}

func (s *ScramblerState) RuneAt(index int) rune {
	return s.Runes[s.RealIndex(index)]
}

func (s *ScramblerState) SetRune(index int, letter rune) {
	s.Runes[s.RealIndex(index)] = letter
}

func (s *ScramblerState) SwapPos(i, j int) {
	tmp := s.RuneAt(i)
	s.SetRune(i, s.RuneAt(j))
	s.SetRune(j, tmp)
}

func (s *ScramblerState) SwapLetter(a, b rune) {
	for i := 0; i < s.Length; i++ {
		if s.Runes[i] == a {
			s.Runes[i] = b
			continue
		}
		if s.Runes[i] == b {
			s.Runes[i] = a
			continue
		}
	}
}

func (s *ScramblerState) RotateRight(num int) {
	s.StartIndex = lib.IntMod(s.StartIndex-num, s.Length)
}

func (s *ScramblerState) RotateLeft(num int) {
	s.StartIndex = lib.IntMod(s.StartIndex+num, s.Length)
}

func (s *ScramblerState) RotateLetter(target rune) {
	num := 1
	for i := 0; i < s.Length; i++ {
		if s.RuneAt(i) == target {
			num += i
			break
		}
	}

	if num > 4 {
		num++
	}

	s.RotateRight(num)
}

func (s *ScramblerState) UndoRotateLetter(target rune) {
	index := -1
	for i := 0; i < s.Length; i++ {
		if s.RuneAt(i) == target {
			index = i
			break
		}
	}

	left := -1
	switch {
	case index == 0:
		left = 9
	case index%2 == 0:
		left = 5 + index/2
	case index%2 == 1:
		left = (index + 1) / 2
	}

	s.RotateLeft(left)
}

func (s *ScramblerState) ReverseSlice(i, j int) {
	for i < j {
		s.SwapPos(i, j)
		i++
		j--
	}
}

func (s *ScramblerState) MovePos(from, to int) {
	tmp := s.RuneAt(from)

	// Remove that letter and insert at the end
	for i := from; i < s.Length-1; i++ {
		s.SetRune(i, s.RuneAt(i+1))
	}

	// Re-Insert
	for i := s.Length - 1; i > to; i-- {
		s.SetRune(i, s.RuneAt(i-1))
	}

	s.SetRune(to, tmp)
}

func (s *ScramblerState) ApplyCommand(cmd Command) {
	switch cmd.Type {
	case SWAP_POS:
		s.SwapPos(cmd.IArg1, cmd.IArg2)
	case SWAP_LETTER:
		s.SwapLetter(cmd.RArg1, cmd.RArg2)
	case REVERSE:
		s.ReverseSlice(cmd.IArg1, cmd.IArg2)
	case ROTATE_LEFT:
		s.RotateLeft(cmd.IArg1)
	case ROTATE_RIGHT:
		s.RotateRight(cmd.IArg1)
	case ROTATE_LETTER:
		s.RotateLetter(cmd.RArg1)
	case MOVE:
		s.MovePos(cmd.IArg1, cmd.IArg2)
	}
}

func (s *ScramblerState) UndoCommand(cmd Command) {
	switch cmd.Type {
	case SWAP_POS:
		s.SwapPos(cmd.IArg1, cmd.IArg2)
	case SWAP_LETTER:
		s.SwapLetter(cmd.RArg1, cmd.RArg2)
	case REVERSE:
		s.ReverseSlice(cmd.IArg1, cmd.IArg2)
	case ROTATE_LEFT:
		s.RotateRight(cmd.IArg1)
	case ROTATE_RIGHT:
		s.RotateLeft(cmd.IArg1)
	case ROTATE_LETTER:
		s.UndoRotateLetter(cmd.RArg1)
	case MOVE:
		s.MovePos(cmd.IArg2, cmd.IArg1)
	}
}

func (s *ScramblerState) Print() {
	for i := 0; i < s.Length; i++ {
		fmt.Print(string(s.RuneAt(i)))
	}

	fmt.Printf("  (%v @ %v)\n", string(s.Runes), s.StartIndex)
}

type CommandType int

const (
	SWAP_POS CommandType = iota
	SWAP_LETTER
	REVERSE
	ROTATE_LEFT
	ROTATE_RIGHT
	ROTATE_LETTER
	MOVE
	UNKNOWN
)

type Command struct {
	Type  CommandType
	Text  string
	IArg1 int
	IArg2 int
	RArg1 rune
	RArg2 rune
}

func NewCommand(line string) Command {
	swapPosRegex := regexp.MustCompile("swap position (\\d+) with position (\\d+)")
	swapLetterRegex := regexp.MustCompile("swap letter (.{1}) with letter (.{1})")
	reverseRegex := regexp.MustCompile("reverse positions (\\d+) through (\\d+)")
	rotateLRRegex := regexp.MustCompile("rotate (left|right) (\\d+) step")
	rotateLetterRegex := regexp.MustCompile("rotate based on position of letter (.{1})")
	moveRegex := regexp.MustCompile("move position (\\d+) to position (\\d+)")

	switch {
	case swapPosRegex.MatchString(line):
		match := swapPosRegex.FindStringSubmatch(line)
		arg1, _ := strconv.Atoi(match[1])
		arg2, _ := strconv.Atoi(match[2])
		return Command{
			Type:  SWAP_POS,
			Text:  line,
			IArg1: arg1,
			IArg2: arg2,
			RArg1: ' ',
			RArg2: ' ',
		}

	case swapLetterRegex.MatchString(line):
		match := swapLetterRegex.FindStringSubmatch(line)
		return Command{
			Type:  SWAP_LETTER,
			Text:  line,
			IArg1: -1,
			IArg2: -1,
			RArg1: []rune(match[1])[0],
			RArg2: []rune(match[2])[0],
		}

	case reverseRegex.MatchString(line):
		match := reverseRegex.FindStringSubmatch(line)
		arg1, _ := strconv.Atoi(match[1])
		arg2, _ := strconv.Atoi(match[2])
		return Command{
			Type:  REVERSE,
			Text:  line,
			IArg1: arg1,
			IArg2: arg2,
			RArg1: ' ',
			RArg2: ' ',
		}

	case rotateLRRegex.MatchString(line):
		match := rotateLRRegex.FindStringSubmatch(line)
		arg1, _ := strconv.Atoi(match[2])

		switch match[1] {
		case "left":
			return Command{
				Type:  ROTATE_LEFT,
				Text:  line,
				IArg1: arg1,
				IArg2: -1,
				RArg1: ' ',
				RArg2: ' ',
			}
		case "right":
			return Command{
				Type:  ROTATE_RIGHT,
				Text:  line,
				IArg1: arg1,
				IArg2: -1,
				RArg1: ' ',
				RArg2: ' ',
			}
		default:
			fmt.Printf("Error understanding rotate left|right: '%v' in line '%v'\n", match[1], line)
		}

	case rotateLetterRegex.MatchString(line):
		match := rotateLetterRegex.FindStringSubmatch(line)
		return Command{
			Type:  ROTATE_LETTER,
			Text:  line,
			IArg1: -1,
			IArg2: -1,
			RArg1: []rune(match[1])[0],
			RArg2: ' ',
		}

	case moveRegex.MatchString(line):
		match := moveRegex.FindStringSubmatch(line)
		arg1, _ := strconv.Atoi(match[1])
		arg2, _ := strconv.Atoi(match[2])
		return Command{
			Type:  MOVE,
			Text:  line,
			IArg1: arg1,
			IArg2: arg2,
			RArg1: ' ',
			RArg2: ' ',
		}

	default:
		fmt.Printf("Unable to parse '%v'\n", line)
	}

	return Command{
		Type:  UNKNOWN,
		Text:  line,
		IArg1: -1,
		IArg2: -1,
		RArg1: ' ',
		RArg2: ' ',
	}
}

func ParseDataString(dat string) (string, []Command) {
	lines := strings.Split(dat, "\n")
	commands := make([]Command, len(lines)-1)

	for i, line := range lines[1:] {
		commands[i] = NewCommand(line)
	}

	return lines[0], commands
}

func LoadData(filename string) (string, []Command) {
	dat := lib.ReadFileData(filename)

	return ParseDataString(strings.Trim(string(dat), "\n"))
}

func RunPartA(filename string) {
	keyword, commands := LoadData(filename)
	scrambler := NewScrambler(keyword)

	scrambler.Print()

	for _, cmd := range commands {
		fmt.Printf("%v: ", cmd.Text)
		scrambler.ApplyCommand(cmd)
		scrambler.Print()
	}

	scrambler.Print()
}

func RunPartB(filename string) {
	keyword := "fbgdceah"
	_, commands := LoadData(filename)
	scrambler := NewScrambler(keyword)

	scrambler.Print()

	var cmd Command
	for i := len(commands) - 1; i >= 0; i-- {
		cmd = commands[i]
		fmt.Printf("%v: ", cmd.Text)
		scrambler.UndoCommand(cmd)
		scrambler.Print()
	}

	scrambler.Print()
}
