package day5

import (
	"strings"

	"fmt"

	"crypto/md5"

	"github.com/gsmcwhirter/advent2016/lib"
)

//LoadData stuff
func LoadData(filename string) string {
	dat := lib.ReadFileData(filename)

	return strings.Trim(string(dat), "\n")
}

func RunPartA(filename string) {
	doorid := LoadData(filename)

	passChars := []byte{}
	i := 0
	zeroByte := []byte("0")[0]

	for len(passChars) < 8 {
		bytesToHash := []byte(fmt.Sprintf("%s%v", doorid, i))
		hash := []byte(fmt.Sprintf("%x", md5.Sum(bytesToHash)))

		interesting := true

		for i := 0; i < 5; i++ {
			if hash[i] != zeroByte {
				interesting = false
				break
			}
		}

		if interesting {
			passChars = append(passChars, hash[5])
		}

		i++
	}

	fmt.Println(string(passChars))

}

func RunPartB(filename string) {
	doorid := LoadData(filename)

	spaceByte := []byte(" ")[0]
	filledChars := 0
	passChars := []byte{spaceByte, spaceByte, spaceByte, spaceByte, spaceByte, spaceByte, spaceByte, spaceByte}
	i := 0
	zeroByte := []byte("0")[0]

	for filledChars < 8 {
		bytesToHash := []byte(fmt.Sprintf("%s%v", doorid, i))
		hash := []byte(fmt.Sprintf("%x", md5.Sum(bytesToHash)))

		interesting := true

		for i := 0; i < 5; i++ {
			if hash[i] != zeroByte {
				interesting = false
				break
			}
		}

		if interesting {
			index := int(hash[5]) - int(zeroByte)
			if index >= 0 && index <= 7 && passChars[index] == spaceByte {
				passChars[index] = hash[6]
				filledChars++
			}
		}

		i++
	}

	fmt.Println(string(passChars))
}
