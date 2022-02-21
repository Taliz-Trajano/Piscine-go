package main

import (
	"fmt"
	"os"
)

func printBoard(gameBoard *[9][9]int) {
	for y := 0; y < 9; y++ {
		if y == 3 || y == 6 {
			fmt.Println("")
		} else {
			fmt.Println("")
		}
		for x := 0; x < 9; x++ {
			if x == 3 || x == 6 {
				fmt.Print(gameBoard[y][x])
			} else {
				fmt.Print(gameBoard[y][x])
			}
		}
	}
	fmt.Println("")
}

func unitValid(unit [9]int) bool {
	for value := 1; value <= 9; value++ {
		count := 0
		for index := 0; index < 9; index++ {
			if unit[index] == value {
				count++
			}
		}
		if count > 1 {
			return false
		}
	}
	return true
}

func getRow(gameBoard *[9][9]int, row int) [9]int {
	return gameBoard[row-1]
}

func getCol(gameBoard *[9][9]int, col int) [9]int {
	var column [9]int
	for row := 0; row < 9; row++ {
		column[row] = gameBoard[row][col-1]
	}
	return column
}

func getBlock(gameBoard *[9][9]int, col, row int) [9]int {
	i := whatBlock(col)*3 - 2
	j := whatBlock(row)*3 - 2
	var block [9]int
	block[0] = gameBoard[j-1][i-1]
	block[1] = gameBoard[j-1][i]
	block[2] = gameBoard[j-1][i+1]
	block[3] = gameBoard[j][i-1]
	block[4] = gameBoard[j][i]
	block[5] = gameBoard[j][i+1]
	block[6] = gameBoard[j+1][i-1]
	block[7] = gameBoard[j+1][i]
	block[8] = gameBoard[j+1][i+1]
	return block
}

func whatBlock(val int) int {
	if val >= 1 && val <= 3 {
		return 1
	} else if val >= 4 && val <= 6 {
		return 2
	} else if val >= 7 && val <= 9 {
		return 3
	}
	return 0
}

func cellValid(gameBoard *[9][9]int, value int, y int, x int) bool {
	oldVal := gameBoard[y-1][x-1]
	gameBoard[y-1][x-1] = value
	row := getRow(gameBoard, y)
	col := getCol(gameBoard, x)
	block := getBlock(gameBoard, x, y)
	possible := unitValid(row) && unitValid(col) && unitValid(block)
	gameBoard[y-1][x-1] = oldVal
	return possible
}

func solveBoard(gameBoard *[9][9]int) {
	for row := 1; row <= 9; row++ {
		for col := 1; col <= 9; col++ {
			if gameBoard[row-1][col-1] == 0 {
				for value := 1; value <= 9; value++ {
					if cellValid(gameBoard, value, row, col) {
						gameBoard[row-1][col-1] = value
						solveBoard(gameBoard)
						gameBoard[row-1][col-1] = 0
					}
				}
				return
			}
		}
	}
	printBoard(gameBoard)
	return
}

func main() {
	board := [9][9]int{}
	memoboard := [9][9]int{}
	check_board := [9]int{}
	gameBoard := Recup(board, memoboard, check_board)
	solveBoard(&gameBoard)
}

func Recup(board [9][9]int, memoboard [9][9]int, check_board [9]int) [9][9]int { // incrementer le tableau, need une fonction tcheck pour le full tableau
	boardline := [9]int{} ///$///$ variable
	end := false          ///$
	for i := 1; i <= 9; i++ {
		for index, line := range os.Args[i] {
			check_board = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
			if line >= rune('1') && line <= rune('9') {
				boardline[index] = int(line - 48)
			} else {
				boardline[index] = 0
			}
			check_board = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0} // ajouter un check du board complet en fin d'increment des valeurs
			for _, value := range boardline {               // les erreurs dues à la recupération par os.Arg sur les colonnes ne sont actuellement pas signalées
				if value > 0 && value <= 9 && !end {
					check_board[value-1]++
				}
			}
		}
		for it := 0; it < 9; it++ {
			if check_board[it] > 1 {
				fmt.Print("eror")
				break
			}
			// fmt.Print(check_board)
			board[i-1] = boardline
		}
	}
	fmt.Print(board)
	return board
}
