package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/MaikelVeen/adventofcode/04/solution"
)

func main() {
	lines, _ := os.ReadFile("input.txt")
	input := strings.Split(string(lines), "\n")

	//fmt.Println(solution.SumPoints(input))
	fmt.Println(solution.SumCards(input))
}
