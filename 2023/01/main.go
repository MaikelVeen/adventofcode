package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/MaikelVeen/adventofcode/01/solution"
)

func main() {
	lines, _ := os.ReadFile("input.txt")
	input := strings.Fields(string(lines))

	fmt.Println(solution.SumLines(input, nil))
	fmt.Println(solution.SumLines(input, solution.Replacer))
}
