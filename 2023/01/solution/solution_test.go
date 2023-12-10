package solution

import (
	"fmt"
	"strings"
	"testing"
)

func Test_SumLines(t *testing.T) {
	input := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	want := 142
	got := SumLines(input, nil)

	if want != got {
		t.Errorf("sumLines() = %v, want %v", got, want)
	}
}

func TestSumLines(t *testing.T) {
	type args struct {
		lines    []string
		replacer *strings.Replacer
	}

	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				lines: []string{
					"two1nine",
					"eightwothree",
					"abcone2threexyz",
					"xtwone3four",
					"4nineeightseven2",
					"zoneight234",
					"7pqrstsixteen",
				},
				replacer: Replacer,
			},
			want: 281,
		},
		{
			args: args{
				lines: []string{
					"oneight",
				},
				replacer: Replacer,
			},
			want: 18,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			if got := SumLines(tt.args.lines, tt.args.replacer); got != tt.want {
				t.Errorf("SumLinesRegex() = %v, want %v", got, tt.want)
			}
		})
	}
}
