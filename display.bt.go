package main

import "fmt"

func displayGridBT(gridBT [][]string) {
	drawHorizontalLine("top")
	for y := 0; y < 9; y++ {
		fmt.Print("│")
		for x := 0; x < 9; x++ {
			if gridBT[y][x] == "0" {
				fmt.Print("   ")
			} else {
				fmt.Print(" " + gridBT[y][x] + " ")
			}
			if (x+1)%3 == 0 {
				fmt.Print("│")
			}
		}
		fmt.Println()
		if (y+1)%3 == 0 {
			switch y {
			case 8:
				drawHorizontalLine("bottom")
			default:
				drawHorizontalLine("")
			}
		}
	}
}

func displayGridRowByRowBT() {
	for y := 0; y < 9; y++ {
		fmt.Printf("R%d : %v\n", (y + 1), gridBT[y])
	}
}
