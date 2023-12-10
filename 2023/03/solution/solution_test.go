package solution

import (
	"os"
	"strings"
	"testing"
)

func TestSumPartNumbers(t *testing.T) {
	type args struct {
		schematicLines func() []string
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{
			name: "Test input",
			args: args{
				schematicLines: func() []string {
					lines, _ := os.ReadFile("../input_test.txt")
					return strings.Split(string(lines), "\n")
				},
			},
			wantSum: 4361,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := SumPartNumbers(tt.args.schematicLines()); gotSum != tt.wantSum {
				t.Errorf("SumPartNumbers() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
