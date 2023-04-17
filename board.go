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

type MoveType int

const (
	MOVE_LEFT MoveType = iota
	MOVE_RIGHT
	MOVE_UP
	MOVE_DOWN
)

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

func (b *Board) move(moveType MoveType) bool {
	res := true
	switch moveType {
	case MOVE_LEFT:
		res = b.moveLeft()
	case MOVE_RIGHT:
		res = b.moveRight()
	case MOVE_UP:
		res = b.moveUp()
	case MOVE_DOWN:
		res = b.moveDown()
	}
	return res
}

func (b *Board) moveLeft() bool {
	isMove := false
	for i := 0; i < b.height; i++ {
		move_idx := 0
		cur_idx := 0
		for j := 1; j < b.width; j++ {
			if b.board[i][j] == 0 {
				continue
			}

			if b.board[i][cur_idx] == 0 {
				cur_idx = j
				continue
			}
			if b.board[i][cur_idx] == b.board[i][j] {
				b.board[i][move_idx] = b.board[i][cur_idx] * 2
				if move_idx != cur_idx {
					b.board[i][cur_idx] = 0
				}
				b.board[i][j] = 0
				move_idx += 1
				cur_idx = j
				isMove = true
			} else {
				if b.board[i][move_idx] != 0 {
					move_idx += 1
					cur_idx = j
				} else {
					b.board[i][move_idx] = b.board[i][cur_idx]
					b.board[i][cur_idx] = 0
					cur_idx = j
					move_idx += 1
					isMove = true
				}
			}
		}
		if (b.board[i][cur_idx] != 0) && (move_idx != cur_idx) {
			b.board[i][move_idx] = b.board[i][cur_idx]
			b.board[i][cur_idx] = 0
			isMove = true
		}
	}
	return isMove
}

func (b *Board) moveRight() bool {
	isMove := false
	for i := 0; i < b.height; i++ {
		move_idx := b.width - 1
		cur_idx := b.width - 1
		for j := b.width - 2; j >= 0; j-- {
			if b.board[i][j] == 0 {
				continue
			}

			if b.board[i][cur_idx] == 0 {
				cur_idx = j
				continue
			}
			if b.board[i][cur_idx] == b.board[i][j] {
				b.board[i][move_idx] = b.board[i][cur_idx] * 2
				if move_idx != cur_idx {
					b.board[i][cur_idx] = 0
				}
				b.board[i][j] = 0
				move_idx -= 1
				cur_idx = j
				isMove = true
			}
		}
		if (b.board[i][cur_idx] != 0) && (move_idx != cur_idx) {
			b.board[i][move_idx] = b.board[i][cur_idx]
			b.board[i][cur_idx] = 0
			isMove = true
		}
	}
	return isMove
}

func (b *Board) moveUp() bool {
	isMove := false
	for i := 0; i < b.width; i++ {
		move_idx := 0
		cur_idx := 0
		for j := 1; j < b.height; j++ {
			if b.board[j][i] == 0 {
				continue
			}

			if b.board[cur_idx][i] == 0 {
				cur_idx = j
				continue
			}
			if b.board[cur_idx][i] == b.board[j][i] {
				b.board[move_idx][i] = b.board[cur_idx][i] * 2
				if move_idx != cur_idx {
					b.board[cur_idx][i] = 0
				}
				b.board[j][i] = 0
				move_idx += 1
				cur_idx = j
				isMove = true
			}
		}
		if (b.board[cur_idx][i] != 0) && (move_idx != cur_idx) {
			b.board[move_idx][i] = b.board[cur_idx][i]
			b.board[cur_idx][i] = 0
			isMove = true
		}
	}
	return isMove
}

func (b *Board) moveDown() bool {
	isMove := false
	for i := 0; i < b.width; i++ {
		move_idx := b.height - 1
		cur_idx := b.height - 1
		for j := b.height - 2; j >= 0; j-- {
			if b.board[j][i] == 0 {
				continue
			}

			if b.board[cur_idx][i] == 0 {
				cur_idx = j
				continue
			}
			if b.board[cur_idx][i] == b.board[j][i] {
				b.board[move_idx][i] = b.board[cur_idx][i] * 2
				if move_idx != cur_idx {
					b.board[cur_idx][i] = 0
				}
				b.board[j][i] = 0
				move_idx -= 1
				cur_idx = j
				isMove = true
			}
		}
		if (b.board[cur_idx][i] != 0) && (move_idx != cur_idx) {
			b.board[move_idx][i] = b.board[cur_idx][i]
			b.board[cur_idx][i] = 0
			isMove = true
		}
	}
	return isMove
}
