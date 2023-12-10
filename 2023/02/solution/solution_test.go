package solution

import (
	"os"
	"strings"
	"testing"
)

func TestSumPossibleGames(t *testing.T) {
	type args struct {
		games    func() []string
		maxRed   int
		maxGreen int
		maxBlue  int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test input",
			args{
				func() []string {
					lines, _ := os.ReadFile("../input_test.txt")
					return strings.Split(string(lines), "\n")
				},
				12, 13, 14,
			},
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumPossibleGames(tt.args.games(), tt.args.maxRed, tt.args.maxGreen, tt.args.maxBlue); got != tt.want {
				t.Errorf("SumPossibleGames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumPower(t *testing.T) {
	type args struct {
		games func() []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test input",
			args{
				func() []string {
					lines, _ := os.ReadFile("../input_test.txt")
					return strings.Split(string(lines), "\n")
				},
			},
			2286,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumPowerGames(tt.args.games()); got != tt.want {
				t.Errorf("SumPossibleGames() = %v, want %v", got, tt.want)
			}
		})
	}
}
