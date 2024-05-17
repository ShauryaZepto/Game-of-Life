package gameoflife

import (
	"fmt"
)

func (B *Board) print_grid(){
	for i:= B.rows{
		for j := B.cols{
			if(B.grid[i][j]){
				fmt.Print("*")
			}
			else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}