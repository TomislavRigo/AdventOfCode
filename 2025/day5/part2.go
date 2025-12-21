package main

import (
	"strconv"
	"strings"
)

func part2(ranges []string) int {
	total := 0

	validItemRanges := []Range{}
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		min, _ := strconv.Atoi(parts[0])
		max, _ := strconv.Atoi(parts[1])

		validItemRanges = append(validItemRanges, Range{min, max})
	}

	for _, rang := range validItemRanges {
		total += rang.max - rang.min
	}

	return total
}
