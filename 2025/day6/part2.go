package main

import (
	"strconv"
	"strings"
)

type accFunction = func(accumulator int, element int) int

func reduce(slice []int, fn accFunction, initial int) int {
	result := initial

	for _, element := range slice {
		result = fn(result, element)
	}

	return result
}

func calculate(numberMatrix []string, operator string) int {
	total := 0

	var numbers []int
	for i := len(operator) - 1; i >= 0; i-- {
		number := ""
		for _, element := range numberMatrix {
			if element[i] != ' ' {
				number += string(element[i])
			}
		}

		parsed, _ := strconv.Atoi(number)
		numbers = append(numbers, parsed)
	}

	switch strings.TrimSpace(operator) {
	case "+":
		total += reduce(numbers, func(acc, elem int) int { return acc + elem }, 0)
	case "*":
		total += reduce(numbers, func(acc, elem int) int { return acc * elem }, 1)
	}

	return total
}

func part2(data []byte) int {
	lines := strings.Split(string(data), "\n")

	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	lastElementIndex := len(lines) - 1
	numbers := lines[:lastElementIndex]
	operators := lines[lastElementIndex]

	total := 0

	leftCursor := 0
	rightCursor := 0
	for leftCursor < len(operators) {
		rightCursor++
		for rightCursor != len(operators) && operators[rightCursor] == ' ' {
			rightCursor++
		}

		// not sure why
		if rightCursor != len(operators)-2 {
			rightCursor--
		}

		var numberMatrix []string
		for _, numberLine := range numbers {
			numberMatrix = append(numberMatrix, string(numberLine[leftCursor:rightCursor]))
		}

		total += calculate(numberMatrix, operators[leftCursor:rightCursor])
		rightCursor++
		leftCursor = rightCursor
	}

	return total
}
