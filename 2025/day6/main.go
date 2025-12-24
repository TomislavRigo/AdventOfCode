package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	// fmt.Println(numberMatrix)
	// fmt.Println(operators)

	fmt.Println("Part 1: ", part1(slices.Clone(data)))
	fmt.Println("Part 2: ", part2(slices.Clone(data)))
}
