package test2

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

const LOCATION_PATTERN = `(?:(\w+)\s+from\s+(\w+)\s+)?x=([+-]?\d+)\s*y=([+-]?\d+)`

type coordination struct {
	x int
	y int
}

type navigate struct {
	name string
	cord int
}

type cordNavigate struct {
	x navigate
	y navigate
}

var startCoordination = coordination{}
var coordinationList = make(map[string]coordination)

type pointToList struct {
	from string
	to   string
	cord cordNavigate
}

var points = []pointToList{}

func main() {
	re := regexp.MustCompile(LOCATION_PATTERN)

	var text string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text += scanner.Text() + "\n"
	}

	// Check for errors in scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	// Find all matches
	matches := re.FindAllStringSubmatch(text, -1)

	// Loop through the matches and print s, d, x, and y
	for _, match := range matches {
		x, _ := strconv.Atoi(match[3])
		y, _ := strconv.Atoi(match[4])

		if match[1] != "" && match[2] != "" {
			// Extract start x and y from matched regex value
			from := match[1]
			to := match[2]

			cord_x, _ := strconv.Atoi(string(match[3][1]))
			cord_y, _ := strconv.Atoi(string(match[4][1]))

			points = append(points, pointToList{
				from: from,
				to:   to,
				cord: cordNavigate{
					x: navigate{
						name: string(match[3][0]),
						cord: cord_x,
					},
					y: navigate{
						name: string(match[4][0]),
						cord: cord_y,
					},
				},
			})
		} else {
			// Set start location
			startCoordination = coordination{
				x: x,
				y: y,
			}
		}
	}

	printResult := func(from string, cx int, cy int) {
		fmt.Printf("%s x=%d y=%d\n", from, cx, cy)
	}

	calcCord := func(ds cordNavigate, st coordination) (f_x int, f_y int) {
		var cord_x int
		var cord_y int

		if ds.x.name == "-" {
			cord_x = st.x - ds.x.cord
		}

		if ds.x.name == "+" {
			cord_x = st.x + ds.x.cord
		}

		if ds.y.name == "-" {
			cord_y = st.y - ds.y.cord
		}

		if ds.y.name == "+" {
			cord_y = st.y + ds.y.cord
		}

		return cord_x, cord_y
	}

	var findCord = func(loc pointToList) {
		if loc.to == "start" {
			cord_x, cord_y := calcCord(loc.cord, startCoordination)

			coordinationList[loc.from] = coordination{
				x: cord_x,
				y: cord_y,
			}
		} else {
			to, ok := coordinationList[loc.to]

			if ok {
				cord_x, cord_y := calcCord(loc.cord, to)
				coordinationList[loc.from] = coordination{
					x: cord_x,
					y: cord_y,
				}
			} else {
				idx := 0
				for i, v := range points {
					if v.from == loc.to{
						idx = i
					}
				}

				cord_x, cord_y := calcCord(points[idx].cord, to)
				coordinationList[loc.from] = coordination{
					x: cord_x,
					y: cord_y,
				}
			}
		}
	}

	ReverseSlice(points)
	for _, loc := range points {
		findCord(loc)
	}

	keys := make([]string, 0, len(coordinationList))

	for k := range coordinationList {
		keys = append(keys, k)
	}

	ReverseSlice(keys)

	for _, k := range keys {
		cord := coordinationList[k]
		printResult(k, cord.x, cord.y)
	}
}


func ReverseSlice[T comparable](s []T) {
    sort.SliceStable(s, func(i, j int) bool {
        return i > j
    })
}