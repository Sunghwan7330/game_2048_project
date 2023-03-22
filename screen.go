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

func (screen ConsoleScreen) draw(board Board) {
	fmt.Println("-------------------------")
	fmt.Println("|     |     |     |     |")
	fmt.Println("-------------------------")
	fmt.Println("|     |     |     |     |")
	fmt.Println("-------------------------")
	fmt.Println("|     |     |     |     |")
	fmt.Println("-------------------------")
	fmt.Println("|     |     |     |     |")
	fmt.Println("-------------------------")
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
