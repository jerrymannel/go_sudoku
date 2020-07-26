package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

//Cell value of a cell
type Cell struct {
	Value      int
	Contenders []int
	Down       *Cell
	Right      *Cell
	GridMember *Cell
}

func main() {
	startTime := time.Now()
	runTests := flag.Bool("t", false, "Run and print test results")
	fmt.Println("Go Sudoku 1.0")
	fmt.Println(*runTests)

	line := "058409020000060089201030700000013400320080901079000608060005817815700304030090200,658479123743162589291538746586913472324687951179254638962345817815726394437891265"
	lines := strings.Split(line, ",")
	subStartTime := time.Now()
	fmt.Println("Split line", (time.Now().Sub(subStartTime)))

	subStartTime = time.Now()
	grid := generateGrid()
	convertToGrid(grid, lines[0])
	fmt.Println("Convert to grid", (time.Now().Sub(subStartTime)))
	displayGrid(grid)

	fmt.Println("Time to solve -", (time.Now().Sub(startTime)))
}
