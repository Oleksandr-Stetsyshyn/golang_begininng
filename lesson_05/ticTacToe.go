package ticTacToe

import (
	"fmt"
	"strings"
)

type Board struct {
	cells map[string]string
}

func (b *Board) Print() {
	fmt.Println("   1   2   3")
	fmt.Println("A ", b.cells["A1"], "|", b.cells["A2"], "|", b.cells["A3"])
	fmt.Println("  ---+---+---")
	fmt.Println("B ", b.cells["B1"], "|", b.cells["B2"], "|", b.cells["B3"])
	fmt.Println("  ---+---+---")
	fmt.Println("C ", b.cells["C1"], "|", b.cells["C2"], "|", b.cells["C3"])
}

func NewBoard() *Board {
	cells := map[string]string{
		"A1": " ",
		"A2": " ",
		"A3": " ",
		"B1": " ",
		"B2": " ",
		"B3": " ",
		"C1": " ",
		"C2": " ",
		"C3": " ",
	}
	return &Board{cells: cells}
}

func (b *Board) CheckWin(symbol string) bool {
	winPatterns := [][]string{
		{"A1", "A2", "A3"},
		{"B1", "B2", "B3"},
		{"C1", "C2", "C3"},
		{"A1", "B1", "C1"},
		{"A2", "B2", "C2"},
		{"A3", "B3", "C3"},
		{"A1", "B2", "C3"},
		{"A3", "B2", "C1"},
	}
	for _, winCells := range winPatterns {

		if b.cells[winCells[0]] == symbol &&
			b.cells[winCells[1]] == symbol &&
			b.cells[winCells[2]] == symbol {
			return true
		}
	}

	return false
}

func clearBoard() {
	fmt.Print("\033[H\033[2J")
}

func StartGame() {
	board := NewBoard()

	var symbol string = "X"
	var winner string = ""
	clearBoard()
	for {
		var input string
		fmt.Println("Player", symbol, "enter cell coordinates")
		board.Print()

		fmt.Scan(&input)
		input = strings.ToUpper(strings.TrimSpace(input))

		_, isCellExist := board.cells[input]

		if !isCellExist {
			clearBoard()
			fmt.Println("!!!Such a cell does not exist!!!")
			continue
		}

		if board.cells[input] != " " {
			clearBoard()
			fmt.Println("!!!This cell is already occupied!!!")
			continue
		}

		board.cells[input] = symbol
		if board.CheckWin(symbol) {
			winner = symbol
			break
		}

		clearBoard()
		if symbol == "X" {
			symbol = "O"
		} else {
			symbol = "X"
		}

	}
	clearBoard()
	board.Print()
	fmt.Printf("🎉🎊 Player %s won. 🎊🎉", winner)
}
