package main

func part1(lines []string) int {
	count := 0

	var beamLocations []int

	// Initialize
	for index, char := range lines[0] {
		if char == Source {
			beamLocations = []int{index}
			break
		}
	}

	// Chech beam location until last row
	for _, line := range lines[1 : len(lines)-1] {
		newLocations := []int{}

		for _, beamLocation := range beamLocations {

			// Check if space already filled
			if contains(newLocations, beamLocation) {
				continue
			}

			locationType := line[beamLocation]

			switch locationType {
			case byte(EmptySpace):
				if !contains(newLocations, beamLocation) {
					newLocations = append(newLocations, beamLocation)
				}
			case byte(Splitter):
				count += 1
				previous := beamLocation - 1
				if previous >= 0 && line[previous] == byte(EmptySpace) && !contains(newLocations, beamLocation-1) {
					newLocations = append(newLocations, beamLocation-1)
				}
				next := beamLocation + 1
				if next < len(line) && line[next] == byte(EmptySpace) && !contains(newLocations, beamLocation+1) {
					newLocations = append(newLocations, beamLocation+1)
				}
			default:
				panic("Unexpected location type")
			}

		}

		beamLocations = newLocations
	}

	return count
}
