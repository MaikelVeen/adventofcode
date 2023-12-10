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

type Symbols map[image.Point]rune

func (g Symbols) TileAdjacentToSymbol(point image.Point) bool {
	for _, delta := range []image.Point{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
	} {

		if _, ok := g[point.Add(delta)]; ok {
			return ok
		}
	}
	return false
}

func SumPartNumbers(schematicLines []string) (sum int) {
	grid := parseSymbols(schematicLines)
	parts := map[image.Point]int{}

	for y, row := range schematicLines {
		for _, match := range regexp.MustCompile(`(?m)[0-9]+`).FindAllStringIndex(row, -1) {
			startPos, endPos := match[0], match[1]
			value, _ := strconv.Atoi(row[startPos:endPos])
			added := false

			for x := startPos; x < endPos; x++ {
				point := image.Point{x, y}
				parts[point] = value

				if grid.TileAdjacentToSymbol(point) {
					if !added {
						sum += value
						added = true
					}
					break
				}
			}
		}
	}
	return sum
}

func parseSymbols(schematicLines []string) Symbols {
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
