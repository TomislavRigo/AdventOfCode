package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	position = 50
	count    = 0
)

func add(amount int) {
	for range amount {
		position++
		if position == 100 {
			position = 0
		}
	}

	if position == 0 {
		count++
	}
}

func subtract(amount int) {
	for range amount {
		position--
		if position == -1 {
			position = 99
		}
	}

	if position == 0 {
		count++
	}
}

func main() {
	data, err := os.ReadFile("../data.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.SplitSeq(strings.TrimSpace(string(data)), "\n")

	for line := range lines {
		indicator := line[:1]
		amount, _ := strconv.Atoi(line[1:])

		if indicator == "R" {
			add(amount)
		} else {
			subtract(amount)
		}
	}

	fmt.Println(count)
}
