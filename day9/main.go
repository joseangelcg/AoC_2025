package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func parseFile(file string) []Point{
	f, err := os.Open(file)
	if err != nil {
		return nil
	}

	var points []Point
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		p := strings.Split(line, ",")
		x, _ := strconv.Atoi(p[0])
		y, _ := strconv.Atoi(p[1])

		points = append(points, Point{x, y})
	}

	return points
}

func sortAndUnique(a []int) []int  {
	if len(a) == 0 {
		return nil
	}

	slices.Sort(a)
	j:=0
	for i:=1;i< len(a);i++{
		if a[i] !=a[j] {
			j++
			a[j] = a[i]
		}
	}

	return a[:j+1]
}

//returns compressed coordinates and maps to each value
func compressPoints(points []Point) ([]Point, map[int]int, map[int]int){

	xpoints:= make([]int, len(points))
	ypoints:= make([]int, len(points))

	//create separate lists
	for i:=range points{
		xpoints[i] = points[i].x
		ypoints[i] = points[i].y
	}

	//sort and eliminate duplicates
	xunique := sortAndUnique(xpoints)
	yunique := sortAndUnique(ypoints)

	//create maps
	xMap := make(map[int]int)
	yMap := make(map[int]int)

	for i,v := range(xunique){
		xMap[v]=i
	}

	for i,v := range(yunique){
		yMap[v]=i
	}

	//create output array
	compressed := make([]Point, len(points))

	for i,p := range(points){
		x_c := xMap[p.x]
		y_c := yMap[p.y]
		compressed[i] = Point{x_c,y_c}
	}

	return compressed, xMap, yMap
}

func generatePoligon(points []Point, size_x int, size_y int) [][]rune {

	//generate empty x*y grid
	grid := make([][]rune,size_y)
	for i := range(grid){
		grid[i] = make([]rune, size_x)
		for j:= range(grid[i]){
			grid[i][j] = 'X'
		}
	}

	//paint perimeter
	for i := 0; i<len(points);i++{
		a := points[i]
		b := points[(i+1)%len(points)]

		//we paint over y axis
		if(a.x == b.x){

			y_begin := min(a.y,b.y)
			y_end:= max(a.y,b.y)

			for y:=y_begin; y<=y_end;y++{
				grid[y][a.x] = '#'
			}
		}else if(a.y == b.y){

			x_begin := min(a.x,b.x)
			x_end:= max(a.x,b.x)

			for x:=x_begin; x<=x_end;x++{
				grid[a.y][x] = '#'
			}
		}
	}

	//paint exterior by DFS ing all points in the borders that are not edges
	queue := []Point{}
	for r:=0; r < size_y; r++{
		for c:=0; c < size_x; c++{
			if (r==0 || r==size_y-1|| c==0||c==size_x-1) && grid[r][c] != '#'{
				grid[r][c] = '.'
				queue= append(queue,Point{c,r})
			}
		}
	}

	directions := []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue)>0 {
		//Dequeue
		curr := queue[0]
		queue = queue[1:]

		for _,dir := range(directions){
			r,c := curr.y+dir.y, curr.x+dir.x
			//# means border   /  . means node is already visited
			if r>0 && r<size_y && c>0 && c<size_x && grid[r][c]!='#' && grid[r][c] != '.'{
				grid[r][c] = '.'
				queue = append(queue, Point{c,r})
			}
		}

	}

	return grid
}

func isRectEnclosed(a Point, b Point, grid [][]rune, xMap map[int]int, yMap map[int]int) bool{
	x1 := xMap[a.x]
	x2 := xMap[b.x]
	y1 := yMap[a.y]
	y2 := yMap[b.y]

	for i:=min(x1,x2) ; i<=max(x1,x2); i++{
		if grid[y1][i] == '.' || grid[y2][i] == '.' {
			return false
		}
	}

	for i:=min(y1,y2) ; i<=max(y1,y2); i++{
		if grid[i][x1] == '.' || grid[i][x2] == '.' {
			return false
		}
	}

	return true

}

func intAbs(num int) int{
	if num<0 {
		return -1*num
	}
	return num
}

func calcBiggestArea(points []Point) int{
	res := 0

	for i := range points{
		for j:=i+1; j<len(points);j++{
			dx := intAbs(points[i].x - points[j].x)
			dy := intAbs(points[i].y - points[j].y)
			area := (dx+1)*(dy+1)
			if area > res{
				res = area
			}
		}
	}

	return res
}

func printGrid(grid [][]rune)  {
	for _,row:=range grid{
		for _,c:=range(row){
			fmt.Printf("%c",c)
		}
		fmt.Println()
	}

}

func main() {

	//gt points
	points := parseFile(os.Args[1])

	//Part1
	fmt.Printf("Part1: %d\n",calcBiggestArea(points))

	//Part2
	compressed, xMap, yMap := compressPoints(points)
	grid := generatePoligon(compressed, len(xMap), len(yMap))

	res:= 0
	for i := range points{
		for j:=i+1; j<len(points);j++{
			if(isRectEnclosed(points[i], points[j], grid, xMap,yMap)){
				dx := intAbs(points[i].x - points[j].x)
				dy := intAbs(points[i].y - points[j].y)
				area := (dx+1)*(dy+1)
				if area > res{
					res = area
				}
			}
		}
	}
	//Part2
	fmt.Printf("Part2: %d\n",res)


}
