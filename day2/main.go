package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkInvalid(str string) bool {
	return str[:len(str)/2] == str[len(str)/2:]
}

func checkInvalid2(str string) bool {

	// check all possible subgroups up to half the string
	for i := 0; i < len(str)/2; i++ {
		invalid := true

		//base group to compare to
		a := str[0 : i+1]

		for j := i + 1; j < len(str); j += i + 1 {

			//if group to compare overflows str, then pattern is valid,
			//jump to next group iteration
			if j+i+1 > len(str) {
				invalid = false
				break
			}

			//create new group to compare to base
			b := str[j : j+i+1]
			if a != b {
				invalid = false
				break
			}
		}

		// after iterating over groups, we always matched with base group.
		// pattern is invalid
		if invalid {
			return true
		}
	}

	return false
}

func processRange(rangeIn string, checkInvalid func(string) bool) int {
	//get range
	arr := strings.Split(rangeIn, "-")
	start, _ := strconv.Atoi(arr[0])
	end, _ := strconv.Atoi(arr[1])

	sum := 0

	for i := start; i <= end; i++ {

		//evaluate if i complies with repeat pattern
		if checkInvalid(strconv.Itoa(i)) {
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
	total2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")

		for _, e := range ranges {
			total += processRange(e, checkInvalid)
			total2 += processRange(e, checkInvalid2)
		}
	}

	fmt.Printf("Sum1 is: %d\n", total)
	fmt.Printf("Sum2 is: %d\n", total2)

}
