package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var nums []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		nums = append(nums, val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	match := map[int]bool{0: false}

	for {
		for _, n := range nums {
			sum += n

			if match[sum] {
				fmt.Println(sum)
				return
			}

			match[sum] = true
		}
	}
}
