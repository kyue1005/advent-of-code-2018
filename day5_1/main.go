package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func main() {
	inputData, err := readFileToArray("input/day5.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", polymerReact(inputData[0]))
}

func readFileToArray(fileName string) ([]string, error) {
	dataB, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(dataB), "\n")
	return split, nil
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
