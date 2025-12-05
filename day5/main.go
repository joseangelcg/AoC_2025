package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const newLine string = "\r\n"

func parseFile(file string) ([][]int, []int) {
	buf, err := os.ReadFile(file)
	if err != nil {
		return nil, nil
	}

	//Remove trailing lines and split by empty lines
	str := strings.TrimRight(string(buf), newLine)
	stringSets := strings.Split(str, newLine+newLine)

	//Obtain array with ranges
	var range_arr [][]int
	for _, r_line := range strings.Split(stringSets[0], newLine) {
		var r []int
		range_s := strings.Split(r_line, "-")
		for _, num_s := range range_s {
			num, _ := strconv.Atoi(num_s)
			r = append(r, num)
		}
		range_arr = append(range_arr, r)
	}

	//Obtain array with items in shelf
	var shelf []int

	for _, shelf_line := range strings.Split(stringSets[1], newLine) {
		shelf_i, _ := strconv.Atoi(shelf_line)
		shelf = append(shelf, shelf_i)
	}

	return range_arr, shelf
}

func countAvailable(freshRanges [][]int, items []int) int {
	count := 0

	//iterate through items
	for _, item := range items {
		for _, r := range freshRanges {
			if item >= r[0] && item <= r[1] {
				//item is fresh, increment and break
				count++
				break
			}
		}
	}
	return count
}

func sortRanges(ranges [][]int) {
	slices.SortFunc(ranges, func(a, b []int) int {
		return a[0] - b[0]
	})
}

func mergeRanges(freshRanges [][]int) [][]int {

	var mergedRanges [][]int
	//sort ranges
	sortRanges(freshRanges)
	mergedRanges = append(mergedRanges, freshRanges[0])
	mergedRange_index := 0

	//merge them basically by cheking limits
	for i := 1; i < len(freshRanges); i++ {

		if freshRanges[i][0] <= freshRanges[i-1][1] {

			if freshRanges[i][1] > freshRanges[i-1][1] {
				//update upper limit of current merged range
				mergedRanges[mergedRange_index][1] = freshRanges[i][1]
			}
		} else {
			// if range was not mergeable, append it and update index
			mergedRanges = append(mergedRanges, freshRanges[i])
			mergedRange_index++
		}

	}

	return mergedRanges
}

func main() {

	fresh, shelf := parseFile(os.Args[1])
	if fresh == nil || shelf == nil {
		return
	}

	//Part 1
	availableCount := countAvailable(fresh, shelf)
	fmt.Printf("Part1: There are %d available items.\n", availableCount)

	//Part 2
	m := mergeRanges(fresh)
	freshCount := 0
	for _, r := range m {
		freshCount += r[1] - r[0] + 1
	}
	fmt.Printf("Part2: There are %d fresh items.\n", freshCount)

}
