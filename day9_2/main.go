package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

/*
	rewrite by using double linked list to avoid array operation to save time
*/

type Marble struct {
	value    int
	next     *Marble
	previous *Marble
}

func main() {
	inputData, err := readFileToArray("input/day9.txt")
	if err != nil {
		log.Fatal(err)
	}

	playerCnt, lastMarblePt := parseGame(inputData[0])
	lastMarblePt = lastMarblePt * 100
	playerScore := make([]int, playerCnt)

	currentMarbleValue := 0
	currentMarble := &Marble{value: currentMarbleValue}
	currentMarbleValue++

	currentMarble.next = &Marble{value: currentMarbleValue, next: currentMarble, previous: currentMarble}
	currentMarble.previous = currentMarble.next
	currentMarbleValue++

	for ; currentMarbleValue <= lastMarblePt; currentMarbleValue++ {
		// fmt.Printf("%v\n", marbleChain)
		if currentMarbleValue%23 == 0 {
			playerIndex := currentMarbleValue % playerCnt

			for i := 0; i < 7; i++ {
				currentMarble = currentMarble.previous
			}
			marbleForRemoval := currentMarble
			marbleForRemoval.next.previous = marbleForRemoval.previous
			marbleForRemoval.previous.next = marbleForRemoval.next

			playerScore[playerIndex] += currentMarbleValue + currentMarble.value

			// fmt.Printf("%v %v %v\n", playerIndex, currentMarble, marbleChain[marbleIndex])
			currentMarble = marbleForRemoval.next
		} else {
			newMarble := &Marble{
				value:    currentMarbleValue,
				next:     currentMarble.next.next,
				previous: currentMarble.next,
			}
			newMarble.previous.next = newMarble
			newMarble.next.previous = newMarble
			currentMarble = newMarble
		}
	}

	fmt.Println(max(playerScore))
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

func max(x []int) int {
	max := 0
	for _, v := range x {
		if v > max {
			max = v
		}
	}
	return max
}
