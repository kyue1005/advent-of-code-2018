package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

const ELF = 0
const GOBLIN = 1

type point struct {
	x, y int
}

type unit struct {
	role int
	pos  point
	hp   int
	atk  int
}

func main() {
	inputData, err := readFileToArray("input/day15.txt")
	if err != nil {
		log.Fatal(err)
	}

	isAllElvesSurvive := false
	elfAtk := 3

	for !isAllElvesSurvive {
		elfAtk++
		caveMap, units := resetCave(inputData, elfAtk)

		// fmt.Println("++++++++ ELFATK: ", elfAtk, " ++++++++++")
		// printCave(caveMap)
		elfCount := 0
		for _, u := range units {
			if u.role == ELF {
				elfCount++
			}
		}

		round := 0
		isEnd := false
		for !isEnd {
			fullRound := 0
			// fmt.Println("++++++++ ROUND: ", round, " ++++++++++")
			sortUnit(units)

			for _, u := range units {
				if u.hp > 0 {
					fullRound++
					moveUnit(u, units, caveMap)
					combat(u, units, caveMap)
					isEnd = isEndGame(units)

					if isEnd {
						break
					}
				}
			}

			aliveCount := 0
			for _, u := range units {
				if u.hp > 0 {
					aliveCount++
				}
			}

			// printCave(caveMap)

			if fullRound >= aliveCount {
				round++
			}
		}

		elfSurvive := 0
		for _, u := range units {
			if u.role == ELF && u.hp > 0 {
				elfSurvive++
			}
		}
		// fmt.Println(elfCount, elfSurvive)
		isAllElvesSurvive = elfCount == elfSurvive

		if elfCount == elfSurvive {
			outcome := 0
			for _, u := range units {
				if u.hp > 0 {
					outcome += u.hp
				}
			}
			fmt.Println(elfAtk, round, outcome*round)
		}
	}

}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func adjacent(pos point) []point {
	return []point{
		point{x: pos.x - 1, y: pos.y},
		point{x: pos.x, y: pos.y - 1},
		point{x: pos.x + 1, y: pos.y},
		point{x: pos.x, y: pos.y + 1},
	}
}

func chooseMove(source point, target point, caveMap [][]rune) point {
	possibeMove := adjacent(source)
	for _, pt := range possibeMove {
		if pt.x == target.x && pt.y == target.y {
			return target
		}
	}

	result := nearestPoint(target, possibeMove, caveMap)
	result = sortPoint(result)

	return result[0]
}

func combat(u *unit, units []*unit, caveMap [][]rune) {
	possibleTargets := []*unit{}
	for _, u1 := range units {
		if u.role != u1.role && u1.hp > 0 && manhattanDist(u.pos, u1.pos) == 1 {
			possibleTargets = append(possibleTargets, u1)
		}
	}

	if len(possibleTargets) > 0 {
		sortUnit(possibleTargets)
		// for _, u := range possibleTargets {
		// 	fmt.Print(u)
		// }
		// fmt.Println()
		target := possibleTargets[0]
		for _, t := range possibleTargets {
			if t.hp < target.hp {
				target = t
			}
		}

		target.hp -= u.atk

		// fmt.Println(u.pos, " attack ", target.pos)

		if target.hp <= 0 {
			caveMap[target.pos.y][target.pos.x] = '.'
		}
	}
}

func countTeamUnit(units []*unit) []int {
	result := make([]int, 2)

	for _, u := range units {
		if u.hp > 0 {
			result[u.role]++
		}
	}

	return result
}

func isEndGame(units []*unit) bool {
	isEnd := false
	count := countTeamUnit(units)

	for _, t := range count {
		if t == 0 {
			isEnd = true
		}
	}

	return isEnd
}

func manhattanDist(source, target point) int {
	return abs(source.x-target.x) + abs(source.y-target.y)
}

