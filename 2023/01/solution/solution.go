package solution

import (
	"strconv"
	"strings"
)

var Replacer = strings.NewReplacer("one", "o1e", "two", "t2o", "three", "t3e", "four",
	"f4r", "five", "f5e", "six", "s6x", "seven", "s7n", "eight", "e8t", "nine", "n9e")

func SumLines(lines []string, replacer *strings.Replacer) int {
	sum := 0
	for _, line := range lines {
		l := line
		if replacer != nil {
			l = replacer.Replace(replacer.Replace(l))
		}
		sum += parseLine([]rune(l))
	}

	return sum
}

func parseLine(runes []rune) int {
	first := 0
	last := 0

	for _, r := range runes {
		if i, err := strconv.Atoi(string(r)); err == nil {
			first = i
			break
		}
	}

	for i := len(runes) - 1; i >= 0; i-- {
		if i, err := strconv.Atoi(string(runes[i])); err == nil {
			last = i
			break
		}
	}

	return (first * 10) + last
}
