package main

import "math/rand"

type Board struct {
	width  int
	height int
	board  [][]int
	rand   Rand
}

type Rand interface {
	Intn(n int) int
}

type defaultRand struct{}

func (r *defaultRand) Intn(n int) int {
	return rand.Intn(n)
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
	for {
		x := b.rand.Intn(b.width)
		y := b.rand.Intn(b.height)
		if b.board[x][y] == 0 {
			b.board[x][y] = generateRandomAddNumber()
			break
		}
	}
}

func NewBoard(width, height int, rand Rand) *Board {
	b := Board{}
	b.width = width
	b.height = height
	b.rand = rand

	if b.rand == nil {
		b.rand = &defaultRand{}
	}

	b.board = make([][]int, width)
	for i := 0; i < width; i++ {
		b.board[i] = make([]int, height)
	}

	for i := 0; i < 2; i++ {
		b.addNumber()
	}

	return &b
}
