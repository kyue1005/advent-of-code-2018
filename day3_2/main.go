package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type claim struct {
	id   string
	left int
	top  int
	wide int
	tall int
}

func readFileToArray(fileName string) ([]string, error) {
	dataB, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(dataB), "\n")
	return split, nil
}

func parseClaim(data string) claim {
	parts := strings.Split(data, " ")

	tempPart2 := parts[2][:len(parts[2])-1]   // remove trailing ;
	position := strings.Split(tempPart2, ",") // split by ,
	dimension := strings.Split(parts[3], "x") // split by x

	left, _ := strconv.Atoi(position[0])
	top, _ := strconv.Atoi(position[1])
	width, _ := strconv.Atoi(dimension[0])
	height, _ := strconv.Atoi(dimension[1])

	return claim{
		id:   parts[0][1:],
		left: left,
		top:  top,
		wide: width,
		tall: height,
	}
}

func main() {
	inputData, err := readFileToArray("input/day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	var claims []claim

	for _, str := range inputData {
		claims = append(claims, parseClaim(str))
	}

	var fabric [1000][1000]int

	// make claims
	for _, c := range claims {
		for x := c.left; x < c.left+c.wide; x++ {
			for y := c.top; y < c.top+c.tall; y++ {
				fabric[x][y]++
			}
		}
	}

	// Check overlap
	for _, c := range claims {
		isOverlap := false
		for x := c.left; x < c.left+c.wide && !isOverlap; x++ {
			for y := c.top; y < c.top+c.tall && !isOverlap; y++ {
				if fabric[x][y] > 1 {
					isOverlap = true
				}
			}
		}

		if !isOverlap {
			fmt.Printf("Valid claim ID: %v\n", c.id)
			break
		}
	}
}
