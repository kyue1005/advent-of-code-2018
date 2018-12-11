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
	inputData, err := readFileToArray("template/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", inputData)
}
