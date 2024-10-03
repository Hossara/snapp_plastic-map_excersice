package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Coordinates struct {
	x, y int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	locations := make(map[string]Coordinates)
	
	//order := []string{}

	// Reading the first line for the start coordinates
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Fields(line)
	startX, _ := strconv.Atoi(strings.Split(parts[1], "=")[1])
	startY, _ := strconv.Atoi(strings.Split(parts[2], "=")[1])
	locations["start"] = Coordinates{startX, startY}

	// Read the remaining lines for other locations
	for scanner.Scan() {
		line := scanner.Text()
		parts = strings.Fields(line)

		println(startX)
		println(startY)
		println(parts)

		// Extract the location name, reference location, and relative coordinates
		// locName := parts[0]
		// refLoc := parts[2]

		// // Handle x and y adjustments safely
		// dx, _ := strconv.Atoi(parts[4][2:]) // get x adjustment
		// if parts[4][1] == '-' {
		// 	dx = -dx
		// }
		// dy, _ := strconv.Atoi(parts[5][2:]) // get y adjustment
		// if parts[5][1] == '-' {
		// 	dy = -dy
		// }

		// // Calculate absolute coordinates
		// refCoords := locations[refLoc]
		// absX := refCoords.x + dx
		// absY := refCoords.y + dy

		// // Store the coordinates and maintain the order for output
		// locations[locName] = Coordinates{absX, absY}
		// order = append(order, locName)
	}

	// Print the results in the order they appeared in input
	// for _, loc := range order {
	// 	coords := locations[loc]
	// 	fmt.Printf("%s x=%d y=%d\n", loc, coords.x, coords.y)
	// }
}
