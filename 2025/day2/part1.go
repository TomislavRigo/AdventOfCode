package main

import (
	"strconv"
	"strings"
)

func part1(input []string) int {
	total := 0

	for _, line := range input {
		parts := strings.Split(line, "-")
		min, _ := strconv.Atoi(parts[0])
		max, _ := strconv.Atoi(parts[1])

		current := min

		for current <= max {
			parsed := strconv.Itoa(current)
			length := len(parsed)

			if length%2 != 0 {
				current++
				continue
			}

			half := length / 2

			if parsed[:half] == parsed[half:] {
				total += current
			}

			current++
		}
	}

	return total
}
