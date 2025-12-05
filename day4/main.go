package main

import (
	"bufio"
	"fmt"
	"os"
)

const neighborhoodSize int = 1

func isAccessible(ptr [][]rune, row int, col int, limit int) bool {

	left := col - neighborhoodSize
	right := col + neighborhoodSize
	up := row - neighborhoodSize
	down := row + neighborhoodSize

	if left < 0 {
		left = 0
	}
	if right >= len(ptr[0]) {
		right = len(ptr[0]) - 1
	}

	if up < 0 {
		up = 0
	}
	if down >= len(ptr) {
		down = len(ptr) - 1
	}

	count := 0
	for i := up; i <= down; i++ {
		for j := left; j <= right; j++ {
			if !(i == row && j == col) &&
				(ptr[i][j] == '@' || ptr[i][j] == 'x') {
				count++
				if count >= limit {
					return false
				}
			}
		}
	}

	return true
}

func parseFile(file string) [][]rune {

	f, err := os.Open(file)
	if err != nil {
		return nil
	}

	scanner := bufio.NewScanner(f)

	var matrix [][]rune

	for scanner.Scan() {
		var row []rune
		line := scanner.Text()
		for _, c := range line {
			row = append(row, c)
		}
		matrix = append(matrix, row)
	}

	return matrix

}

func scanRolls(rolls [][]rune, mark bool) int {
	count := 0
	for i := 0; i < len(rolls); i++ {
		for j := 0; j < len(rolls[i]); j++ {
			if rolls[i][j] == '@' && isAccessible(rolls, i, j, 4) {
				count++
				if mark {
					rolls[i][j] = 'x'
				}
			}
		}
	}
	return count
}

func scanRolls2(rolls [][]rune) int {

	count := 0
	for {
		new_count := scanRolls(rolls, true)
		if new_count == 0 {
			break
		}
		count += new_count
		removeRolls(rolls)
	}
	return count
}

func removeRolls(rolls [][]rune) {
	for i := 0; i < len(rolls); i++ {
		for j := 0; j < len(rolls[i]); j++ {
			if rolls[i][j] == 'x' {
				rolls[i][j] = '.'
			}
		}
	}
}

func main() {
	m := parseFile(os.Args[1])

	//part 1 do it once
	count := scanRolls(m, false)
	fmt.Printf("Part1: %d rolls are accessible\n", count)

	// part 2: do it until no change in count
	count = scanRolls2(m)
	fmt.Printf("Part2: %d rolls are accessible\n", count)
}
