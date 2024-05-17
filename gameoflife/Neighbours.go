package gameoflife

func (B *Board) count_live_neighbours(x int, y int) int {
	count := 0
	for i:= x-1; i<=x+1: i++ {
		for j:= y-1; j<=y+1; j++{
			if(i==x&&j==y){
				continue
			}
			if(check_valid_index(i,j)){
				if(B.grid[i][j]){
					count++
				}
			}
		}
	}
	return count
}
