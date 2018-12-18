package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const GENERATION = 20

type note struct {
	state, pattern string
}

func main() {
	inputData, err := readFileToArray("input/day12.txt")
	if err != nil {
		log.Fatal(err)
	}

	state := parseState(inputData[0])
	notes := []*note{}
	zeroIndex := 0

	for i := 2; i < len(inputData); i++ {
		n := parseNote(inputData[i])
		notes = append(notes, &n)
	}

	for g := 0; g < GENERATION; g++ {
		// Expand state
		first := strings.Index(state, "#")
		if first < 5 {
			remain := 5 - first
			zeroIndex += remain
			for c := 0; c < remain; c++ {
				state = "." + state
			}
		}

		last := strings.LastIndex(state, "#")
		if last > len(state)-5 {
			remain := len(state) + 4 - last
			for c := 0; c < remain; c++ {
				state = state + "."
			}
		}

		currentState := []rune(state)
		nextState := []rune(state)

		for i := 2; i < len(currentState)-2; i++ {
			str := string(state[i-2 : i+3])

			hasPlant := checkPattern(str, notes)

			if hasPlant {
				nextState[i] = '#'
			} else {
				nextState[i] = '.'
			}

		}

		currentState = nextState
		state = string(currentState)
		// fmt.Println(state)
	}

	fmt.Println(getSum(state, zeroIndex))
}

func readFileToArray(fileName string) ([]string, error) {
	dataB, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(dataB), "\n")
	return split, nil
}

func getSum(state string, zeroIndex int) int {
	sum := 0
	for i, c := range state {
		if c == '#' {
			sum += i - zeroIndex
		}
	}

	return sum
}

func checkPattern(data string, notes []*note) bool {
	state := false
	for _, n := range notes {
		if n.pattern == data {
			if n.state == "#" {
				state = true
			}

			break
		}
	}
	return state
}

func parseState(data string) string {
	state := ""
	parts := strings.Split(data, " ")
	if len(parts) == 3 {
		state = parts[2]
	}

	return state
}

func parseNote(data string) note {
	state := ""
	pattern := ""
	parts := strings.Split(data, " => ")
	if len(parts) == 2 {
		pattern = parts[0]
		state = parts[1]
	}
	return note{
		state:   state,
		pattern: pattern,
	}
}
