package hw02unpackstring

import (
	"errors"
)

var ErrInvalidString = errors.New("invalid string")

const asciiZero = 48

func Unpack(s string) (res string, err error) {
	src := []rune(s)

	inCounter := false
	for pos := range src {
		if inCounter {
			inCounter = false
			continue
		}

		if isDigit(src[pos]) {
			err = ErrInvalidString
			break
		}

		if pos == len(src)-1 {
			res += string(src[pos])
			break
		}

		if isDigit(src[pos+1]) {
			if atoi(src[pos+1]) > 0 {
				res += repeat(src[pos], atoi(src[pos+1]))
			}
			inCounter = true
		} else {
			res += string(src[pos])
		}
	}
	return
}

func atoi(s rune) int {
	return int(s) - asciiZero
}

func isDigit(s rune) bool {
	return s >= asciiZero && s <= asciiZero+9
}

func repeat(s rune, n int) (res string) {
	for i := 0; i < n; i++ {
		res += string(s)
	}
	return
}
