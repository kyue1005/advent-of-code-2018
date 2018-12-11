package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

const SLEEP = 0
const AWAKE = 1
const START = 2

type event struct {
	minute    int
	eventType int
	guard     int
}

type guard struct {
	sleepTime    int
	sleepMinutes [60]int
}

func main() {
	inputData, err := readFileToArray("input/day4.txt")
	if err != nil {
		log.Fatal(err)
	}

	sort.Strings(inputData)

	guards := make(map[int]*guard)

	dutyGuard := 0
	startSleepMinute := 0

	for _, str := range inputData {
		e := parseEvent(str)

		if e.eventType == SLEEP {
			startSleepMinute = e.minute
		} else if e.eventType == AWAKE {
			g := guards[dutyGuard]
			for i := startSleepMinute; i < e.minute; i++ {
				g.sleepMinutes[i]++
			}
			g.sleepTime += e.minute - startSleepMinute
			guards[dutyGuard] = g
		} else {
			dutyGuard = e.guard
			if _, ok := guards[e.guard]; !ok {
				guards[e.guard] = &guard{}
			}
		}
	}

	// Strategy 2, the guard most frequently asleep on the same minute
	maxCount, maxSleepGuard, maxSleepMin := 0, 0, 0
	for gid, g := range guards {
		for min, count := range g.sleepMinutes {
			if count > maxCount {
				maxCount = count
				maxSleepMin = min
				maxSleepGuard = gid
			}
		}
	}

	fmt.Printf("%v\n", maxSleepMin*maxSleepGuard)
}

func readFileToArray(fileName string) ([]string, error) {
	dataB, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(dataB), "\n")
	return split, nil
}

func parseEvent(data string) event {
	parts := strings.Split(data, "] ")

	timeParts := strings.Split(parts[0], ":")
	minute, _ := strconv.Atoi(timeParts[1])

	guardNo, eventType := 0, 0
	if strings.Contains(parts[1], "falls asleep") {
		eventType = SLEEP
	} else if strings.Contains(parts[1], "wakes up") {
		eventType = AWAKE
	} else {
		strs := strings.Split(parts[1], " ")
		guardNo, _ = strconv.Atoi(strs[1][1:len(strs[1])])
		eventType = START
	}

	return event{
		minute:    minute,
		eventType: eventType,
		guard:     guardNo,
	}
}
