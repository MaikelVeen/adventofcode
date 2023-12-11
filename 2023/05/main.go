package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/MaikelVeen/adventofcode/05/solution"
)

func main() {
	lines, _ := os.ReadFile("input.txt")
	input := strings.Split(string(lines), "\n")
	fmt.Println(solution.LowestLocation(input, []solution.SourceLocation{
		{
			LineStart: 3,
			LineEnd:   40,
		},
		{
			LineStart: 42,
			LineEnd:   61,
		},
		{
			LineStart: 63,
			LineEnd:   72,
		},
		{
			LineStart: 74,
			LineEnd:   100,
		},
		{
			LineStart: 102,
			LineEnd:   148,
		},
		{
			LineStart: 150,
			LineEnd:   175,
		},
		{
			LineStart: 177,
			LineEnd:   197,
		},
	}))
}
