package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

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

	// fmt.Println(numberMatrix)
	// fmt.Println(operators)

	fmt.Println("Part 1: ", part1(numberMatrix, operators))
	fmt.Println("Part 2: ", part2(numberMatrix, operators))
}
