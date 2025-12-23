package main

func part1(numberMatrix [][]int, operators []string) int {
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
