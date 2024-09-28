package main

import "fmt"

// Location types
type coordination struct {
	x int
	y int
}

type locations map[string]coordination

var startLocation coordination

func main() {
	locationList := make(locations)

	// Main loop
	for {
		// Define variables
		var (
			name string
			x    int
			y    int
		)

		// Read line by line
		args, err := fmt.Scanf("%s %d %d", &name, &x, &y)

		// End of line
		if args < 3 {
			break
		}

		// Handle wrong format
		if err != nil {
			panic("name x y format is not correct")
		}

		// Handle list limitations
		if len(locationList) == 1000 {
			panic("you reach maximum number of coordinations")
		}

		// Handle x, y limitations
		if x <= -1000000 || x >= 1000000 || y <= -1000000 || y >= 1000000 {
			panic("x and y number must be between -1000000 and 1000000")
		}

		// Handle first element rule
		if (coordination{}) == startLocation {
			// Set start coordination
			startLocation = coordination{
				x: x,
				y: y,
			}

			if name != "start" {
				panic("first element should be start")
			}

			continue
		}

		println("Im not here", name)

		// Push to map
		locationList[name] = coordination{
			x: x,
			y: y,
		}
	}
}
