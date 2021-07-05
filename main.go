package main

func isRobotBounded(instructions string) bool {
	// north 0, east 1, south 2 west 3
	// x, y
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	// initial position
	x := 0
	y := 0

	// facing direction is north
	index := 0

	for _, v := range instructions {
		if v == 'L' {
			// Modulo here is needed to deal with the situation - facing west, idx = 3, turn to the right to face north, idx = 0
			// west
			index = (index + 3) % 4
		} else if v == 'R' {
			// east
			index = (index + 1) % 4
		} else {
			x += directions[index][0]
			y += directions[index][1]
		}
	}
	return (x == 0 && y == 0) || index != 0
}
