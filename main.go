package main

import (
	"fmt"
	"strconv"
)

var board *Board
var screen *Screen

func inputGameSize() (int, int) {
	var inputWidth string
	var inputHeight string
	var inputCheck string
	fmt.Print("게임의 가로 크기 입력 : ")
	fmt.Scanln(&inputWidth)
	fmt.Print("게임의 세로 크기 입력 : ")
	fmt.Scanln(&inputHeight)

	fmt.Printf("입력하신 게임의 크기는 가로 %s, 세로 %s 입니다. 맞습니까? (y/n) : ", inputWidth, inputHeight)
	fmt.Scanln(&inputCheck)

	w, _ := strconv.Atoi(inputWidth)
	h, _ := strconv.Atoi(inputHeight)
	return w, h
}

func initGame() {
	fmt.Println("Welcome to game 2048.")
	w, h := inputGameSize()
	board = NewBoard(w, h, nil)
	screen = NewScreen(board, ConsoleScreen{})
}

func doingGame() {
	fmt.Println("Start Game")
	var input string
	for {
		input = ""
		screen.draw()
		fmt.Println("1 : up")
		fmt.Println("2 : down")
		fmt.Println("3 : left")
		fmt.Println("4 : right")
		fmt.Println("1 ~ 4 중 하나를 입력해주세요.")
		fmt.Scanln(&input)
		isMove := false
		switch input {
		case "1":
			isMove = board.move(MOVE_UP)
		case "2":
			isMove = board.move(MOVE_DOWN)
		case "3":
			isMove = board.move(MOVE_LEFT)
		case "4":
			isMove = board.move(MOVE_RIGHT)
		default:
			fmt.Println("입력 오류!")
		}
		if isMove {
			board.addNumber()
		}
	}
}

func endGame() {
	fmt.Println("Game over.")
	//TODO print game result. (score)
}

func main() {
	initGame()
	doingGame()
	endGame()
}
