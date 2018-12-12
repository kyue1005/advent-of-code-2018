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
	count := findRegion(coords, bound)

	fmt.Printf("%v\n", count)
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

func findRegion(coords []coord, bound boundary) int {
	region := 0

	for x := bound.xMin; x <= bound.xMax; x++ {
		for y := bound.yMin; y <= bound.yMax; y++ {
			if totalDist(coords, x, y) < 10000 {
				region++
			}
		}
	}

	return region
}

func totalDist(coords []coord, x int, y int) int {
	dist := 0

	for _, c := range coords {
		dist += abs(c.x-x) + abs(c.y-y)
	}

	return dist
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
