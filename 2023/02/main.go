package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/MaikelVeen/adventofcode/02/solution"
)

func main() {
	lines, _ := os.ReadFile("input.txt")
	input := strings.Split(string(lines), "\n")

	fmt.Println(solution.SumPossibleGames(input, 12, 13, 14))
	fmt.Println(solution.SumPowerGames(input))
}
