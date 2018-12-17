package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type node struct {
	child       int
	childValues []int
	metadata    int
	value       int
}

func main() {
	inputData, err := readFileToArray("input/day8.txt")
	if err != nil {
		log.Fatal(err)
	}

	tree := strings.Split(inputData[0], " ")
	nodeStack := []*node{}
	var lastNode *node

	for {
		if len(nodeStack) > 0 {
			lastNode = nodeStack[len(nodeStack)-1]
		} else {
			lastNode = &node{
				child:    1,
				metadata: 1,
			}
		}

		if len(tree) > 0 {
			if lastNode.child > 0 {
				c, _ := strconv.Atoi(tree[0])
				m, _ := strconv.Atoi(tree[1])
				tree = tree[2:]

				if c == 0 {
					childSum := 0

					for i := 0; i < m; i++ {
						val, _ := strconv.Atoi(tree[0])
						childSum += val
						tree = tree[1:]
					}

					lastNode.child--
					lastNode.childValues = append(lastNode.childValues, childSum)
				} else {
					nodeStack = append(nodeStack, &node{
						child:    c,
						metadata: m,
					})
				}
			} else {
				nodeSum := 0

				for i := 0; i < lastNode.metadata; i++ {
					val, _ := strconv.Atoi(tree[0])

					if len(lastNode.childValues) >= val {
						nodeSum += lastNode.childValues[val-1]
					}
					tree = tree[1:]
				}

				lastNode.value = nodeSum

				if len(nodeStack) > 1 {
					nodeStack = nodeStack[:len(nodeStack)-1]

					lastNode = nodeStack[len(nodeStack)-1]
					lastNode.childValues = append(lastNode.childValues, nodeSum)

					lastNode.child--
				}
			}
		} else {
			break
		}
	}

	fmt.Printf("%v\n", nodeStack[0].value)
}

func readFileToArray(fileName string) ([]string, error) {
	dataB, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(dataB), "\n")
	return split, nil
}
