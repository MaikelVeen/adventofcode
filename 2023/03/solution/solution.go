package solution

import (
	"image"
	"regexp"
	"strconv"
	"unicode"
)

type Tile uint

const Void Tile = 0
const Digit Tile = 1
const Symbol Tile = 2

type Grid map[image.Point]rune

func (g Grid) TileAdjacentToSymbol(y, x int) bool {
	for _, delta := range []image.Point{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
	} {

		if _, ok := g[image.Point{x, y}.Add(delta)]; ok {
			return ok
		}
	}
	return false
}

func SumPartNumbers(schematicLines []string) (sum int) {
	grid := parseGrid(schematicLines)
	regex := regexp.MustCompile(`(?m)[0-9]+`)

	for y, row := range schematicLines {
		for _, match := range regex.FindAllStringIndex(row, -1) {
			startPos := match[0]
			endPos := match[1]

			for x := startPos; x < endPos; x++ {
				if grid.TileAdjacentToSymbol(y, x) {
					value, _ := strconv.Atoi(row[startPos:endPos])
					sum += value
					break
				}
			}
		}
	}
	return sum
}

func parseGrid(schematicLines []string) Grid {
	grid := map[image.Point]rune{}

	for y, line := range schematicLines {
		for x, char := range line {
			if char == '.' || unicode.IsDigit(char) {
				continue
			}
			grid[image.Point{x, y}] = char
		}
	}
	return grid
}
