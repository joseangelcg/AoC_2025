package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const newLine string = "\r\n"

func parseFile(file string)  []string {
	buf, err := os.ReadFile(file)
	if err != nil {
		return nil
	}

	//Remove trailing lines and split by empty lines
	str := strings.TrimRight(string(buf), newLine)
	lines := strings.Split(str, newLine)

	return lines
}

//function to proces file lines as part 1
func parseString1(lines []string) [][]int {

	var numbers [][]int

	var strs[][]string
	for _,line := range(lines){
		strs = append(strs, regexp.MustCompile(`\s+`).Split(strings.Trim(line, " "), -1))
	}

	for i:= 0; i<len(strs[0]);i++{
		var nums_col[]int
		for _, l := range strs {
			num, _ := strconv.Atoi(l[i])
			nums_col = append(nums_col, num)
		}
		numbers = append(numbers, nums_col)
	}

	return numbers
}

func parseOperators(line string) []string {
	ops := regexp.MustCompile(`\s+`).Split(strings.Trim(line, " "), -1)
	return ops

}

func parseString2(lines []string, num_ops int) [][]int {

	// var numbers [][]int
	numbers := make([][]int, num_ops)

	curr_op := 0
	// iterate char by char each line
	for i := 0; i < len(lines[0]); i++ {

		num := 0
		for _, line := range(lines) {
			if line[i] != ' ' {
				d, _ := strconv.Atoi(string(line[i]))
				num = num*10 + d
			}
		}

		if num > 0 {
			numbers[curr_op] = append(numbers[curr_op], num)
		} else {
			curr_op++
		}
	}

	return numbers
}

func processOperations(nums [][]int, ops []string) int {
	if len(nums) != len(ops) {
		//mismatch between ops and nums
		fmt.Println(fmt.Errorf("Sizes of nums (%d) and ops (%d) mismatch", len(nums), len(ops)))
		return -1
	}

	finalRes := 0
	for i, op := range ops {

		var res int

		if op == "*" {
			res = 1
		} else if op == "+" {
			res = 0
		}
		for _, num := range nums[i] {
			if op == "*" {
				res *= num
			} else if op == "+" {
				res += num
			}
		}
		finalRes += res
	}
	return finalRes
}

func main() {

	lines := parseFile(os.Args[1])

	//get operators from last line
	ops := parseOperators(lines[len(lines)-1])

	//parse all lines except the last one
	nums1 := parseString1(lines[:len(lines)-1])
	fmt.Printf("Part1: result is %d\n", processOperations(nums1, ops))

	nums2 := parseString2(lines[:len(lines)-1], len(ops))
	fmt.Printf("Part1: result is %d\n", processOperations(nums2, ops))
}
