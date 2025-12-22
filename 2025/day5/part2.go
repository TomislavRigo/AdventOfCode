package main

import (
	"strconv"
	"strings"
)

func mergeRange(rng Range, mergedRanges []Range) []Range {
	currentRange := rng
	for _, mergedRange := range mergedRanges {
		if currentRange.min >= mergedRange.min && currentRange.max <= mergedRange.max {
			return mergedRanges
		}

		if currentRange.min < mergedRange.min && currentRange.max >= mergedRange.min && currentRange.max <= mergedRange.max {
			currentRange = Range{currentRange.min, mergedRange.min - 1}
			continue
		}

		if currentRange.max > mergedRange.max && currentRange.min >= mergedRange.min && currentRange.min <= mergedRange.max {
			currentRange = Range{mergedRange.max + 1, currentRange.max}
			continue
		}

		if currentRange.min < mergedRange.min && currentRange.max > mergedRange.max {
			mergedRanges = mergeRange(Range{currentRange.min, mergedRange.min - 1}, mergedRanges)
			mergedRanges = mergeRange(Range{mergedRange.max + 1, currentRange.max}, mergedRanges)

			return mergedRanges
		}
	}

	return append(mergedRanges, currentRange)
}

func part2(input []string) int {
	count := 0

	ranges := []Range{}
	for _, r := range input {
		parts := strings.Split(r, "-")
		min, _ := strconv.Atoi(parts[0])
		max, _ := strconv.Atoi(parts[1])

		ranges = append(ranges, Range{min, max})
	}

	mergedRanges := []Range{}
	for _, rng := range ranges {
		mergedRanges = mergeRange(rng, mergedRanges)
	}

	for _, rng := range mergedRanges {
		count += rng.max - rng.min + 1
	}

	return count
}
