package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	// "strings"
)

func processRotation(curPos int, rotation string) (int, int) {
	//Get direction of rotation
	s := rotation[0]
	var newPos int

	//Get how far we will rotate from curPos, %100 to avoid multiple turns
	turn, _ := strconv.Atoi(rotation[1:])
	overs := turn / 100
	turn = turn % 100

	if s == 'L' {

		//we will substract from curPos
		newPos = curPos - turn
		//if we are negative, then we overflowed. add 100
		if newPos < 0 {
			newPos += 100

			//only add if we didnt start nor land on 0.
			if curPos != 0 && newPos != 0 {
				overs++
			}
		}

	} else if s == 'R' {

		//we will substract from curPos
		newPos = curPos + turn
		//if we are more than 99, then we overflowed. substract 100
		if newPos > 99 {
			newPos -= 100
			//only add if we didnt start nor land on 0.
			if curPos != 0 && newPos != 0 {
				overs++
			}
		}
	}

	//landed on 0. increment over count here
	if newPos == 0 {
		overs++
	}

	fmt.Printf("Old Position %d moves: %s to position %d, overs: %d\n", curPos, rotation, newPos, overs)

	return newPos, overs
}

func main() {
	var dialPos int = 50
	var password int = 0

	f, err := os.Open("directions.txt")
	if err != nil {
		fmt.Println("Error file does not open")
		return
	}
	//Scan directions
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var overs int = 0
		dialPos, overs = processRotation(dialPos, scanner.Text())
		//Add number of times we clicked over 0 or landed
		password += overs
	}

	fmt.Printf("The password is: %d\n", password)

}
