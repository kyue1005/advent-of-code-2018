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

	matchCode := make([][]string, 16)

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
				possibleCode := []string{}
				for k, o := range opcodes {
					// fmt.Println(before, instr, after)
					result := execOpcode(o, instr, before)
					// fmt.Println(k, o, ": ", result)
					if intSliceEqual(result, after) {
						matchResultCount++
						possibleCode = append(possibleCode, k)
					}
				}
				// fmt.Println("Similar op count: ", matchResultCounto)

				if len(matchCode[instr[0]]) > 0 {
					existCode := matchCode[instr[0]]
					for i := 0; len(existCode) > i; {
						exist := false
						for _, c1 := range possibleCode {
							if existCode[i] == c1 {
								exist = true
								break
							}
						}
						if !exist {
							existCode = append(existCode[:i], existCode[i+1:]...)
						} else {
							i++
						}
					}
					matchCode[instr[0]] = existCode
				} else {
					matchCode[instr[0]] = possibleCode
				}

			}
		} else {
			emptyLineCount++
		}
		lineCount++
	}

	for knownCode := []string{}; len(knownCode) < 16; {
		for i, codes := range matchCode {
			if len(codes) == 1 {
				duplicate := false
				for _, k := range knownCode {
					if codes[0] == k {
						duplicate = true
					}
				}
				if !duplicate {
					knownCode = append(knownCode, codes[0])
				}
			} else {
				j := 0
				for len(codes) > j {
					duplicate := false
					c := codes[j]
					for _, k := range knownCode {
						if k == c {
							duplicate = true
							codes = append(codes[:j], codes[j+1:]...)
						}
					}
					if !duplicate {
						j++
					}
				}
				matchCode[i] = codes
			}

		}
	}

	reg := make([]int, 4)
	for len(inputData) > lineCount {
		instr := make([]int, 4)
		parts := strings.Split(inputData[lineCount], " ")
		for i := 0; i < 4; i++ {
			val, _ := strconv.Atoi(parts[i])
			instr[i] = val
		}
		code := opcodes[matchCode[instr[0]][0]]
		reg = execOpcode(code, instr, reg)

		lineCount++
	}
	fmt.Println(reg)
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
