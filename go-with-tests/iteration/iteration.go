package iteration

import "strings"

const repeatCount = 5

func Repeat0(character string) string {
	var result string
	for range 5 {
		result += character
	}
	return result
}

func Repeat1(character string) string {
	var result strings.Builder
	for range repeatCount {
		result.WriteString(character)
	}
	return result.String()
}

func Repeat(character string, count int) string {
	var result strings.Builder
	for range count {
		result.WriteString(character)
	}
	return result.String()
}
