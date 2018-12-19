package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

const STEP = 846601

func main() {
	// fmt.Printf("%v\n", inputData)
	elves := []int{0, 1}
	scoreboard := []int{3, 7}

	targetLen := STEP + STEP + 1

	for len(scoreboard) < targetLen {
		digits := splitDigit(scoreboard[elves[0]] + scoreboard[elves[1]])
		scoreboard = append(scoreboard, digits...)

		elves[0] = (elves[0] + scoreboard[elves[0]] + 1) % len(scoreboard)
		elves[1] = (elves[1] + scoreboard[elves[1]] + 1) % len(scoreboard)

		// fmt.Println(scoreboard)
		// fmt.Println(elves)
	}

	sum := ""
	for i := STEP; i < STEP+10; i++ {
		sum += strconv.Itoa(scoreboard[i])
	}

	fmt.Println(sum)
}

func splitDigit(num int) []int {
	digits := []int{}
	n := 1
	for exp := 0; exp <= num; n++ {
		exp = int(math.Pow(10.0, float64(n)))
	}

	for n--; n > 0; n-- {
		exp := int(math.Pow(10.0, float64(n)))
		d := num % exp / (exp / 10)
		digits = append(digits, d)
	}

	return digits
}

func readFileToArray(fileName string) ([]string, error) {
	dataB, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(dataB), "\n")
	return split, nil
}
