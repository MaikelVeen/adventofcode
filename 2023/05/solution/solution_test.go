package solution

import (
	"os"
	"strings"
	"testing"
)

func TestFindDestination(t *testing.T) {
	type args struct {
		source   int
		mappings func() []Mapping
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testcase1",
			args: args{
				mappings: func() []Mapping {
					lines, _ := os.ReadFile("../input_test.txt")
					input := strings.Split(string(lines), "\n")
					return ParseMappings(input[3:5])
				},
				source: 79,
			},
			want: 81,
		},
		{
			name: "testcase1",
			args: args{
				mappings: func() []Mapping {
					lines, _ := os.ReadFile("../input_test.txt")
					input := strings.Split(string(lines), "\n")
					return ParseMappings(input[3:5])
				},
				source: 55,
			},
			want: 57,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDestination(tt.args.source, tt.args.mappings()); got != tt.want {
				t.Errorf("FindDestination() = %v, want %v", got, tt.want)
			}
		})
	}
}
