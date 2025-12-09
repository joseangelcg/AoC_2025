package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
	z int
}

type Distance struct {
	dist float64
	a    int
	b    int
}

func parseFile(file string) []Point {

	f, err := os.Open(file)
	if err != nil {
		return nil
	}

	scanner := bufio.NewScanner(f)

	var points []Point

	for scanner.Scan() {
		line := scanner.Text()

		p := strings.Split(line, ",")
		x, _ := strconv.Atoi(p[0])
		y, _ := strconv.Atoi(p[1])
		z, _ := strconv.Atoi(p[2])

		points = append(points, Point{x, y, z})
	}

	return points
}

//return sorted distances
func calculateDistances(points []Point) []Distance {

	var distances []Distance

	for i := range points {
		for j := i + 1; j < len(points); j++ {
			dx := points[i].x - points[j].x
			dy := points[i].y - points[j].y
			dz := points[i].z - points[j].z

			d := math.Sqrt(
				math.Pow(float64(dx), 2) +
					math.Pow(float64(dy), 2) +
					math.Pow(float64(dz), 2))
			distances = append(distances, Distance{d, i, j})
		}
	}

	slices.SortFunc(distances, func(a Distance, b Distance) int {
		if a.dist > b.dist {
			return 1
		}
		if a.dist < b.dist {
			return -1
		}
		return 0
	})

	return distances
}

func findID(IDs []int, i int) int {
	if IDs[i] != i {
		return findID(IDs, IDs[i])
	}
	return i
}

func connectID(IDs []int, i int, j int) bool {
	id_i := findID(IDs, i)
	id_j := findID(IDs, j)

	if id_i == id_j {
		return false
	}
	//Perform connection only if ids are diferent
	if id_i < id_j {
		IDs[id_j] = IDs[id_i]
	} else if id_j < id_i {
		IDs[id_i] = IDs[id_j]
	}

	return true
}

func performNConnections(distances []Distance, points []Point, connections int) int {

	circuitID := make([]int, len(points))
	circuit_sizes := make([]int, len(points))

	for i := range len(circuitID) {
		circuitID[i] = i
	}

	//perform n connections
	n := 0
	for i := range connections {
		a := distances[i].a
		b := distances[i].b
		connectID(circuitID, a, b)
	}

	for i := range circuitID {
		circuit_ID := findID(circuitID, i)
		circuit_sizes[circuit_ID]++
	}

	slices.Sort(circuit_sizes)
	n = len(circuit_sizes)

	return circuit_sizes[n-1] * circuit_sizes[n-2] * circuit_sizes[n-3]
}

func connectAll(distances []Distance, points []Point) int {

	circuitID := make([]int, len(points))

	for i := range len(circuitID) {
		circuitID[i] = i
	}

	//perform connections until all points are connected
	n := len(points) - 1
	res := 0
	for i := range distances {
		a := distances[i].a
		b := distances[i].b

		//check if new connection was made
		if connectID(circuitID, a, b) {
			n--
		}

		if n == 0 {
			res = points[a].x * points[b].x
			break
		}
	}
	return res
}
func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Incorrect number of arguments\n")
		return
	}
	file := os.Args[1]
	n, _ := strconv.Atoi(os.Args[2])
	points := parseFile(file)
	distances := calculateDistances(points)

	//Part1
	fmt.Printf("Part1: magic number is: %d\n", performNConnections(distances, points, n))

	//Part2
	fmt.Printf("Part2: magic number is: %d\n", connectAll(distances, points))

}
