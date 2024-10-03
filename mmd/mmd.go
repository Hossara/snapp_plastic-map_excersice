package mmd

import (
	"errors"
	"fmt"
)

// Location types
type coordination struct {
	x int
	y int
}

type location struct {
	from string
	to   string
	coordination
}

type locations map[string]location

var locationList = make(locations)
var coordinationList = make(map[string]coordination)

var startLocation coordination

func findElement(pointer string) (location, error) {
	for _, loc := range locationList {
		if loc.from == pointer {
			return loc, nil
		}
	}

	return location{}, errors.New("no more location")
}

func mmd() {
	// Main loop
	for {
		// Define variables
		var (
			from string
			to   string
			x    string
			y    string
		)

		if (coordination{}) == startLocation {
			_, err := fmt.Scanf("%s x=%s y=%s", &from, &x, &y)

			if err != nil {
				panic("start x y format is not correct")
			}
		} else {
			args, err := fmt.Scanf("%s from %s x=%s y=%s", &to, &from, &x, &y)

			// End of line
			if args < 4 {
				break
			}

			// Handle wrong format
			if err != nil {
				panic("start from end x=x y=y format is not correct")
			}
		}

		// Handle list limitations
		if len(locationList) == 1000 {
			panic("you reach maximum number of coordinations")
		}

		// Handle x, y limitations
		// if x <= -1000000 || x >= 1000000 || y <= -1000000 || y >= 1000000 {
		// 	panic("x and y number must be between -1000000 and 1000000")
		// }

		// Handle first element rule
		if (coordination{}) == startLocation {
			// Set start coordination
			// startLocation = coordination{
			// 	x: x,
			// 	y: y,
			// }

			if from != "start" {
				panic("first element should be start")
			}

			continue
		}

		println(x)
		println(y)

		// Push to map
		// locationList[from] = location{
		// 	from: from,
		// 	to:   to,
		// 	coordination: coordination{
		// 		x: x,
		// 		y: y,
		// 	},
		// }
	}

	// findLocation := func(cord location) {
	// 	it := coordination{}

	// 	if cord.from == "start" {
	// 		it.x = cord.x + startLocation.x
	// 		it.y = cord.y + startLocation.y
	// 	} else {
	// 		next, err := findElement(cord.to)

	// 		if err != nil {
	// 			return
	// 		}

	// 		it.x = cord.x + next.x
	// 		it.y = cord.y + next.y

	// 		fmt.Println()
	// 	}

	// 	fmt.Printf("%s from %s x=")
	// }

	// println(startLocation.x)
	// println(startLocation.y)

	// for _, cord := range locationList {

	// 	println(cord.coordination.x)
	// 	println(cord.coordination.y)
	// 	println(cord.from)
	// 	println(cord.to)
	// 	println(cord.x)
	// 	println(cord.y)
	// 	//findLocation(cord)
	// }
}
