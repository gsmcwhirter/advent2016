package day7

import (
	"strings"

	"fmt"
	"regexp"

	"github.com/gsmcwhirter/advent2016/lib"
)

type CharSequenceList struct {
	SeqSize   int
	Sequences [][]rune
}

func (csl *CharSequenceList) hasABBA() bool {
	for i, seq := range csl.Sequences[:len(csl.Sequences)-csl.SeqSize] {
		isABBA := true
		for j := 0; j < csl.SeqSize; j++ {
			for k := j + 1; k < csl.SeqSize; k++ {
				if seq[j] == seq[k] {
					isABBA = false
				}
			}

			if seq[j] != csl.Sequences[i+csl.SeqSize][csl.SeqSize-j-1] {
				isABBA = false
			}
		}

		if isABBA {
			return true
		}
	}

	return false
}

func (csl *CharSequenceList) getABAs() [][]rune {
	abas := [][]rune{}

	for _, seq := range csl.Sequences {
		if seq[0] == seq[2] && seq[0] != seq[1] {
			abas = append(abas, seq)
		}
	}

	return abas
}

func StringToCharSequenceList(seqSize int, str string) CharSequenceList {
	strRunes := []rune(str)
	numSequences := len(strRunes) - seqSize + 1
	if numSequences < 0 {
		numSequences = 0
	}

	seqs := make([][]rune, numSequences)

	for i := 0; i < numSequences; i++ {
		seqs[i] = strRunes[i : i+seqSize]
	}

	return CharSequenceList{
		SeqSize:   seqSize,
		Sequences: seqs,
	}
}

type IPv7 struct {
	Hypernets []CharSequenceList
	Supernets []CharSequenceList
}

func (ip *IPv7) supportsTLS() bool {
	for _, hypernet := range ip.Hypernets {
		if hypernet.hasABBA() {
			return false
		}
	}

	for _, supernet := range ip.Supernets {
		if supernet.hasABBA() {
			return true
		}
	}

	return false
}

func (ip *IPv7) supportsSSL() bool {
	hyperABAs := [][]rune{}

	for _, hypernet := range ip.Hypernets {
		hyperABAs = append(hyperABAs, hypernet.getABAs()...)
	}

	superABAs := [][]rune{}
	for _, supernet := range ip.Supernets {
		superABAs = append(superABAs, supernet.getABAs()...)
	}

	for _, h := range hyperABAs {
		for _, s := range superABAs {
			if h[0] == s[1] && h[1] == s[0] {
				return true
			}
		}
	}

	return false
}

func StringToIPv7(line string, seqSize int) IPv7 {
	hypernetRegex := regexp.MustCompile("\\[([^\\]]+)\\]")
	supernetRegex := regexp.MustCompile("(^|\\])([^\\[]+)")

	hypernetStrings := hypernetRegex.FindAllString(line, -1)
	hypernets := make([]CharSequenceList, len(hypernetStrings))

	for i, hypernet := range hypernetStrings {
		hypernets[i] = StringToCharSequenceList(seqSize, strings.Trim(hypernet, "[]"))
	}

	supernetStrings := supernetRegex.FindAllString(line, -1)
	supernets := make([]CharSequenceList, len(supernetStrings))

	for i, supernet := range supernetStrings {
		supernets[i] = StringToCharSequenceList(seqSize, strings.Trim(supernet, "[]"))
	}

	return IPv7{
		Hypernets: hypernets,
		Supernets: supernets,
	}
}

func ParseDataString(dat string, seqSize int) []IPv7 {
	ips := []IPv7{}
	for _, line := range strings.Split(dat, "\n") {
		ips = append(ips, StringToIPv7(line, seqSize))
	}

	return ips
}

func LoadData(filename string, seqSize int) []IPv7 {
	dat := lib.ReadFileData(filename)

	return ParseDataString(strings.Trim(string(dat), "\n"), seqSize)
}

func RunPartA(filename string) {
	ips := LoadData(filename, 2)
	count := 0

	for _, ip := range ips {
		if ip.supportsTLS() {
			count++
		}
	}

	fmt.Println(count)
}

func RunPartB(filename string) {
	ips := LoadData(filename, 3)
	count := 0

	for _, ip := range ips {
		if ip.supportsSSL() {
			count++
		}
	}

	fmt.Println(count)
}
