package main

import (
	"strings"
)

func generateGrid() *Cell {
	numberSet := "123456789"
	C00 := &Cell{"00", numberSet, nil, nil, nil}
	C10 := &Cell{"10", numberSet, nil, nil, nil}
	C00.Down = C10

	C20 := &Cell{"20", numberSet, nil, nil, nil}
	C10.Down = C20

	C30 := &Cell{"30", numberSet, nil, nil, nil}
	C20.Down = C30

	C40 := &Cell{"40", numberSet, nil, nil, nil}
	C30.Down = C40

	C50 := &Cell{"50", numberSet, nil, nil, nil}
	C40.Down = C50

	C60 := &Cell{"60", numberSet, nil, nil, nil}
	C50.Down = C60

	C70 := &Cell{"70", numberSet, nil, nil, nil}
	C60.Down = C70

	C80 := &Cell{"80", numberSet, nil, nil, nil}
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
		p1.Right = &Cell{"0", numberSet, nil, nil, nil}
		p1 = p1.Right

		p2.Right = &Cell{"1", numberSet, nil, nil, nil}
		p2 = p2.Right
		p1.Down = p2

		p3.Right = &Cell{"2", numberSet, nil, nil, nil}
		p3 = p3.Right
		p2.Down = p3

		p4.Right = &Cell{"3", numberSet, nil, nil, nil}
		p4 = p4.Right
		p3.Down = p4

		p5.Right = &Cell{"4", numberSet, nil, nil, nil}
		p5 = p5.Right
		p4.Down = p5

		p6.Right = &Cell{"5", numberSet, nil, nil, nil}
		p6 = p6.Right
		p5.Down = p6

		p7.Right = &Cell{"6", numberSet, nil, nil, nil}
		p7 = p7.Right
		p6.Down = p7

		p8.Right = &Cell{"7", numberSet, nil, nil, nil}
		p8 = p8.Right
		p7.Down = p8

		p9.Right = &Cell{"8", numberSet, nil, nil, nil}
		p9 = p9.Right
		p8.Down = p9
	}

	head := C00
	for i := 0; i < 3; i++ {
		r := head
		for j := 0; j < 3; j++ {
			c := r
			for k := 0; k < 3; k++ {
				c.GridMember = c.Right
				c.Right.GridMember = c.Right.Right
				if k != 2 {
					c.Right.Right.GridMember = c.Down
				}
				c = c.Down
			}
			r = r.Right.Right.Right
		}
		head = head.Down.Down.Down
	}
	return C00
}

func convertToGrid(_grid *Cell, _line string) {
	puzzle := strings.Split(_line, "")
	head := _grid
	for index, element := range puzzle {
		_grid.Value = element
		if _grid.Value != "0" {
			_grid.Contenders = ""
		}
		_grid = _grid.Right
		if index != 0 && (index+1)%9 == 0 {
			_grid = head.Down
			head = _grid
		}
	}
}

func convertToLine(_startingCell *Cell) string {
	line := ""
	for _startingCell != nil {
		rCell := _startingCell
		for rCell != nil {
			line += rCell.Value
			rCell = rCell.Right
		}
		_startingCell = _startingCell.Down
	}
	return line
}

func updateMissingValuesForEachGrid(_startingCell *Cell) {
	for i := 0; i < 3; i++ {
		r := _startingCell
		for j := 0; j < 3; j++ {
			cell := r
			for cell != nil {
				if cell.Value != "0" {
					c := r
					for c != nil {
						if strings.Index(c.Contenders, cell.Value) != -1 && c.Value == "0" {
							c.Contenders = strings.ReplaceAll(c.Contenders, cell.Value, "")
						}
						c = c.GridMember
					}
				}
				cell = cell.GridMember
			}
			r = r.Right.Right.Right
		}
		_startingCell = _startingCell.Down.Down.Down
	}
}

func updateMissingValuesForEachRow(_startingCell *Cell) {
	for _startingCell != nil {
		rCell := _startingCell
		for rCell != nil {
			if rCell.Value != "0" {
				c := _startingCell
				for c != nil {
					if strings.Index(c.Contenders, rCell.Value) != -1 && c.Value == "0" {
						c.Contenders = strings.ReplaceAll(c.Contenders, rCell.Value, "")
					}
					c = c.Right
				}
			}
			rCell = rCell.Right
		}
		_startingCell = _startingCell.Down
	}
}

