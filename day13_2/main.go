package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

const LEFT_TURN = 0
const STRAGHT_TURN = 1
const RIGHT_TURN = 2

const LEFT_DIR = 0
const UP_DIR = 1
const RIGHT_DIR = 2
const DOWN_DIR = 3

type cart struct {
	x, y    int
	dir     int
	turn    int
	crashed bool
}

func main() {
	inputData, err := readFileToArray("input/day13.txt")
	if err != nil {
		log.Fatal(err)
	}

	mineMap := make([][]rune, len(inputData))
	carts := []*cart{}

	for i, str := range inputData {
		mineMap[i] = []rune(str)

		for j, c := range mineMap[i] {
			switch c {
			case '^':
				carts = append(carts, &cart{
					x:       j,
					y:       i,
					dir:     UP_DIR,
					turn:    LEFT_TURN,
					crashed: false,
				})
				mineMap[i][j] = '|'
				break
			case 'v':
				carts = append(carts, &cart{
					x:       j,
					y:       i,
					dir:     DOWN_DIR,
					turn:    LEFT_TURN,
					crashed: false,
				})
				mineMap[i][j] = '|'
				break
			case '<':
				carts = append(carts, &cart{
					x:       j,
					y:       i,
					dir:     LEFT_DIR,
					turn:    LEFT_TURN,
					crashed: false,
				})
				mineMap[i][j] = '-'
				break
			case '>':
				carts = append(carts, &cart{
					x:       j,
					y:       i,
					dir:     RIGHT_DIR,
					turn:    LEFT_TURN,
					crashed: false,
				})
				mineMap[i][j] = '-'
				break
			}
		}
	}

	// printMap(mineMap, carts)
	for existCart(carts) > 1 {
		// for _, c := range carts {
		// 	fmt.Print(c)
		// }
		// fmt.Println()
		// fmt.Println("=====================STEP=====================")
		carts = sortCart(carts)
		carts = moveCart(mineMap, carts)
		// printMap(mineMap, carts)
	}

	for _, c := range carts {
		if !c.crashed {
			fmt.Println(c.x, c.y)
		}
	}
}

func moveCart(mineMap [][]rune, carts []*cart) []*cart {
	for i, c := range carts {
		if !c.crashed {
			// decide direction
			switch mineMap[c.y][c.x] {
			case '+':
				switch c.turn {
				case LEFT_TURN:
					// turn anti-clockwise
					if c.dir == 0 {
						c.dir = 3
					} else {
						c.dir--
					}
					break
				case RIGHT_TURN:
					// turn clockwise
					c.dir = (c.dir + 1) % 4
					break
				}
				// switch turn dir at next cross
				c.turn = (c.turn + 1) % 3
				break
			case '/':
				if c.dir%2 == 0 {
					// turn anti-clockwise
					if c.dir == 0 {
						c.dir = 3
					} else {
						c.dir--
					}
				} else {
					// turn clockwise
					c.dir = (c.dir + 1) % 4
				}
				break
			case '\\':
				if c.dir%2 == 0 {
					// turn clockwise
					c.dir = (c.dir + 1) % 4
				} else {
					// turn anti-clockwise
					if c.dir == 0 {
						c.dir = 3
					} else {
						c.dir--
					}
				}
				break
			}

			// move cart
			switch c.dir {
			case LEFT_DIR:
				c.x--
				break
			case UP_DIR:
				c.y--
				break
			case RIGHT_DIR:
				c.x++
				break
			case DOWN_DIR:
				c.y++
				break
			}

			for j, c2 := range carts {
				if i != j && c.x == c2.x && c.y == c2.y && !c2.crashed {
					carts[i].crashed = true
					carts[j].crashed = true
					break
				}
			}
		}
	}

	return carts
}

func sortCart(carts []*cart) []*cart {
	sort.Slice(carts, func(i int, j int) bool {
		return carts[i].x < carts[j].x
	})

	return carts
}

func existCart(carts []*cart) int {
	count := 0
	for _, c := range carts {
		if !c.crashed {
			count++
		}
	}
	return count
}

func printMap(mineMap [][]rune, carts []*cart) {
	tempMap := make([][]rune, len(mineMap))

	for i, row := range mineMap {
		tempMap[i] = make([]rune, len(row))
		copy(tempMap[i], row)
	}

	for _, c := range carts {
		if !c.crashed {
			dir := ' '
			switch c.dir {
			case LEFT_DIR:
				dir = '<'
				break
			case UP_DIR:
				dir = '^'
				break
			case RIGHT_DIR:
				dir = '>'
				break
			case DOWN_DIR:
				dir = 'v'
				break
			}
			tempMap[c.y][c.x] = dir
		}
	}

	for _, row := range tempMap {
		fmt.Println(string(row))
	}
}

func readFileToArray(fileName string) ([]string, error) {
	dataB, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(dataB), "\n")
	return split, nil
}
