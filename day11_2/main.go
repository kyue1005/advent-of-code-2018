package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const GRID_SIZE = 300

func main() {
	inputData, err := readFileToArray("input/day11.txt")
	if err != nil {
		log.Fatal(err)
	}

	serial, _ := strconv.Atoi(inputData[0])
	maxPower := 0
	xmax, ymax, sizeMax := 0, 0, 0
	powers := make([][]int, GRID_SIZE)

	for y := 1; y <= GRID_SIZE; y++ {
		row := make([]int, GRID_SIZE)
		for x := 1; x <= GRID_SIZE; x++ {
			row[x-1] = calcPower(x, y, serial)
		}
		powers[y-1] = row
	}

	for s := 1; s <= GRID_SIZE; s++ {
		bound := 300 - s + 1
		for y := 1; y <= bound; y++ {
			for x := 1; x <= bound; x++ {
				power := calcGridPower(x, y, s, powers)

				if maxPower < power {
					maxPower = power
					xmax = x
					ymax = y
					sizeMax = s
				}
			}
		}
	}

	fmt.Println(maxPower, xmax, ymax, sizeMax)
}

func calcGridPower(x int, y int, size int, powers [][]int) int {
	power := 0

	for i := y; i < y+size; i++ {
		for j := x; j < x+size; j++ {
			power += powers[i-1][j-1]
		}
	}
	return power
}

func calcPower(x int, y int, serial int) int {
	rackID := x + 10
	val := ((rackID * y) + serial) * rackID
	level := (val % 1000 / 100) - 5
	return level
}

func readFileToArray(fileName string) ([]string, error) {
	dataB, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(dataB), "\n")
	return split, nil
}
