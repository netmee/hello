package iteration

import "strings"

// Repeat 5 time input string
func Repeat(character string, n int) (result string) {
	result = strings.Repeat(character, n)
	return
}
