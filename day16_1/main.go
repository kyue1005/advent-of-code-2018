package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var opcodes = map[string]op{
	"addr": {action: '+', a: 'r', b: 'r'},
	"addi": {action: '+', a: 'r', b: 'v'},
	"mulr": {action: '*', a: 'r', b: 'r'},
	"muli": {action: '*', a: 'r', b: 'v'},
	"banr": {action: '&', a: 'r', b: 'r'},
	"bani": {action: '&', a: 'r', b: 'v'},
	"borr": {action: '|', a: 'r', b: 'r'},
	"bori": {action: '|', a: 'r', b: 'v'},
	"setr": {action: 'a', a: 'r', b: 'r'},
	"seti": {action: 'a', a: 'v', b: 'r'},
	"gtir": {action: '>', a: 'v', b: 'r'},
	"gtri": {action: '>', a: 'r', b: 'v'},
	"gtrr": {action: '>', a: 'r', b: 'r'},
	"eqir": {action: '=', a: 'v', b: 'r'},
	"eqri": {action: '=', a: 'r', b: 'v'},
	"eqrr": {action: '=', a: 'r', b: 'r'},
}

func main() {
	inputData, err := readFileToArray("input/day16.txt")
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	lineCount := 0
	emptyLineCount := 0

	for emptyLineCount < 3 {
		if len(inputData[lineCount]) > 0 {
			emptyLineCount = 0
			parts := strings.Split(inputData[lineCount], ": [")

			if parts[0] == "Before" {
				before, instr, after := make([]int, 4), make([]int, 4), make([]int, 4)

				// parse Before
				parts[1] = parts[1][:len(parts[1])-1]
				bStrParts := strings.Split(parts[1], ", ")
				for i := 0; i < 4; i++ {
					before[i], _ = strconv.Atoi(bStrParts[i])
				}

				lineCount++
				// parse op
				parts = strings.Split(inputData[lineCount], " ")
				for i := 0; i < 4; i++ {
					val, _ := strconv.Atoi(parts[i])
					instr[i] = val
				}

				lineCount++
				// parse After
				parts = strings.Split(inputData[lineCount], ":  [")
				parts[1] = parts[1][:len(parts[1])-1]
				aStrParts := strings.Split(parts[1], ", ")
				for i := 0; i < 4; i++ {
					after[i], _ = strconv.Atoi(aStrParts[i])
				}

				matchResultCount := 0
				for _, o := range opcodes {
					// fmt.Println(before, instr, after)
					result := execOpcode(o, instr, before)
					// fmt.Println(k, o, ": ", result)
					if intSliceEqual(result, after) {
						matchResultCount++
						// fmt.Println("Equal ", k)
					}
				}
				// fmt.Println("Similar op count: ", matchResultCounto)

				if matchResultCount >= 3 {
					count++
				}
			}
		} else {
			emptyLineCount++
		}
		lineCount++
	}

	fmt.Println(count)
}

func intSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func execOpcode(code op, instr []int, reg []int) []int {
	result := make([]int, 4)
	for i := 0; i < 4; i++ {
		result[i] = reg[i]
	}
	a, b := instr[1], instr[2]

	if code.a == 'r' {
		a = reg[a]
	}

	if code.b == 'r' {
		b = reg[b]
	}

	switch code.action {
	case '+':
		result[instr[3]] = a + b
		break
	case '*':
		result[instr[3]] = a * b
		break
	case '&':
		result[instr[3]] = a & b
		break
	case '|':
		result[instr[3]] = a | b
		break
	case 'a':
		result[instr[3]] = a
		break
	case '>':
		if a > b {
			result[instr[3]] = 1
		} else {
			result[instr[3]] = 0
		}
		break
	case '=':
		if a == b {
			result[instr[3]] = 1
		} else {
			result[instr[3]] = 0
		}
		break
	}

	return result
}

func readFileToArray(fileName string) ([]string, error) {
	dataB, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(dataB), "\n")
	return split, nil
}

type op struct {
	a          rune
	b          rune
	action     rune
	matchCount []byte
}
