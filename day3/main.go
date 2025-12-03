package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func allocateArray(col int, rows int) [][]int {

	arr := make([][]int, rows)
	for i := range len(arr) {
		arr[i] = make([]int, col)
	}

	return arr
}

func processBank(line string, m int) int {

	//allocate our memoized array of m*len(line)
	dp := allocateArray(len(line)+1, m+1)

	//parse line to int array to process
	arr := make([]int, len(line))
	for i, c := range line {
		arr[i], _ = strconv.Atoi(string(c))
	}

	//implement memoized dp
	for i := 1; i < len(dp); i++ {
		for j := i; j < len(dp[i]); j++ {
			dp[i][j] = max(dp[i-1][j-1]*10+arr[j-1], dp[i][j-1])
		}
	}
	return dp[m][len(line)]
}

func main() {

	f, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Error file does not open")
		return
	}

	//Scan banks
	scanner := bufio.NewScanner(f)
	joltage1 := 0
	joltage2 := 0

	for scanner.Scan() {
		joltage1 += processBank(scanner.Text(), 2)
		joltage2 += processBank(scanner.Text(), 12)
	}

	fmt.Printf("The maximum joltage in Part1 is: %d\n", joltage1)
	fmt.Printf("The maximum joltage in Part2 is: %d\n", joltage2)

}
