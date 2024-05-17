package gameoflife

type Board struct {
	rows int
	cols int
	grid [][]bool
}

func createBoard(rows int, cols int) *Board {
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

func (B *Board) check_valid_index(x int, y int) bool {
	if x < 0 || x >= B.rows || y < 0 || y >= B.cols {
		return false
	}
	return true
}
