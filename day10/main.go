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

type Machine struct {
	//use bitwise rep
	lights  int
	buttons []int
	joltage []int
}

type Matrix struct {
	data [][]float64
	rows int
	cols int
	ind  []int
	dep  []int
}

func parseFile(file string) []Machine {
	f, err := os.Open(file)
	if err != nil {
		return nil
	}

	var machines []Machine
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Split(line, " ")

		m := Machine{}
		//first string are lights
		str_lights := strings.Trim(line_arr[0], "[]")
		for i, c := range str_lights {
			if c == '#' {
				m.lights |= 1 << i
			}
		}

		//next in string, list of number per button between ()
		for i := 1; i < len(line_arr)-1; i++ {
			str_button := strings.Trim(line_arr[i], "()")
			str_nums := strings.Split(str_button, ",")

			button := 0
			for _, s := range str_nums {
				n, _ := strconv.Atoi(s)
				button |= 1 << n
			}
			m.buttons = append(m.buttons, button)
		}

		//last string contains joltages

		str_joltages := strings.Split(strings.Trim(line_arr[len(line_arr)-1], "{}"), ",")
		for _, j := range str_joltages {
			n, _ := strconv.Atoi(j)
			m.joltage = append(m.joltage, n)
		}

		machines = append(machines, m)
	}

	return machines
}

func (m *Matrix) printMatrix() {
	for _, r := range m.data {
		fmt.Printf("%v\n", r)
	}
	fmt.Println()
}

func (m *Matrix) swapRows(a, b int) {
	for i := 0; i <= m.cols; i++ {
		m.data[a][i], m.data[b][i] = m.data[b][i], m.data[a][i]
	}
}

func (m *Matrix) gaussianElimination() {

	//start at top left corner
	pivot_row := 0
	cur_col := 0

	for {
		//condition to exit loop
		if pivot_row >= m.rows || cur_col >= m.cols {
			break
		}

		//iterate for rows in current col to find best match
		val := 0.0
		row_val := pivot_row
		for i := pivot_row; i < m.rows; i++ {
			tmp := math.Abs(m.data[i][cur_col])
			if tmp > val {
				row_val = i
				val = tmp
			}
		}

		if val < 1e-9 {
			//value is 0, variable of cur column is independent
			//add to independent list and continue
			m.ind = append(m.ind, cur_col)
			cur_col++
			continue
		}

		//variable is dependent, lets swap...
		if row_val != pivot_row {
			//swap values
			m.swapRows(row_val, pivot_row)
		}
		//add it to the list
		m.dep = append(m.dep, cur_col)

		//and normalize the row
		divisor := m.data[pivot_row][cur_col]
		for i := cur_col; i <= m.cols; i++ {
			m.data[pivot_row][i] /= divisor
		}

		//now just add/substract this row from the others.
		for i := 0; i < m.rows; i++ {
			//obviously, dont operate over pivot row
			if i != pivot_row {
				factor := m.data[i][cur_col]
				if math.Abs(factor) > 1e-9 {
					for c := cur_col; c <= m.cols; c++ {
						m.data[i][c] -= factor * m.data[pivot_row][c]
					}
				}
			}
		}

		pivot_row++
		cur_col++
	}

	for i := cur_col; i < m.cols; i++ {
		m.ind = append(m.ind, i)
	}

}

func createMatrix(machine Machine) Matrix {

	rows := len(machine.joltage)
	cols := len(machine.buttons)

	matrix := make([][]float64, rows)
	for i := range matrix {
		matrix[i] = make([]float64, cols+1)
	}

	for c, b := range machine.buttons {

		for r := range rows {
			if b&(1<<r) != 0 {
				matrix[r][c] = 1.0
			}
		}
	}

	for i, j := range machine.joltage {
		matrix[i][cols] = float64(j)
	}

	m := Matrix{
		rows: rows,
		cols: cols,
		data: matrix,
	}

	m.gaussianElimination()

	return m
}

func (m *Matrix) solveP2(max_pres int) int {

	if len(m.ind) == 0 {
		s := 0
		for i := range len(m.dep) {
			s += int(math.Round(m.data[i][m.cols]))
		}
		return s
	}

	values := make([]int, len(m.ind))

	min_val := math.MaxInt
	for {
		if values[len(values)-1] > max_pres {
			break
		}

		total := 0
		for _, v := range values {
			total += v
		}
		valid := false

		//iterate for each row the values that are in the
		//values array for ind variables
		for r := 0; r < len(m.dep); r++ {
			res := m.data[r][m.cols]

			for c := 0; c < len(m.ind); c++ {
				res -= m.data[r][m.ind[c]] * float64(values[c])
			}

			if res < -1e-9 {
				//negative number...
				break
			}

			//number is fractional, discard
			rounded := math.Abs(math.Round(res))
			if math.Abs(res-rounded) > 1e-9 {
				break
			}

			total += int(rounded)
			if r == len(m.dep)-1 {
				valid = true
			}
		}

		//check solution
		if valid {
			min_val = min(min_val, total)
		}

		//update values for ind variables
		values[0]++
		for i := 0; i < len(values)-1; i++ {
			if values[i] > max_pres {
				values[i+1]++
				values[i] = 0
				break
			}
		}
	}

	return min_val
}

func findMinimumPressMachine(m Machine) int {

	//max presses can be len(buttons). Return len(res)+1 in case of no solution
	res := len(m.buttons) + 1

	//mask will tell wich buttons will be pressed
	for mask := 1; mask < (1 << len(m.buttons)); mask++ {
		//for each mask, lights start all off
		cur_lights := 0
		presses := 0

		for i, b := range m.buttons {
			//button i is enabled, press it
			if mask&(1<<i) != 0 {
				cur_lights = cur_lights ^ b
				presses++
			}
		}

		//after pressing all possible buttons, check
		if cur_lights == m.lights {
			res = min(res, presses)
		}
	}

	return res
}

func sumAllMinPresses(machines []Machine) int {
	sum := 0
	for _, m := range machines {
		sum += findMinimumPressMachine(m)
	}

	return sum
}

func main() {

	machines := parseFile(os.Args[1])

	//Part1
	fmt.Printf("Part1 sol: %d\n", sumAllMinPresses(machines))

	//Part2
	s := 0
	for _, mach := range machines {
		m := createMatrix(mach)
		s += m.solveP2(slices.Max(mach.joltage))
	}
	fmt.Printf("Part2 sol: %d\n", s)

}
