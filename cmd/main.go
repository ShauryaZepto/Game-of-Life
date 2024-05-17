package main

import (
	"fmt"
	"gameoflife/gameoflife"
	"log"
)

var path string = "/Users/shaurya/Desktop/Game_of_life/configs/config1.yml"

func main() {

	config, e := gameoflife.InitializeConfig(path)

	if e != nil {
		log.Fatalf("Error reading data from config file: %v", e)
	}

	simulation := gameoflife.CreateBoard(config.Rows, config.Cols)
	simulation.Set_Board(config.Matrix, config.Rows, config.Cols)

	fmt.Printf("Generation 0 : \n")
	simulation.Print_Grid()
	simulation.Simulate(config.Iterations)
	fmt.Printf("Generation %d :\n", config.Iterations)
	simulation.Print_Grid()

}
