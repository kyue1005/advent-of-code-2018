package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func readFileToArray(fileName string) ([]string, error) {
	dataB, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(dataB), "\n")
	return split, nil
}

func main() {
	count2 := 0
	count3 := 0

	inputData, err := readFileToArray("day2_1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, str := range inputData {
		count := make(map[rune]int)

		hasDouble := false
		hasTriple := false

		for _, c := range str {
			count[c]++
		}

		for _, v := range count {
			if v == 2 && !hasDouble {
				count2++
				hasDouble = true
			}

			if v == 3 && !hasTriple {
				count3++
				hasTriple = true
			}
		}
	}

	fmt.Printf("%v\n", count2*count3)
}
