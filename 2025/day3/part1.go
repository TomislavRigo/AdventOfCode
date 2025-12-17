package main

import "strconv"

func part1(lines []string) int {
	total := 0

	for _, line := range lines {
		first := 0
		firstIndex := 0

		second := 0

		for i := 0; i < len(line)-1; i++ {
			parsed, _ := strconv.Atoi(string(line[i]))
			if first < parsed {
				first = parsed
				firstIndex = i
			}

		}

		for i := firstIndex + 1; i < len(line); i++ {
			parsed, _ := strconv.Atoi(string(line[i]))
			if second < parsed {
				second = parsed
			}
		}

		total += ((first * 10) + second)
	}

	return total
}
