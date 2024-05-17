package gameoflife

import (
	"fmt"
)

func (B *Board) Print_Grid() {
	for i := range B.rows {
		for j := range B.cols {
			if B.grid[i][j] {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
