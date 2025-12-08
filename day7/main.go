package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func countTimelines(input [][]rune) int {

	var timelines []int

	for _, c := range input[0] {
		n := 0
		if c == 'S' {
			n = 1
		}
		timelines = append(timelines, n)
	}

	//last line is not processed
	for i := 1; i < len(input)-1; i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == '^' {
				//timeline generation
				if j > 0 {
					timelines[j-1] += timelines[j]
				}
				if j < len(input[i])-1 {
					timelines[j+1] += timelines[j]
				}
				timelines[j] = 0
			}
		}
	}

	res := 0
	for _, n := range timelines {
		res += n
	}

	return res
}

func countSplits(input [][]rune) int {
	splits := 0

	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 'S' || input[i][j] == '|' {
				//check if ray can continue downward
				if input[i+1][j] == '.' {
					input[i+1][j] = '|'
				} else if input[i+1][j] == '^' {
					//ray hits splitter then
					splits++
					//ray can be duplicated.
					if j > 0 {
						input[i+1][j-1] = '|'
					}
					if j < len(input[i+1])-1 {
						input[i][j+1] = '|'
					}
				}
			}
		}
	}
	return splits

}

func main() {
	m := parseFile(os.Args[1])
	fmt.Printf("Part1: beam splits %d times\n", countSplits(m))
	fmt.Printf("Part2: beam has %d timelines\n", countTimelines(m))

}
