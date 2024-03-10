package main

import (
	"os"

	"github.com/01-edu/z01"
)

const n = 9

func main() {
	info := "Error"
	mssg := ""
	args := os.Args
	arguments := []string(args[1:])
	lenA := 0
	for index := range arguments {
		lenA = index + 1
	}
	field := make([][]rune, n)
	for i := range field {
		field[i] = make([]rune, n)
	}
	if lenA == n {
		for index := range field {
			field[index] = []rune(arguments[index])
		}
		if FindSolution(field) == false {
			mssg = info
		}
	} else {
		mssg = info
	}
	if mssg == "" {
		PrintField(field)
	} else {
		for _, l := range []rune(mssg) {
			z01.PrintRune(l)
		}
		z01.PrintRune('\n')
	}
}

func FindSolution(field [][]rune) bool {
	x := -1
	y := -1
	min := 10
	for row := 0; row < n; row++ {
		for column := 0; column < n; column++ {
			if field[row][column] == '.' {
				counter := 0
				for v := '1'; v < rune('1'+n); v++ {
					field[row][column] = rune(v)
					if CheckPos(field, row, column) {
						counter++
					}
					field[row][column] = '.'
				}
				if min > counter {
					min = counter
					x = row
					y = column
				}
			}
		}
	}
	if x == -1 {
		return true
	}
	for v := '1'; v < rune('1'+n); v++ {
		field[x][y] = rune(v)
		if CheckPos(field, x, y) {
			if FindSolution(field) {
				return true
			}
		}
	}
	field[x][y] = '.'
	return false
}

func PrintField(field [][]rune) {
	for i := range field {
		for _, l := range field[i] {
			z01.PrintRune(l)
			z01.PrintRune(' ')
		}
		z01.PrintRune('\n')
	}
}

func CheckPos(field [][]rune, row, column int) bool {
	value := field[row][column]
	count := 0
	//check in column
	for i := 0; i < n; i++ {
		if field[i][column] == value {
			count++
		}
	}
	//check in row
	for j := 0; j < n; j++ {
		if field[row][j] == value {
			count++
		}
	}
	//square 123456789 for sub-grid
	for i := (row / 3) * 3; i < (row/3+1)*3; i++ {
		for j := (column / 3) * 3; j < (column/3+1)*3; j++ {
			if field[i][j] == value {
				count++
			}
		}
	}
	if count == 3 {
		return true
	}
	return false
}
