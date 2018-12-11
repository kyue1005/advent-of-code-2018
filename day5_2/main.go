package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

const ALPHABET = "abcdefghijklmnopqrstuvwxyz"

func main() {
	inputData, err := readFileToArray("input/day5.txt")
	if err != nil {
		log.Fatal(err)
	}

	minLength := len(inputData[0])
	targetUnit := ""

	// Try remove unit A to Z
	for _, c := range ALPHABET {
		str := removePolymerUnit(string(c), inputData[0])
		resultLength := polymerReact(str)

		if resultLength < minLength {
			targetUnit = string(c)
			minLength = resultLength
		}
	}

	fmt.Printf("%v %v\n", targetUnit, minLength)
}

func readFileToArray(fileName string) ([]string, error) {
	dataB, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(dataB), "\n")
	return split, nil
}

func removePolymerUnit(char string, str string) string {
	char = strings.ToLower(char)
	newStr := strings.Replace(str, char, "", -1)

	char = strings.ToUpper(char)
	newStr = strings.Replace(newStr, char, "", -1)

	return newStr
}

func polymerReact(str string) int {
	index := 0

	for index < len(str)-1 {
		// Compare current rune with next one
		floatDiff := float64(int(str[index]) - int(str[index+1]))

		if int(math.Abs(floatDiff)) == 32 {
			newStr := ""

			if index > 0 {
				newStr += str[:index]
			}

			newStr += str[index+2:]
			str = newStr

			index = 0 // reset index
		} else {
			index++
		}
	}

	return len(str)
}
