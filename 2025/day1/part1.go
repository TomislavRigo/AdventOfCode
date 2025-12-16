package main

import (
	"strconv"
)

func part1Add(amount int, position int, count int) (int, int) {
	for range amount {
		position++
		if position == 100 {
			position = 0
		}
	}

	if position == 0 {
		count++
	}

	return position, count
}

func part1Subtract(amount int, position int, count int) (int, int) {
	for range amount {
		position--
		if position == -1 {
			position = 99
		}
	}

	if position == 0 {
		count++
	}

	return position, count
}

func part1(lines []string) int {
	position := 50
	count := 0

	for _, line := range lines {
		indicator := line[:1]
		amount, _ := strconv.Atoi(line[1:])

		if indicator == "R" {
			position, count = part1Add(amount, position, count)
		} else {
			position, count = part1Subtract(amount, position, count)
		}
	}

	return count
}
