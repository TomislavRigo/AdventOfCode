package main

import (
	"regexp"
	"strconv"
	"strings"
)

func part1(data []byte) int {
	re := regexp.MustCompile(` +`)
	normalizedData := re.ReplaceAllString(string(data), " ")
	lines := strings.Split(strings.TrimSpace(normalizedData), "\n")

	lastElementIndex := len(lines) - 1
	numberLines := lines[:lastElementIndex]
	operators := strings.Split(lines[lastElementIndex], " ")

	var numberMatrix [][]int
	for _, numberLine := range numberLines {
		numbers := strings.Split(strings.TrimSpace(numberLine), " ")
		var newElement []int
		for _, number := range numbers {
			parsed, _ := strconv.Atoi(number)
			newElement = append(newElement, parsed)
		}

		numberMatrix = append(numberMatrix, newElement)

	}

	total := 0

	for x := range numberMatrix[0] {
		result := 0
		for y := range numberMatrix {
			switch operators[x] {
			case "+":
				result += numberMatrix[y][x]
			case "*":
				if result == 0 {
					result = numberMatrix[y][x]
				} else {
					result *= numberMatrix[y][x]
				}
			}
		}

		total += result
	}

	return total
}
