package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var b strings.Builder
	var buffered string
	var escaped = false

	for _, currentRune := range s {
		currentChar := string(currentRune)

		switch {
		case isEscapeSymbol(currentRune) && !escaped:
			b.WriteString(buffered)
			buffered = ""
			escaped = true

		case isEscapeSymbol(currentRune) && escaped:
			b.WriteString(buffered)
			buffered = currentChar
			escaped = false

		case unicode.IsDigit(currentRune) && !escaped:
			if buffered == "" {
				return "", ErrInvalidString
			}
			n, err := strconv.Atoi(currentChar)
			if err != nil {
				return "", err
			}
			b.WriteString(strings.Repeat(buffered, n))
			buffered = ""

		case unicode.IsDigit(currentRune) && escaped:
			b.WriteString(buffered)
			buffered = currentChar
			escaped = false

		default:
			b.WriteString(buffered)
			buffered = currentChar
			if escaped {
				escaped = false
			}
		}
	}

	if escaped {
		return "", ErrInvalidString
	}

	b.WriteString(buffered)

	return b.String(), nil
}

func isEscapeSymbol(r rune) bool {
	return string(r) == `\`
}
