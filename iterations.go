package code

import "strings"

func Repeat(input string, repeatCount int) string {
	var sb strings.Builder
	for i := 0; i < repeatCount; i++ {
		sb.WriteString(input)
	}

	return sb.String()
}
