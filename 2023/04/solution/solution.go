package solution

import (
	"math"
	"regexp"
	"slices"
	"strings"
)

func SumPoints(input []string) (sum int) {
	for _, i := range input {
		line := strings.Split(strings.Split(i, ":")[1], "|")

		winningNumbers := parseLine(line[0])
		playingNumbers := parseLine(line[1])

		hits := float64(-1)
		for _, num := range playingNumbers {
			if slices.Contains(winningNumbers, num) {
				hits += 1
			}
		}

		sum += int(math.Pow(2, hits))
	}

	return sum
}

func SumCards(input []string) (sum int) {
	copies := make([]int, len(input))

	for i, s := range input {
		copies[i]++

		line := strings.Split(strings.Split(s, ":")[1], "|")

		winningNumbers := parseLine(line[0])
		playingNumbers := parseLine(line[1])

		hits := 0
		for _, num := range playingNumbers {
			if slices.Contains(winningNumbers, num) {
				hits += 1
				copies[i+hits] += copies[i]
			}
		}

		sum += copies[i]
	}

	return sum
}

func parseLine(s string) []string {
	re := regexp.MustCompile(`\d+`)
	return re.FindAllString(s, -1)
}
