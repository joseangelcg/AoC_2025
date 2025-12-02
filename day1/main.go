package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	// "strings"
)

func processRotation(curPos int ,rotation string) int{
	//Get direction of rotation
	s := rotation[0]
	var newPos int

	//Get how far we will rotate from curPos, %100 to avoid multiple turns
	turn,_ := strconv.Atoi(rotation[1:])
	turn = turn%100

	if (s == 'L') {

		//we will substract from curPos
		newPos = curPos - turn
		//if we are negative, then we overflowed. add 100
		if(newPos<0) {newPos += 100}

	} else if (s == 'R'){

		//we will substract from curPos
		newPos = curPos + turn
		//if we are more than 99, then we overflowed. substract 100
		if(newPos>99) {newPos -= 100}

	}

	return newPos
}

func main() {
	var dialPos int = 50
	var password int = 0

	f,err := os.Open("directions.txt")
	if(err != nil){
		fmt.Println("Error file does not open")
		return
	}
	//Scan directions
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		dialPos = processRotation(dialPos, scanner.Text())
		// fmt.Printf("New pos is: %d\n", dialPos)
		if(dialPos==0) {password++}
	}

	fmt.Printf("The password is: %d\n", password)

}
