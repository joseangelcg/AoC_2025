package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const UNVISITED = -1

type Graph struct {
	vertices []Vertex
	edges    map[Vertex][]Vertex
}

type Vertex = string

type Edge struct {
	src string
	dst string
}

func parseFile(file string) *Graph {
	f, err := os.Open(file)
	if err != nil {
		return nil
	}

	g := &Graph{
		vertices: make([]Vertex, 0),
		edges:    make(map[Vertex][]Vertex),
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		a := strings.Split(line, ":")

		//add vertex to vertices array
		v := a[0]
		g.vertices = append(g.vertices, v)

		//add edges
		edges := strings.Split(strings.Trim(a[1], " "), " ")
		for _, e := range edges {
			g.edges[v] = append(g.edges[v], e)
		}
	}

	return g
}

func (g *Graph) dfs_aux(paths map[Vertex]int, src, dst Vertex) int {

	//found a path
	if src == dst {
		return 1
	}

	//if path is determined for vertex, retun it
	if paths[src] != UNVISITED {
		return paths[src]
	}

	//explore edges
	s := 0
	for _, v := range g.edges[src] {
		s += g.dfs_aux(paths, v, dst)
	}

	//save result of accumulation in paths
	paths[src] = s
	return s

}

func (g *Graph) dfs(src, dst Vertex) int {
	//create a map to store path count from specific vertex
	paths := make(map[Vertex]int)
	for _, v := range g.vertices {
		paths[v] = UNVISITED
	}

	g.dfs_aux(paths, src, dst)

	return paths[src]
}

func solveP1(g *Graph) int {

	res := g.dfs("you", "out")
	return res
}

func solveP2(g *Graph) int {

	// svr -> dac -> fft -> out
	res := g.dfs("svr", "dac") * g.dfs("dac", "fft") * g.dfs("fft", "out")
	// svr -> fft -> dac -> out
	res += g.dfs("svr", "fft") * g.dfs("fft", "dac") * g.dfs("dac", "out")

	return res
}

func main() {
	g := parseFile(os.Args[1])

	//part 1
	fmt.Printf("part1: %d\n", solveP1(g))

	//part 2
	fmt.Printf("part2: %d\n", solveP2(g))

}
