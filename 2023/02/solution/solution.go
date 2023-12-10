package solution

import (
	"strconv"
	"strings"
)

type Game map[string]int

const Red = "red"
const Green = "green"
const Blue = "blue"

func (g Game) ExceedsMax(max int, colour string) bool {
	return g[colour] > max
}

func (g Game) Power() int {
	return g[Red] * g[Green] * g[Blue]
}

func SumPossibleGames(games []string, maxRed, maxGreen, maxBlue int) (sum int) {
	for i, rawGame := range games {
		game := parseGame(rawGame)

		if game.ExceedsMax(maxRed, Red) || game.ExceedsMax(maxGreen, Green) || game.ExceedsMax(maxBlue, Blue) {
			continue
		}

		sum += (i + 1)
	}

	return sum
}

func SumPowerGames(games []string) (sum int) {
	for _, rawGame := range games {
		sum += parseGame(rawGame).Power()
	}

	return sum
}

func parseGame(s string) Game {
	game := map[string]int{
		Red:   0,
		Green: 0,
		Blue:  0,
	}

	values := strings.Split(s, ":")

	for _, throw := range strings.Split(values[1], ";") {
		cubes := strings.Split(throw, ",")
		for _, cube := range cubes {
			v := strings.Split(strings.TrimPrefix(cube, " "), " ")
			value, _ := strconv.Atoi(v[0])

			if value > game[v[1]] {
				game[v[1]] = value
			}
		}
	}

	return game
}
