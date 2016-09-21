package util

import (
	"strings"
	"unicode/utf8"
)

// MaskString keep first `showNum`` count char of a string `str`` and change all remaining chars to "*"
func MaskString(str string, showNum int) string {
	var res string
	var i int
	var c rune
	for i, c = range str {
		if i < showNum {
			res += string(c)
		} else {
			break
		}
	}
	if i != showNum {
		i++
	}
	length := utf8.RuneCountInString(str)
	if i < length {
		res += strings.Repeat("*", length-i)
	}
	return res
}
