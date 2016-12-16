package day14

import (
	"fmt"
	"strings"

	"crypto/md5"

	"github.com/gsmcwhirter/advent2016/lib"
)

func LoadData(filename string) string {
	dat := lib.ReadFileData(filename)

	return strings.Trim(string(dat), "\n")
}

func GenerateHash(salt string, index int) []rune {
	return []rune(fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%v%v", salt, index)))))
}

func GenerateStretchedHash(salt string, index int) []rune {
	hash := string(GenerateHash(salt, index))
	for i := 0; i < 2016; i++ {
		hash = fmt.Sprintf("%x", md5.Sum([]byte(hash)))
	}

	return []rune(hash)
}

func HashHas5Consecutive(hash []rune, target rune) bool {
	for i := 0; i < len(hash)-4; i++ {
		fiveOkay := true
		for j := 0; j < 5; j++ {
			if hash[i+j] != target {
				fiveOkay = false
				break
			}
		}

		if fiveOkay {
			return true
		}
	}

	return false
}

func HashIndicatesKey(salt string, index int, hash []rune, HashGen func(string, int) []rune, hashHistory map[int][]rune) (bool, map[int][]rune) {
	extraHashes := map[int][]rune{}
	keyRune := 'x' // hash is hex, so this will never be valid
	for i, iRune := range hash[:len(hash)-2] {
		if hash[i+1] == iRune && hash[i+2] == iRune {
			keyRune = iRune
			break
		}
	}

	if keyRune == 'x' {
		return false, extraHashes
	}

	for j := 1; j <= 1000; j++ {
		nextHash, exists := hashHistory[index+j]
		if !exists {
			nextHash = HashGen(salt, index+j)
			extraHashes[index+j] = nextHash
		}

		if HashHas5Consecutive(nextHash, keyRune) {
			return true, extraHashes
		}
	}

	return false, extraHashes
}

func RunPartA(filename string) {
	salt := LoadData(filename)

	keysFound := 0

	var lastKey int

	index := 0
	for keysFound < 64 {
		hash := GenerateHash(salt, index)
		isKey, _ := HashIndicatesKey(salt, index, hash, GenerateHash, map[int][]rune{})
		if isKey {
			lastKey = index
			keysFound++
			fmt.Printf("Key %v is %v\n", keysFound, lastKey)
		}

		index++
	}

	fmt.Println(lastKey)
}

func RunPartB(filename string) {
	salt := LoadData(filename)

	keysFound := 0

	var lastKey int

	hashLookup := map[int][]rune{
		0: GenerateStretchedHash(salt, 0),
	}

	index := 0
	for keysFound < 64 {
		hash, exists := hashLookup[index]
		if !exists {
			hash = GenerateStretchedHash(salt, index)
		}

		isKey, extraHashes := HashIndicatesKey(salt, index, hash, GenerateStretchedHash, hashLookup)

		if isKey {
			lastKey = index
			keysFound++
			fmt.Printf("Key %v is %v\n", keysFound, lastKey)
		}

		for k, v := range extraHashes {
			hashLookup[k] = v
		}

		index++
	}

	fmt.Println(lastKey)
}
