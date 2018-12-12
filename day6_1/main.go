package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type coord struct {
	x    int
	y    int
	area int
}

type boundary struct {
	xMin int
	xMax int
	yMin int
	yMax int
}

func main() {
	inputData, err := readFileToArray("input/day6.txt")
	if err != nil {
		log.Fatal(err)
	}

	coords := []coord{}

	for _, data := range inputData {
		coords = append(coords, parseCoord(data))
	}

	bound := findBoundary(coords)
	area := findLargestArea(coords, bound)

	fmt.Printf("%v\n", area)
}

func readFileToArray(fileName string) ([]string, error) {
	dataB, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(dataB), "\n")
	return split, nil
}

func findBoundary(coords []coord) boundary {
	bound := boundary{
		xMin: coords[0].x,
		xMax: coords[0].x,
		yMin: coords[0].y,
		yMax: coords[0].y,
	}

	for _, c := range coords {
		if c.x < bound.xMin {
			bound.xMin = c.x
		}
		if c.x > bound.xMax {
			bound.xMax = c.x
		}
		if c.y < bound.yMin {
			bound.yMin = c.y
		}
		if c.y > bound.yMax {
			bound.yMax = c.y
		}
	}
	return bound
}

func findLargestArea(coords []coord, bound boundary) int {
	for x := bound.xMin; x <= bound.xMax; x++ {
		for y := bound.yMin; y <= bound.yMax; y++ {
			coordIndex := closestCoordIndex(coords, x, y)

			// point not have equal min distance
			if coordIndex >= 0 {
				// mark coord has inifinte area if point lines at boudary
				if x == bound.xMin || x == bound.xMax || y == bound.yMin || y == bound.yMax {
					coords[coordIndex].area = -1
				}

				if coords[coordIndex].area >= 0 {
					coords[coordIndex].area++
				}
			}
		}
	}

	maxArea := 0

	for _, c := range coords {
		if c.area > maxArea {
			maxArea = c.area
		}
	}

	return maxArea
}

func closestCoordIndex(coords []coord, x int, y int) int {
	minIndex, minDist := 0, -1

	for i, c := range coords {
		dist := abs(c.x-x) + abs(c.y-y)

		// fmt.Printf("%v %v,%v %v\n", dist, x, y, c)
		if minDist < 0 || dist < minDist {
			minDist = dist
			minIndex = i
		} else if dist == minDist {
			minIndex = -1
		}
	}

	return minIndex
}

func parseCoord(data string) coord {
	parts := strings.Split(data, ", ")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])

	return coord{
		x:    x,
		y:    y,
		area: 0,
	}
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
