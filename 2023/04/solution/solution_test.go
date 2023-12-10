package solution

import (
	"os"
	"strings"
	"testing"
)

func TestSumPoints(t *testing.T) {
	type args struct {
		input func() []string
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{
			"test case",
			args{
				input: func() []string {
					lines, _ := os.ReadFile("../input_test.txt")
					return strings.Split(string(lines), "\n")
				},
			},
			13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := SumPoints(tt.args.input()); gotSum != tt.wantSum {
				t.Errorf("SumPoints() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func TestSumCards(t *testing.T) {
	type args struct {
		input func() []string
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{
			"test case",
			args{
				input: func() []string {
					lines, _ := os.ReadFile("../input_test.txt")
					return strings.Split(string(lines), "\n")
				},
			},
			30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := SumCards(tt.args.input()); gotSum != tt.wantSum {
				t.Errorf("SumCards() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
