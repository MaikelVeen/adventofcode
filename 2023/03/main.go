package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/MaikelVeen/adventofcode/03/solution"
)

func main() {
	lines, _ := os.ReadFile("input.txt")
	input := strings.Split(string(lines), "\n")

	fmt.Println(solution.SumPartNumbers(input))
}
