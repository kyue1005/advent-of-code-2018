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
	inputData, err := readFileToArray("day2_2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, str := range inputData {
		for _, val := range inputData {
			diffCount := 0
			same := ""
			for i, r := range val {
				if str[i] != byte(r) {
					diffCount++
				} else {
					same += string(r)
				}
			}

			if diffCount == 1 {
				fmt.Printf("%s\n", same)
				return
			}
		}
	}
}
