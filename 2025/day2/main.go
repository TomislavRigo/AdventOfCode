package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	input := strings.Split(strings.TrimSpace(string(data)), ",")

	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}
