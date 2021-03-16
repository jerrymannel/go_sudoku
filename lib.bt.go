package main

import (
	"strings"
)

func convertToGridBT(_line string) [][]string {
	grid := [][]string{{}, {}, {}, {}, {}, {}, {}, {}, {}}
	puzzle := strings.Split(_line, "")
	for y := 0; y < 9; y++ {
		grid[y] = puzzle[(y*8)+y : (y*8)+y+9]
	}
	return grid
}

func convertToLineBT(grid [][]string) string {
	line := ""
	for x := 0; x < 9; x++ {
		line += strings.Join(grid[x], "")
	}
	return line
}

func solvedBT() bool {
	for y := 0; y < 9; y++ {
		row := strings.Join(gridBT[y], "")
		if strings.Index(row, "0") != -1 {
			return false
		}
	}
	return true
}

func isValidEntryForRow(_entry string, _rowIndex int) bool {
	row := strings.Join(gridBT[_rowIndex], "")
	if strings.Index(row, _entry) != -1 {
		return false
	}
	return true
}

func isValidEntryForColumn(_entry string, _colIndex int) bool {
	for y := 0; y < 9; y++ {
		if gridBT[y][_colIndex] == _entry {
			return false
		}
	}
	return true
}

func isValidEntryForGrid(_entry string, _rowIndex int, _colIndex int) bool {
	_rowIndex = _rowIndex - (_rowIndex % 3)
	_colIndex = _colIndex - (_colIndex % 3)
	for rowIndex := _rowIndex; rowIndex < _rowIndex+3; rowIndex++ {
		for colIndex := _colIndex; colIndex < _colIndex+3; colIndex++ {
			if gridBT[rowIndex][colIndex] == _entry {
				return false
			}
		}
	}
	return true
}

func isValid(_entry string, _rowIndex int, _colIndex int) bool {
	return isValidEntryForGrid(_entry, _rowIndex, _colIndex) && isValidEntryForRow(_entry, _rowIndex) && isValidEntryForColumn(_entry, _colIndex)
}
