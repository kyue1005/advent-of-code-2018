package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

const TARGET = "846601"

func main() {
	// fmt.Printf("%v\n", inputData)
	elves := []int{0, 1}
	scoreboard := []int{3, 7}
	target := []int{}

	for _, c := range TARGET {
		d, _ := strconv.Atoi(string(c))
		target = append(target, d)
	}

	for {
		digits := splitDigit(scoreboard[elves[0]] + scoreboard[elves[1]])
		scoreboard = append(scoreboard, digits...)

		elves[0] = (elves[0] + scoreboard[elves[0]] + 1) % len(scoreboard)
		elves[1] = (elves[1] + scoreboard[elves[1]] + 1) % len(scoreboard)

		// fmt.Println(scoreboard)
		// fmt.Println(elves)
		if len(scoreboard) > len(target) {
			lastSlice := ""
			for i := len(scoreboard) - len(target); i < len(scoreboard); i++ {
				lastSlice += strconv.Itoa(scoreboard[i])
			}
			if TARGET == lastSlice {
				fmt.Println(len(scoreboard) - len(target))
				break
			}

			secondLastSlice := ""
			for i := len(scoreboard) - len(target) - 1; i < len(scoreboard)-1; i++ {
				secondLastSlice += strconv.Itoa(scoreboard[i])
			}

			if TARGET == secondLastSlice {
				fmt.Println(len(scoreboard) - len(target) - 1)
				break
			}
		}
	}

	// fmt.Println(scoreboard)

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
