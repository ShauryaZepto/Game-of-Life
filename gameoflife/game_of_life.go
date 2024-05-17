package main

import (
	"fmt"
)

var n, m int = 100, 100
var initial_alive, iterations_number int

func make_grid(grid [][]bool) {
	for i := range n {
		grid[i] = make([]bool, m)
	}
}

func set_alive(grid [][]bool) {
	for i := 0; i < initial_alive; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		grid[x][y] = true
	}
}

func check_valid_index(x, y int) bool {
	if x < 0 || x >= n || y < 0 || y >= m {
		return false
	}
	return true
}

func count_neighbours(grid [][]bool, x int, y int) int {
	count := 0
	for i := range 3 {
		for j := range 3 {
			if i == 1 && j == 1 {
				continue
			}
			dx, dy := x-1+i, y-1+j
			if check_valid_index(dx, dy) {
				if grid[dx][dy] {
					count++
				}
			}
		}
	}
	return count
}

func print_grid(grid [][]bool) {
	for i := range n {
		for j := range m {
			if grid[i][j] {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func iterations_generation(grid [][]bool) {
	temp_grid := make([][]bool, n)
	make_grid(temp_grid)
	fmt.Printf("Generation 0 :")
	print_grid(grid)
	for iter := range iterations_number {

		for i := range n {
			for j := range m {
				number_of_neighers := count_neighbours(grid, i, j)
				if number_of_neighers < 2 || number_of_neighers > 3 {
					temp_grid[i][j] = false
				} else if number_of_neighers == 2 {
					temp_grid[i][j] = grid[i][j]
				} else {
					temp_grid[i][j] = true
				}
			}
		}
		fmt.Printf("Generation %d :", iter+1)
		print_grid(grid)
	}
}

func main() {
	fmt.Print("Enter the grid size (n * m): ")
	// fmt.Scan(&n,&m)

	grid := make([][]bool, n)
	make_grid(grid)

	fmt.Print("Enter the number of alive cells: ")
	fmt.Scan(&initial_alive)

	fmt.Print("Enter the number of iterations: ")
	fmt.Scan(&iterations_number)

	iterations_generation(grid)

}
