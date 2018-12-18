package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	inputData, err := readFileToArray("input/day11.txt")
	if err != nil {
		log.Fatal(err)
	}

	seiral, _ := strconv.Atoi(inputData[0])
	maxPower := 0
	xmax, ymax := 0, 0

	for y := 1; y <= 298; y++ {
		for x := 1; x <= 298; x++ {
			power := calcPower(x, y, seiral)
			if maxPower < power {
				maxPower = power
				xmax = x
				ymax = y
			}
		}
	}

	fmt.Printf("%v %v %v\n", maxPower, xmax, ymax)
}

func calcPower(x int, y int, serial int) int {
	sum := 0

	for i := y; i < y+3; i++ {
		for j := x; j < x+3; j++ {
			rackID := j + 10
			product := float64(((rackID * i) + serial) * rackID % 1000 / 100.0)
			d := int(math.Ceil(product))
			sum += d - 5
		}
	}

	return sum
}

func readFileToArray(fileName string) ([]string, error) {
	dataB, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(dataB), "\n")
	return split, nil
}
