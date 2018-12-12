package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"
)

const WORKER_NO = 5

type node struct {
	childs   []string
	activate []string
	isChild  bool
}

type worker struct {
	node       string
	timeRemain int
}

func main() {
	inputData, err := readFileToArray("input/day7.txt")
	if err != nil {
		log.Fatal(err)
	}

	steps := make(map[string]*node)

	for _, str := range inputData {
		parent, child := parseStep(str)
		if _, ok := steps[parent]; !ok {
			steps[parent] = &node{
				childs:  []string{},
				isChild: false,
			}
		}

		if _, ok := steps[child]; !ok {
			steps[child] = &node{
				childs:  []string{},
				isChild: false,
			}
		}

		steps[parent].childs = append(steps[parent].childs, child)
		steps[child].activate = append(steps[child].activate, parent)
		steps[child].isChild = true
	}

	result := ""
	second := 0
	availableStep := []string{}

	workers := []*worker{}
	// create workers
	for i := 0; i < WORKER_NO; i++ {
		workers = append(workers, &worker{})
	}

	// find Head node
	for k, s := range steps {
		// fmt.Printf("%s, %v\n", k, s)
		if !s.isChild {
			availableStep = append(availableStep, k)
		}
	}

	for {
		sort.Strings(availableStep)

		// check finish Job
		for _, w := range workers {
			if w.timeRemain == 0 && len(w.node) > 0 {
				// job finished
				result += w.node

				for _, c := range steps[w.node].childs {
					activate := steps[c].activate
					if checkActivate(activate, result) {
						availableStep = append(availableStep, c)
					}
				}
				// reset node
				w.node = ""
			}
		}

		// assign job
		for _, w := range workers {
			if w.timeRemain == 0 {
				if len(availableStep) > 0 && len(w.node) == 0 {
					w.node = availableStep[0]
					w.timeRemain = 60 + (int(w.node[0]) - 64) // A=61, B=62, etc.

					// shift out first available node
					availableStep = availableStep[1:]

				}
			}

			if w.timeRemain > 0 {
				w.timeRemain--
			}
		}

		if len(result) == len(steps) {
			break
		}

		second++
	}

	fmt.Printf("%v %v\n", result, second)
}

func readFileToArray(fileName string) ([]string, error) {
	dataB, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(dataB), "\n")
	return split, nil
}

func checkActivate(activate []string, result string) bool {
	isActive := true
	for _, a := range activate {
		if !strings.Contains(result, string(a)) {
			isActive = false
			break
		}
	}

	return isActive
}

func parseStep(data string) (string, string) {
	p, c := "", ""

	re := regexp.MustCompile("Step ([A-Z]) must be finished before step ([A-Z]) can begin.")
	matched := re.FindAllStringSubmatch(data, -1)

	if len(matched) > 0 && len(matched[0]) > 1 {
		p, c = matched[0][1], matched[0][2]
	}

	return p, c
}
