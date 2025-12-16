package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Present [][]rune

type Region struct {
	w      int
	h      int
	counts []int
}

type Problem struct {
	presents []Present
	regions  []Region
}

func parseFile(file string) *Problem{

	f,err := os.Open(file)
	if err != nil {
		return nil
	}

	p := new(Problem)
	p.presents = make([]Present,0)
	p.regions =  make([]Region,0)

	scanner := bufio.NewScanner(f)

	//while file has lines
	for scanner.Scan() {
		//get line
		line := scanner.Text()

		p_regex := regexp.MustCompile(`\d+:$`)
		a_regex := regexp.MustCompile(`\d+x\d+:*`)

		//lines define present
		if p_regex.MatchString(line) {
			present := Present{}
			for scanner.Scan() {
				line = scanner.Text()
				if line == "" {break}

				r := make([]rune, len(line))
				for i,c :=range line{
					r[i] = c
				}
				present = append(present, r)
			}
			p.presents = append(p.presents, present)
		} else if a_regex.MatchString(line){

			//matches region
			arr := strings.Split(line,":")
			area := strings.Split(arr[0],"x")
			counts := strings.Split(strings.TrimSpace(arr[1])," ")

			region := Region{}
			region.w ,_ = strconv.Atoi(area[0])
			region.h ,_ = strconv.Atoi(area[1])
			region.counts = make([]int, 0)

			for _,c := range(counts){
				count, _ := strconv.Atoi(c)
				region.counts = append(region.counts, count)
			}

			p.regions = append(p.regions, region)
		}
	}

	return p
}

func solveP1(p *Problem)  int{
	res := 0

	for _,r := range(p.regions){

		//lets be greedy for now and assume presents will be stacked 3x3
		// hint: turns out this was enough :P
		presents_area := 0
		for _,c := range r.counts{
			presents_area += c * 9
		}

		useful_area := (r.w - r.w%3) * (r.h - r.h%3)
		if presents_area <= useful_area{
			res++
		}
	}

	return res
}

func main() {

	prob:= parseFile(os.Args[1])
	fmt.Printf("Part1: %d\n", solveP1(prob))

}
