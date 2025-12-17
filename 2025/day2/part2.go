package main

import (
	"strconv"
	"strings"
)

func part2(lines []string) int {
	total := 0

	for _, line := range lines {
		parts := strings.Split(line, "-")
		min, _ := strconv.Atoi(parts[0])
		max, _ := strconv.Atoi(parts[1])

		current := min
		for current <= max {
			parsed := strconv.Itoa(current)
			sep := ""

			for index := range len(parsed) {
				sep += string(parsed[index])

				if len(parsed)%len(sep) != 0 || len(parsed) == len(sep) {
					continue
				}

				match := true
				for i := len(sep); i < len(parsed); i += len(sep) {
					if parsed[i:i+len(sep)] != sep {
						match = false
						break
					}
				}

				if match {
					total += current
					break
				}
			}

			current++
		}
	}

	return total
}
