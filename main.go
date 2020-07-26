package main

import (
	"flag"
	"fmt"
	"time"
)

//Cell value of a cell
type Cell struct {
	Value      string
	Contenders string
	Down       *Cell
	Right      *Cell
	GridMember *Cell
}

// Sudoku accepts a string represenation of the grid
// and returns a solution string
func Sudoku(line string) (time.Duration, string) {
	startTime := time.Now()
	grid := generateGrid()
	convertToGrid(grid, line)
	for !solved(grid) {
		findMissingValuesAtEachCell(grid)
		cleanUpEachRow(grid)
		findMissingValuesAtEachCell(grid)
		cleanUpEachColumn(grid)
		findMissingValuesAtEachCell(grid)
		cleanUpGrid(grid)
	}
	// displayGridRowBsyRow(grid)
	// displayGrid(grid)
	return time.Now().Sub(startTime), convertToLine(grid)
}

func main() {
	input := flag.String("i", "", "Input puzzle")
	output := flag.String("o", "", "Expected output")
	flag.Parse()
	if *input == "" {
		flag.PrintDefaults()
		panic("No input provided")
	}
	var line string = *input
	duration, solution := Sudoku(line)
	if *output != "" {
		fmt.Println("Solved?", (solution == *output))
	}
	fmt.Printf("%v,%s\n", duration, solution)
}
