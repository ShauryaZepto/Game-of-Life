package gameoflife

import "fmt"

func (B *Board) Update_Board() {
	temp := CreateBoard(B.rows, B.cols)
	for i := range B.rows {
		for j := range B.cols {
			number_of_neighers := B.Count_Live_Neighbours(i, j)
			if number_of_neighers < 2 || number_of_neighers > 3 {
				temp.grid[i][j] = false
			} else if number_of_neighers == 2 {
				temp.grid[i][j] = B.grid[i][j]
			} else {
				temp.grid[i][j] = true
			}
		}
	}
	copy(B.grid, temp.grid)
}

func (B *Board) Iterations(number_of_iterations int) {

	fmt.Printf("Generation %d :", 0)
	B.Print_Grid()

	for i := range number_of_iterations {
		B.Update_Board()
		fmt.Printf("Generation %d :", i+1)
		B.Print_Grid()
	}

	// fmt.Printf("Generation %d :", number_of_iterations+1)
	// B.Print_Grid()
}
