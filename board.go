package main

type Board struct {
	width  int
	height int
	board  [][]int
}

func createBoard(width, height int) *Board {
	b := Board{}
	b.width = width
	b.height = height

	b.board = make([][]int, width)
	for i := 0; i < width; i++ {
		b.board[i] = make([]int, height)
	}

	return &b
}
