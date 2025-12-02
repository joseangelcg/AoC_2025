package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processRange(rangeIn string) int {
	//get range
	arr := strings.Split(rangeIn, "-")
	start, _ := strconv.Atoi(arr[0])
	end, _ := strconv.Atoi(arr[1])

	sum := 0

	for i := start; i <= end; i++ {

		//evaluate if i complies with repeat pattern
		i_str := strconv.Itoa(i)
		if strings.Compare(i_str[:len(i_str)/2], i_str[len(i_str)/2:]) == 0 {
			sum += i
			// fmt.Printf("matched: %s\n", i_str)
		}
	}

	return sum
}

func main() {

	f, _ := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")

		for _, e := range ranges {
			total += processRange(e)
		}
	}

	fmt.Printf("sum is: %d\n", total)

}
