package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x, y int
}

type adjacency []coordinate

func adjacencies(height, width int) (result []adjacency) {
	// verticals
	for x := 0; x < width; x++ {
		for y := 0; y < height-3; y++ {
			result = append(result, adjacency{
				coordinate{x, y},
				coordinate{x, y + 1},
				coordinate{x, y + 2},
				coordinate{x, y + 3},
			})
		}
	}

	// horizontals
	for x := 0; x < width-3; x++ {
		for y := 0; y < height; y++ {
			result = append(result, adjacency{
				coordinate{x, y},
				coordinate{x + 1, y},
				coordinate{x + 2, y},
				coordinate{x + 3, y},
			})
		}
	}

	// diagonals
	for x := 0; x < width-3; x++ {
		for y := 0; y < height; y++ {
			// go up
			if y > 2 {
				result = append(result, adjacency{
					coordinate{x, y},
					coordinate{x + 1, y - 1},
					coordinate{x + 2, y - 2},
					coordinate{x + 3, y - 3},
				})
			}
			// go down
			if y < height-3 {
				result = append(result, adjacency{
					coordinate{x, y},
					coordinate{x + 1, y + 1},
					coordinate{x + 2, y + 2},
					coordinate{x + 3, y + 3},
				})
			}
		}
	}

	return result
}

func main() {
	file, _ := os.Open("input.txt")

	var grid [][]int

	s := bufio.NewScanner(file)
	for s.Scan() {
		var line []int
		for _, f := range strings.Fields(s.Text()) {
			i, _ := strconv.ParseInt(f, 10, 32)
			line = append(line, int(i))
		}
		grid = append(grid, line)
	}

	max := 0

	for _, a := range adjacencies(len(grid), len(grid[0])) {
		p := 1
		for _, c := range a {
			p *= grid[c.y][c.x]
			if p > max {
				max = p
			}
		}
	}

	fmt.Printf("%d\n", max)
}
