package main

func checkPosition2(coordinate Coordinates, lines []string, xBound int, yBound int) bool {
	count := 0

	for xoffset := -1; xoffset <= 1; xoffset++ {
		currentX := coordinate.x + xoffset
		if currentX == -1 || currentX > xBound {
			continue
		}

		for yoffset := -1; yoffset <= 1; yoffset++ {
			currentY := coordinate.y + yoffset
			if currentY == -1 || currentY > yBound || (currentX == coordinate.x && currentY == coordinate.y) {
				continue
			}

			if lines[currentX][currentY] == '@' {
				count++
			}
		}
	}

	return count < 4
}

func part2(lines []string) int {
	count := 0

	yBound := len(lines) - 1
	xBound := len(lines[0]) - 1

	for x := 0; x < len(lines); x++ {
		for y := 0; y < len(lines[0]); y++ {
			if lines[x][y] == '@' {
				if checkPosition2(Coordinates{x, y}, lines, xBound, yBound) {
					count++
					lines[x] = lines[x][:y] + "." + lines[x][y+1:]
					x = -1
					break
				}
			}
		}
	}

	return count
}
