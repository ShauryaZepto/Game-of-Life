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

func (G *Board) initialize(row int, col int) {
	G.row = row
	G.col = col
grid:
}