func updateMissingValuesForEachColumn(_startingCell *Cell) {
	for _startingCell != nil {
		rCell := _startingCell
		for rCell != nil {
			if rCell.Value != "0" {
				c := _startingCell
				for c != nil {
					if strings.Index(c.Contenders, rCell.Value) != -1 && c.Value == "0" {
						c.Contenders = strings.ReplaceAll(c.Contenders, rCell.Value, "")
					}
					c = c.Down
				}
			}
			rCell = rCell.Down
		}
		_startingCell = _startingCell.Right
	}
}

func cleanUpEachRow(_startingCell *Cell) {
	// fmt.Println("cleanUpEachRow()")
	for _startingCell != nil {
		rCell := _startingCell
		frequencyMap := make(map[string]int)
		for rCell != nil {
			if rCell.Contenders != "" {
				contenders := strings.Split(rCell.Contenders, "")
				for _, contender := range contenders {
					if _, ok := frequencyMap[contender]; !ok {
						frequencyMap[contender] = 0
					}
					frequencyMap[contender]++
				}
			}
			rCell = rCell.Right
		}
		// fmt.Println(frequencyMap)
		rCell = _startingCell
		for rCell != nil {
			if rCell.Contenders != "" {
				contenders := strings.Split(rCell.Contenders, "")
				for _, contender := range contenders {
					if v, ok := frequencyMap[contender]; ok && v == 1 {
						rCell.Value = contender
						rCell.Contenders = ""
					}
				}
			}
			rCell = rCell.Right
		}
		_startingCell = _startingCell.Down
	}
}

func cleanUpEachColumn(_startingCell *Cell) {
	// fmt.Println("cleanUpEachColumn()")
	for _startingCell != nil {
		rCell := _startingCell
		frequencyMap := make(map[string]int)
		for rCell != nil {
			if rCell.Contenders != "" {
				contenders := strings.Split(rCell.Contenders, "")
				for _, contender := range contenders {
					if _, ok := frequencyMap[contender]; !ok {
						frequencyMap[contender] = 0
					}
					frequencyMap[contender]++
				}
			}
			rCell = rCell.Down
		}
		// fmt.Println(frequencyMap)
		rCell = _startingCell
		for rCell != nil {
			if rCell.Contenders != "" {
				contenders := strings.Split(rCell.Contenders, "")
				for _, contender := range contenders {
					if v, ok := frequencyMap[contender]; ok && v == 1 {
						rCell.Value = contender
						rCell.Contenders = ""
					}
				}
			}
			rCell = rCell.Down
		}
		_startingCell = _startingCell.Right
	}
}

func cleanUpGrid(_startingCell *Cell) {
	// fmt.Println("cleanUpGrid()")
	rCell := _startingCell
	frequencyMap := make(map[string]int)
	for rCell != nil {
		if rCell.Contenders != "" {
			contenders := strings.Split(rCell.Contenders, "")
			for _, contender := range contenders {
				if _, ok := frequencyMap[contender]; !ok {
					frequencyMap[contender] = 0
				}
				frequencyMap[contender]++
			}
		}
		rCell = rCell.GridMember
	}
	// fmt.Println(frequencyMap)
	rCell = _startingCell
	for rCell != nil {
		if rCell.Contenders != "" {
			contenders := strings.Split(rCell.Contenders, "")
			for _, contender := range contenders {
				if v, ok := frequencyMap[contender]; ok && v == 1 {
					rCell.Value = contender
					rCell.Contenders = ""
				}
			}
		}
		rCell = rCell.GridMember
	}
}

func cleanUpEachGrid(_startingCell *Cell) {
	for i := 0; i < 3; i++ {
		r := _startingCell
		for j := 0; j < 3; j++ {
			cleanUpGrid(r)
			r = r.Right.Right.Right
		}
		_startingCell = _startingCell.Down.Down.Down
	}
}

func solved(_startingCell *Cell) bool {
	for _startingCell != nil {
		rCell := _startingCell
		for rCell != nil {
			if len(rCell.Contenders) != 0 {
				return false
			}
			rCell = rCell.Right
		}
		_startingCell = _startingCell.Down
	}
	return true
}

func findMissingValuesAtEachCell(_startingCell *Cell) {
	updateMissingValuesForEachGrid(_startingCell)
	updateMissingValuesForEachRow(_startingCell)
	updateMissingValuesForEachColumn(_startingCell)
}
