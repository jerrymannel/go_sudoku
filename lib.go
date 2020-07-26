package main

import (
	"fmt"
	"strconv"
	"strings"
)

func generateGrid() *Cell {
	C00 := &Cell{0, nil, nil, nil, nil}
	C10 := &Cell{10, nil, nil, nil, nil}
	C00.Down = C10

	C20 := &Cell{20, nil, nil, nil, nil}
	C10.Down = C20

	C30 := &Cell{30, nil, nil, nil, nil}
	C20.Down = C30

	C40 := &Cell{40, nil, nil, nil, nil}
	C30.Down = C40

	C50 := &Cell{50, nil, nil, nil, nil}
	C40.Down = C50

	C60 := &Cell{60, nil, nil, nil, nil}
	C50.Down = C60

	C70 := &Cell{70, nil, nil, nil, nil}
	C60.Down = C70

	C80 := &Cell{80, nil, nil, nil, nil}
	C70.Down = C80

	p1 := C00
	p2 := C10
	p3 := C20
	p4 := C30
	p5 := C40
	p6 := C50
	p7 := C60
	p8 := C70
	p9 := C80
	for x := 0; x < 8; x++ {
		p1.Right = &Cell{00 + x + 1, nil, nil, nil, nil}
		p1 = p1.Right

		p2.Right = &Cell{10 + x + 1, nil, p1, nil, nil}
		p2 = p2.Right
		p1.Down = p2

		p3.Right = &Cell{20 + x + 1, nil, p2, nil, nil}
		p3 = p3.Right
		p2.Down = p3

		p4.Right = &Cell{30 + x + 1, nil, p3, nil, nil}
		p4 = p4.Right
		p3.Down = p4

		p5.Right = &Cell{40 + x + 1, nil, p4, nil, nil}
		p5 = p5.Right
		p4.Down = p5

		p6.Right = &Cell{50 + x + 1, nil, p5, nil, nil}
		p6 = p6.Right
		p5.Down = p6

		p7.Right = &Cell{60 + x + 1, nil, p6, nil, nil}
		p7 = p7.Right
		p6.Down = p7

		p8.Right = &Cell{70 + x + 1, nil, p7, nil, nil}
		p8 = p8.Right
		p7.Down = p8

		p9.Right = &Cell{80 + x + 1, nil, p8, nil, nil}
		p9 = p9.Right
		p8.Down = p9
	}
	return C00
}

func convertToGrid(_grid *Cell, _line string) {
	puzzle := strings.Split(_line, "")
	fmt.Println(puzzle)
	head := _grid
	for index, element := range puzzle {
		_grid.Value, _ = strconv.Atoi(element)
		_grid = _grid.Right
		if index != 0 && (index+1)%9 == 0 {
			_grid = head.Down
			head = _grid
		}
	}
}

func findMissingValuesAtEachCell(_grid *Cell) {

}
