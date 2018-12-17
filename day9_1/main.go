package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inputData, err := readFileToArray("input/day9.txt")
	if err != nil {
		log.Fatal(err)
	}

	playerCnt, lastMarblePt := parseGame(inputData[0])
	// playerCnt, lastMarblePt = 10, 1618

	playerIndex, marbleIndex := 0, 0
	playerScore := make([]int, playerCnt)
	marbleChain := make([]int, 0)

	currentMarble := 0
	marbleChain = append(marbleChain, currentMarble)
	currentMarble++

	for ; currentMarble <= lastMarblePt; currentMarble++ {
		// fmt.Printf("%v\n", marbleChain)
		if currentMarble%23 == 0 {
			marbleIndex = (marbleIndex - 7) % len(marbleChain)
			if marbleIndex < 0 {
				marbleIndex += len(marbleChain)
			}
			playerScore[playerIndex] += currentMarble + marbleChain[marbleIndex]

			// fmt.Printf("%v %v %v\n", playerIndex, currentMarble, marbleChain[marbleIndex])

			copy(marbleChain[marbleIndex:], marbleChain[marbleIndex+1:])
			marbleChain[len(marbleChain)-1] = 0
			marbleChain = marbleChain[:len(marbleChain)-1]
		} else {
			marbleIndex += 2

			if marbleIndex > len(marbleChain) {
				marbleIndex = 1
			}

			marbleChain = append(marbleChain, 0)
			copy(marbleChain[marbleIndex+1:], marbleChain[marbleIndex:])
			marbleChain[marbleIndex] = currentMarble
		}

		playerIndex = currentMarble % playerCnt
	}
	// fmt.Printf("%v\n", playerScore)

	maxScore := 0
	for p := 0; p < playerCnt; p++ {
		if maxScore < playerScore[p] {
			maxScore = playerScore[p]
		}
	}

	fmt.Printf("%v\n", maxScore)
}

func readFileToArray(fileName string) ([]string, error) {
	dataB, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(dataB), "\n")
	return split, nil
}

func parseGame(data string) (int, int) {
	playerCnt, lastMarblePt := 0, 0
	re := regexp.MustCompile("([0-9]+) players; last marble is worth ([0-9]+) points")
	match := re.FindStringSubmatch(data)

	if len(match) > 0 {
		playerCnt, _ = strconv.Atoi(match[1])
		lastMarblePt, _ = strconv.Atoi(match[2])
	}
	return playerCnt, lastMarblePt
}
