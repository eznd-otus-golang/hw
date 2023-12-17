package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

const asciiZero = 48

func Unpack(s string) (string, error) {
	src := []rune(s)
	b := strings.Builder{}

	inCounter := false
	for pos := range src {
		if inCounter {
			inCounter = false
			continue
		}

		if isDigit(src[pos]) {
			return "", ErrInvalidString
		}

		if pos == len(src)-1 {
			_, _ = b.WriteRune(src[pos])
			break
		}

		if isDigit(src[pos+1]) {
			if atoi(src[pos+1]) > 0 {
				_, _ = b.WriteString(strings.Repeat(string(src[pos]), atoi(src[pos+1])))
			}
			inCounter = true
			continue
		}

		b.WriteRune(src[pos])
	}
	return b.String(), nil
}

func atoi(s rune) int {
	return int(s) - asciiZero
}

func isDigit(s rune) bool {
	return s >= asciiZero && s <= asciiZero+9
}
