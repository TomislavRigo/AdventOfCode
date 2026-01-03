package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	Source     = 'S'
	Splitter   = '^'
	EmptySpace = '.'
)

func contains(arr []int, item int) bool {
	for _, element := range arr {
		if element == item {
			return true
		}
	}
	return false
}

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Println("Part 1: ", part1(lines))
	fmt.Println("Part 2: ", part2(lines))
}
