package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type lightPoint struct {
	x, y, vx, vy int
	next         *lightPoint
}

type boundary struct {
	xmin, xmax, ymin, ymax int
}

func main() {
	inputData, err := readFileToArray("input/day10.txt")
	if err != nil {
		log.Fatal(err)
	}

	head := &lightPoint{}
	currentStar := head

	for _, str := range inputData {
		point := parseData(str)

		currentStar.next = &point
		currentStar = currentStar.next
	}

	currentArea := findArea(head, 0)
	nextArea := findArea(head, 1)
	time := 2

	// Find smallest area, which mean drone is closest, hopefully close to the target time
	for currentArea > nextArea {
		currentArea = nextArea
		nextArea = findArea(head, time)
		time++
	}

	fmt.Printf("Estimate time: %v\n", time)

	// plot position near by estimated time, check manually
	for t := -5; t < 5; t++ {
		currentTime := time + t

		bound := findBoundary(head, currentTime)

		plotGraph(head, currentTime, bound)
	}
}

func readFileToArray(fileName string) ([]string, error) {
	dataB, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(dataB), "\n")
	return split, nil
}

func plotGraph(head *lightPoint, time int, bound boundary) {
	xmax := bound.xmax + abs(bound.xmin) + 1
	ymax := bound.ymax + abs(bound.ymin) + 1

	// fmt.Printf("%v, %v\n", xmax, ymax)

	// Output as text
	mapper := make([][]bool, ymax)
	for y := 0; y < len(mapper); y++ {
		mapper[y] = make([]bool, xmax)
	}

	// Output as image
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{xmax, ymax}})
	cyan := color.RGBA{100, 200, 200, 0xff}

	for currentStar := head; currentStar.next != nil; currentStar = currentStar.next {
		posx := currentStar.x + time*currentStar.vx + abs(bound.xmin)
		posy := currentStar.y + time*currentStar.vy + abs(bound.ymin)
		if posx >= 0 && posx <= xmax && posy >= 0 && posy < ymax {
			img.Set(posx, posy, cyan)
		}
	}

	f, _ := os.Create(fmt.Sprintf("day10/image%d.png", time))
	png.Encode(f, img)
}

func parseData(data string) lightPoint {
	re := regexp.MustCompile("position=<(.+), (.+)> velocity=<(.+), (.+)>")
	match := re.FindStringSubmatch(data)

	if len(match) > 1 {
		posx, _ := strconv.Atoi(strings.TrimSpace(match[1]))
		posy, _ := strconv.Atoi(strings.TrimSpace(match[2]))
		velx, _ := strconv.Atoi(strings.TrimSpace(match[3]))
		vely, _ := strconv.Atoi(strings.TrimSpace(match[4]))

		return lightPoint{
			x:  posx,
			y:  posy,
			vx: velx,
			vy: vely,
		}
	}

	return lightPoint{}
}

func findBoundary(head *lightPoint, time int) boundary {
	bound := boundary{
		xmin: 0,
		xmax: 0,
		ymin: 0,
		ymax: 0,
	}

	for currentStar := head; currentStar.next != nil; currentStar = currentStar.next {
		xtemp := currentStar.x + time*currentStar.vx
		ytemp := currentStar.y + time*currentStar.vy

		if bound.xmin > xtemp {
			bound.xmin = xtemp
		} else if bound.xmax < xtemp {
			bound.xmax = xtemp
		}

		if bound.ymin > ytemp {
			bound.ymin = ytemp
		} else if bound.ymax < ytemp {
			bound.ymax = ytemp
		}
	}

	return bound
}

func findArea(head *lightPoint, time int) int {
	bound := findBoundary(head, time)
	return abs(bound.xmin-bound.xmax) * abs(bound.ymin-bound.ymax)
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
