package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type coordination struct {
	x int
	y int
}

type location struct {
	name  string
	from  string
	index int
	cord  coordination
}

func parseCoordinate(coordination string) int {

	if len(coordination) < 2 {
		return 0
	}

	valueStr := coordination[2:]

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0
	}
	return value
}

func main() {
	var (
		locations         []location
		loc               location
		lines             []string
		i                 int
		startCoordination coordination
	)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		lines = strings.Fields(scanner.Text())

		if i == 0 {
			x := parseCoordinate(lines[1])
			y := parseCoordinate(lines[2])

			if x <= -1000000 || x >= 1000000 || y <= -1000000 || y >= 1000000 {
				panic("x and y number must be between -1000000 and 1000000")
			}

			startCoordination = coordination{
				x: x,
				y: y,
			}
		} else {

			x := parseCoordinate(lines[3])
			y := parseCoordinate(lines[4])

			if x <= -1000000 || x >= 1000000 || y <= -1000000 || y >= 1000000 {
				//panic("x and y number must be between -1000000 and 1000000")
				return
			}

			loc = location{
				name: lines[0],
				from: lines[2],
				cord: coordination{
					x: x,
					y: y,
				},
				index: i,
			}
			locations = append(locations, loc)
		}

		i++
	}

	// Handle list limitations
	if len(locations) == 1000 {
		return
		//panic("you reach maximum number of coordinations")
	}

	sort.Slice(locations, func(a, b int) bool {
		if locations[a].from == "start" && locations[b].from != "start" {
			return true
		} else {
			return false
		}
	})

	for i := range locations {
		if locations[i].from == "start" {
			locations[i].cord.x += startCoordination.x
			locations[i].cord.y += startCoordination.y
		} else {
			for j := 0; j < i; j++ {
				if locations[j].name == locations[i].from {
					locations[i].cord.x = locations[j].cord.x + locations[i].cord.x
					locations[i].cord.y = locations[j].cord.y + locations[i].cord.y
					break
				}
			}
		}
	}

	sort.Slice(locations, func(a, b int) bool {
		return locations[a].index < locations[b].index
	})

	for _, p := range locations {
		fmt.Printf("%s x=%d y=%d\n", p.name, p.cord.x, p.cord.y)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
