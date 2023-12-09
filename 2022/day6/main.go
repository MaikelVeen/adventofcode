package main

import (
	"fmt"
	"io/ioutil"
)

func readFile(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func slidingWindow(s string, size int) int {
	for i := 0; i <= len(s)-size; i++ {
		if unique(s[i : i+size]) {
			return i + size
		}
	}

	return -1
}

func unique(s string) bool {
	seen := make([]bool, 26)

	for _, r := range s {
		if seen[r-'a'] {
			return false
		}
		seen[r-'a'] = true
	}

	return true
}

func main() {
	content, err := readFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Start of packet marker -> 4 unique chars.
	res := slidingWindow(content, 4)
	fmt.Println(res)

	// Start of message marker -> 14 unique chars.
	res2 := slidingWindow(content, 14)
	fmt.Println(res2)
}