func moveUnit(u *unit, units []*unit, caveMap [][]rune) {
	// fmt.Println(u)
	targets := []unit{}
	for j, u1 := range units {
		if u1.hp > 0 && u.role != u1.role {
			targets = append(targets, *units[j])
		}
	}

	inRange := []point{}
	for _, t := range targets {
		inRange = append(inRange, unitRange(t.pos, caveMap)...)
	}
	// fmt.Println(inRange)

	isInRange := false
	for _, r := range inRange {
		if r.x == u.pos.x && r.y == u.pos.y {
			isInRange = true
			break
		}
	}

	if !isInRange {
		// choose closest range
		choosenRange := nearestPoint(u.pos, inRange, caveMap)

		if len(choosenRange) > 0 {
			// choose first read order point if have multiple chosen range
			choosenRange = sortPoint(choosenRange)
			// fmt.Println("Range: ", choosenRange)
			move := chooseMove(u.pos, choosenRange[0], caveMap)
			// fmt.Println("Move: ", move)

			updateCave(u.pos, move, caveMap)
			u.pos = move
		}
	}
}

func nearestPoint(source point, targets []point, caveMap [][]rune) []point {
	result := []point{}
	visited := make([][]bool, len(caveMap))

	for j, row := range caveMap {
		visited[j] = make([]bool, len(row))
		for i, r := range row {
			if r != '.' {
				visited[j][i] = true
			}
		}
	}

	isReachTarget := false
	layer := adjacent(source)
	for len(layer) != 0 && !isReachTarget {
		nextlayer := []point{}
		for _, pt := range layer {
			if pt.x < 0 || pt.x >= len(visited[0]) || pt.y < 0 || pt.y >= len(visited) || visited[pt.y][pt.x] {
				continue
			}

			for _, t := range targets {
				if t.x == pt.x && t.y == pt.y {
					isReachTarget = true
					result = append(result, t)
				}
			}

			visited[pt.y][pt.x] = true
			nextlayer = append(nextlayer, adjacent(pt)...)
		}

		layer = nextlayer
	}

	return result
}

func printCave(caveMap [][]rune) {
	for _, row := range caveMap {
		for _, c := range row {
			fmt.Print(string(c))
		}
		fmt.Println()
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

func resetCave(data []string, elfAtk int) ([][]rune, []*unit) {
	caveMap := make([][]rune, len(data))
	units := []*unit{}

	for j, row := range data {
		caveMap[j] = make([]rune, len(row))
		for i, c := range row {
			if c == 'G' {
				units = append(units, &unit{
					role: GOBLIN,
					pos:  point{x: i, y: j},
					hp:   200,
					atk:  3,
				})
			} else if c == 'E' {
				units = append(units, &unit{
					role: ELF,
					pos:  point{x: i, y: j},
					hp:   200,
					atk:  elfAtk,
				})
			}
			caveMap[j][i] = c
		}
	}

	return caveMap, units
}

func sortUnit(units []*unit) {
	sort.Slice(units, func(i, j int) bool {
		if units[i].pos.y < units[j].pos.y {
			return true
		} else if units[i].pos.y == units[j].pos.y && units[i].pos.x < units[j].pos.x {
			return true
		}
		return false
	})
}

func sortPoint(points []point) []point {
	sort.Slice(points, func(i, j int) bool {
		if points[i].y < points[j].y {
			return true
		} else if points[i].y == points[j].y && points[i].x < points[j].x {
			return true
		}
		return false
	})
	return points
}

func unitRange(pos point, caveMap [][]rune) []point {
	inRange := []point{}
	neighbours := adjacent(pos)

	for _, n := range neighbours {
		// not wall
		if caveMap[n.y][n.x] != '#' {
			inRange = append(inRange, n)
		}
	}
	return inRange
}

func updateCave(src, dest point, caveMap [][]rune) {
	mark := caveMap[src.y][src.x]
	caveMap[src.y][src.x] = '.'
	caveMap[dest.y][dest.x] = mark
}
