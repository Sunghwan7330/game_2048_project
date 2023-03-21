package main

import "math/rand"

type Board struct {
	width  int
	height int
	board  [][]int
}

func generateRandomAddNumber() int {
	//1/3 chance to generate 4.
	//2/3 chance to generate 2.
	v := rand.Intn(3)
	if v == 0 {
		return 4
	}

	return 2
}

func (b Board) addNumber() {
	x := rand.Intn(b.width)
	y := rand.Intn(b.height)

	b.board[x][y] = generateRandomAddNumber()
}

func createBoard(width, height int) *Board {
	b := Board{}
	b.width = width
	b.height = height

	b.board = make([][]int, width)
	for i := 0; i < width; i++ {
		b.board[i] = make([]int, height)
	}

	for i := 0; i < 2; i++ {
		b.addNumber()
	}

	return &b
}
