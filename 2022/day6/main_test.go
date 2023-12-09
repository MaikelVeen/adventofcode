package main

import (
	"fmt"
	"testing"
)

func Test_slidingWindow(t *testing.T) {
	tests := map[string]struct {
		s    string
		size int
		want int
	}{
		"test case 1": {
			s:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			size: 4,
			want: 5,
		},
		"test case 2": {
			s:    "nppdvjthqldpwncqszvftbrmjlhg",
			size: 4,
			want: 6,
		},
		"test case 3": {
			s:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			size: 4,
			want: 10,
		},
		"test case 4": {
			s:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			size: 4,
			want: 11,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := slidingWindow(tt.s, tt.size); got != tt.want {
				t.Errorf("slidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_slidingWindowPart2(t *testing.T) {
	tests := []struct {
		s    string
		size int
		want int
	}{
		{
			s:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			size: 14,
			want: 19,
		},
		{
			s:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			size: 14,
			want: 23,
		},
		{
			s:    "nppdvjthqldpwncqszvftbrmjlhg",
			size: 14,
			want: 23,
		},
		{
			s:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			size: 14,
			want: 29,
		},
		{
			s:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			size: 14,
			want: 26,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			if got := slidingWindow(tt.s, tt.size); got != tt.want {
				t.Errorf("slidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
