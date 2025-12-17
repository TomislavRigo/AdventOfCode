package main

import (
	"strconv"
)

func checkOffset(value string, currentIndex int, offset int) bool {
	if currentIndex+offset >= len(value) {
		return true
	}

	for i := 1; i <= offset; i++ {
		parsedCurrent, _ := strconv.Atoi(string(value[currentIndex]))
		parsedOnOffset, _ := strconv.Atoi(string(value[currentIndex+i]))

		if parsedCurrent < parsedOnOffset {
			return true
		}
	}

	return false
}

func part2(lines []string) int {
	total := 0

	for _, line := range lines {

		result := line
		index := 0
		for len(result) > 12 {
			offset := len(result) - 12

			shouldRemove := checkOffset(result, index, offset)
			if shouldRemove {
				result = result[:index] + result[index+1:]
				continue
			}

			index++
		}

		parsed, _ := strconv.Atoi(result)
		total += parsed
	}

	return total
}
