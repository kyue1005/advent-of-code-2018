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
	x, y int
	dir  int
	turn int
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
					x:    j,
					y:    i,
					dir:  UP_DIR,
					turn: LEFT_TURN,
				})
				mineMap[i][j] = '|'
				break
			case 'v':
				carts = append(carts, &cart{
					x:    j,
					y:    i,
					dir:  DOWN_DIR,
					turn: LEFT_TURN,
				})
				mineMap[i][j] = '|'
				break
			case '<':
				carts = append(carts, &cart{
					x:    j,
					y:    i,
					dir:  LEFT_DIR,
					turn: LEFT_TURN,
				})
				mineMap[i][j] = '-'
				break
			case '>':
				carts = append(carts, &cart{
					x:    j,
					y:    i,
					dir:  RIGHT_DIR,
					turn: LEFT_TURN,
				})
				mineMap[i][j] = '-'
				break
			}
		}
	}

	// printMap(mineMap, carts)

	isCollide, cX, cY := false, 0, 0
	for !isCollide {
		// fmt.Println("=====================STEP=====================")
		carts = sortCart(carts)
		isCollide, cX, cY = moveCart(mineMap, carts)
	}
	// printMap(mineMap, carts)
	fmt.Println(cX, cY)
}

func moveCart(mineMap [][]rune, carts []*cart) (bool, int, int) {
	isCollide := false
	cX, cY := 0, 0

	for i, c := range carts {
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

		// check collision
		for j, c1 := range carts {
			if i != j && c.x == c1.x && c.y == c1.y {
				isCollide = true
				cX = c.x
				cY = c.y
				break
			}
		}
	}

	return isCollide, cX, cY
}

func printMap(mineMap [][]rune, carts []*cart) {
	tempMap := make([][]rune, len(mineMap))

	for i, row := range mineMap {
		tempMap[i] = make([]rune, len(row))
		copy(tempMap[i], row)
	}

	for _, c := range carts {
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

	for _, row := range tempMap {
		fmt.Println(string(row))
	}
}

func sortCart(carts []*cart) []*cart {
	sort.Slice(carts, func(i int, j int) bool {
		return carts[i].x < carts[j].x
	})

	return carts
}

func readFileToArray(fileName string) ([]string, error) {
	dataB, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(dataB), "\n")
	return split, nil
}
