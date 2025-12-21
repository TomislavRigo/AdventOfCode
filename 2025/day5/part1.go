package main

import (
	"strconv"
	"strings"
)

func part1(ranges []string, items []string) int {
	count := 0

	validItemRanges := []Range{}
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		min, _ := strconv.Atoi(parts[0])
		max, _ := strconv.Atoi(parts[1])

		validItemRanges = append(validItemRanges, Range{min, max})
	}

	for _, item := range items {
		parsed, _ := strconv.Atoi(item)

		for _, r := range validItemRanges {
			if r.inRange(parsed) {
				count++
				break
			}
		}
	}

	return count
}
