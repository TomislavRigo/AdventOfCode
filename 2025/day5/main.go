package main

import (
	"fmt"
	"os"
	"strings"
)

type Range struct {
	min int
	max int
}

func (r Range) inRange(item int) bool {
	return r.min <= item && item <= r.max
}

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	ranges := []string{}
	items := []string{}
	for index := range lines {
		if lines[index] == "" {
			ranges = lines[:index]
			items = lines[index+1:]
		}
	}

	fmt.Println("Part 1: ", part1(ranges, items))
	fmt.Println("Part 2: ", part2(ranges))
}
