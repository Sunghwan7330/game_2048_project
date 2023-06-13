package main

import "fmt"

type ScreenInterfece interface {
	draw(board Board)
}

type Screen struct {
	borad           *Board
	screenInterface ScreenInterfece
}

type ConsoleScreen struct{}

func drawLine(b Board) {
	for i := 0; i < b.width; i++ {
		fmt.Print("------")
	}
	fmt.Println("-")
}
func (screen ConsoleScreen) draw(board Board) {
	drawScore(board)
	drawGameBoard(board)
}

func drawScore(board Board) {
	fmt.Println("현재 점수 : ", board.score)
}

func drawGameBoard(board Board) {
	drawLine(board)
	for i := 0; i < board.height; i++ {
		for j := 0; j < board.width; j++ {
			var val string
			if board.board[i][j] == 0 {
				val = ""
			} else {
				val = fmt.Sprintf("%d", board.board[i][j])
			}
			fmt.Printf("|%5s", val)
		}
		fmt.Println("|")
		drawLine(board)
	}
}

func (screen Screen) draw() {
	screen.screenInterface.draw(*screen.borad)
}
func NewScreen(board *Board, screen ScreenInterfece) *Screen {
	s := Screen{}
	s.borad = board
	s.screenInterface = screen
	return &s
}
