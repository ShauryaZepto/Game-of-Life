package gameoflife

type Board struct {
	rows int
	cols int
	grid [][]bool
}

func CreateBoard(rows int, cols int) *Board {
	board := &Board{
		rows: rows,
		cols: cols,
	}
	board.grid = make([][]bool, rows)
	for i := range rows {
		board.grid[i] = make([]bool, cols)
	}
	return board
}

func (B *Board) Set_Board(matrix [][]bool, rows int, cols int) {
	for i := range rows {
		for j := range cols {
			B.grid[i][j] = matrix[i][j]
		}
	}
}

func (B *Board) Check_Valid_Index(x int, y int) bool {
	if x < 0 || x >= B.rows || y < 0 || y >= B.cols {
		return false
	}
	return true
}
