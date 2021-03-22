package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

// Unpack осуществляет примитивную распаковку строки, содержащую повторяющиеся символы/руны.
func Unpack(s string) (string, error) {
	var b strings.Builder

	if s == "" {
		return s, nil
	}

	// Строка не должна начинаться с цифры
	if unicode.IsDigit(rune(s[0])) {
		return "", ErrInvalidString
	}

	var prevRune rune
	var initPrevRune bool

	for _, r := range s {
		if unicode.IsDigit(r) {
			if unicode.IsDigit(prevRune) {
				return "", ErrInvalidString
			}
			c, _ := strconv.Atoi(string(r))
			b.WriteString(strings.Repeat(string(prevRune), c))
		}

		if initPrevRune && !unicode.IsDigit(prevRune) && !unicode.IsDigit(r) {
			b.WriteString(string(prevRune))
		}

		prevRune = r
		initPrevRune = true
	}

	if !unicode.IsDigit(prevRune) {
		b.WriteString(string(prevRune))
	}
	return b.String(), nil
}
