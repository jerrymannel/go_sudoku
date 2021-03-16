package main

import (
	"flag"
	"fmt"
	"time"
)

//Cell value of a cells
type Cell struct {
	Value      string
	Contenders string
	Down       *Cell
	Right      *Cell
	GridMember *Cell
}

var gridBT [][]string
var debug *bool
var numberStrings = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

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

// SudokuBackTrack accepts the starting cell of the grid
// and returns a solution string
// This uses the backtracking algo
func SudokuBackTrack() bool {
	// displayGridBT()
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if gridBT[row][col] == "0" {
				for _, val := range numberStrings {
					if isValid(val, row, col) {
						gridBT[row][col] = val
						// fmt.Println(val, row, col, isValid(val, row, col))
						if *debug {
							fmt.Println(convertToLineBT(gridBT))
						}
						success := SudokuBackTrack()
						if !success {
							gridBT[row][col] = "0"
						}
					}
				}
				return gridBT[row][col] != "0"
			}
		}
	}
	return true
}

func main() {
	startTime := time.Now()
	input := flag.String("i", "", "Input puzzle")
	output := flag.String("o", "", "Expected output")
	showGrid := flag.Bool("g", false, "Show grid")
	debug = flag.Bool("d", false, "Show iterations")
	// backtracking := flag.Bool("b", true, "Use backtracking algorithm")
	flag.Parse()
	if *input == "" {
		flag.PrintDefaults()
		panic("No input provided")
	}
	var line string = *input
	var duration time.Duration
	var solution string
	gridBT = convertToGridBT(line)
	SudokuBackTrack()
	solution = convertToLineBT(gridBT)
	// if *backtracking {
	// 	// displayGridRowByRowBT()
	// } else {
	// 	duration, solution = Sudoku(line)
	// }
	duration = time.Now().Sub(startTime)
	if *output != "" {
		fmt.Println("Solved?", (solution == *output))
	}
	fmt.Printf("Time taken :: %v\n", duration)
	fmt.Printf("Solution :: %s\n", solution)
	if *showGrid {
		problem := convertToGridBT(line)
		fmt.Println("--------------- PROBLEM ------------------")
		displayGridBT(problem)
		if *output != "" {
			answer := convertToGridBT(*output)
			fmt.Println("--------------- GIVEN SOLUTION ------------------")
			displayGridBT(answer)
		}
		fmt.Println("--------------- COMPUTED SOLUTION ------------------")
		displayGridBT(gridBT)
	}
}
