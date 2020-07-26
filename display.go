package main

import (
	"fmt"
	"strconv"
)

func drawHorizontalLine(pos string) {
	s := "├─"
	switch pos {
	case "top":
		s = "┌─"
	case "bottom":
		s = "└─"
	}
	for i := 0; i < 9; i++ {
		s += "──"
		if i == 2 {
			switch pos {
			case "top":
				s += "┬"
			case "bottom":
				s += "┴"
			default:
				s += "┼"
			}
		} else if i == 5 {
			switch pos {
			case "top":
				s += "─┬"
			case "bottom":
				s += "─┴"
			default:
				s += "─┼"
			}
		} else {
			s += "─"
		}
	}
	switch pos {
	case "top":
		s += "┐"
	case "bottom":
		s += "┘"
	default:
		s += "┤"
	}
	fmt.Println(s)
}

//DisplayGrid display the sudoku _grid
func displayGrid(_grid *Cell) {
	head := _grid
	drawHorizontalLine("top")
	yCounter := 0
	for head != nil {
		fmt.Print("│")
		xCounter := 0
		for _grid != nil {
			if _grid.Value == 0 {
				fmt.Print(" " + " " + " ")
			} else {
				fmt.Print(" " + strconv.Itoa(_grid.Value) + " ")
			}
			xCounter++
			_grid = _grid.Right
			if xCounter%3 == 0 {
				fmt.Print("│")
			}
		}
		_grid = head.Down
		head = _grid
		yCounter++
		fmt.Println()
		if yCounter%3 == 0 {
			switch yCounter {
			case 9:
				drawHorizontalLine("bottom")
			default:
				drawHorizontalLine("")
			}
		}
	}
}
