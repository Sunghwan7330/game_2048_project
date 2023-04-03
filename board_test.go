package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockRand struct {
	number []int
}

func (r *mockRand) Intn(n int) int {
	vec := r.number
	deqVal := vec[0]
	r.number = vec[1:]

	return deqVal
}

func TestCreateBoard(t *testing.T) {
	width := 4
	height := 4

	b := NewBoard(width, height, &mockRand{number: []int{2, 2, 2, 2, 3, 3}})

	assert := assert.New(t)
	assert.Equal(width, b.width)
	assert.Equal(height, b.height)
	assert.Equal(width, len(b.board))
	assert.Equal(height, len(b.board[0]))

	zeroCnt := 0
	valueCnt := 0

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if b.board[i][j] == 0 {
				zeroCnt++
			} else {
				valueCnt++
			}
		}
	}
	assert.Equal(valueCnt, 2)
	assert.Equal(zeroCnt, (width*height)-2)

	assert.NotEqual(0, b.board[2][2])
	assert.NotEqual(0, b.board[3][3])
}

func TestMoveBoardLeft(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{2, 2, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	expect := [][]int{
		{4, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	b.move(MOVE_LEFT)

	assert := assert.New(t)
	assert.Equal(expect, b.board)
}

func TestMoveBoardLeft2(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{2, 2, 2, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	expect := [][]int{
		{4, 2, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	b.move(MOVE_LEFT)

	assert := assert.New(t)
	assert.Equal(expect, b.board)
}

func TestMoveBoardLeft3(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{4, 2, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	expect := [][]int{
		{4, 2, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	b.move(MOVE_LEFT)

	assert := assert.New(t)
	assert.Equal(expect, b.board)
}

func TestMoveBoardLeft4(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{0, 2, 0, 2},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	expect := [][]int{
		{4, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	b.move(MOVE_LEFT)

	assert := assert.New(t)
	assert.Equal(expect, b.board)
}

func TestMoveBoardLeft5(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{4, 2, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	expect := [][]int{
		{4, 2, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	res := b.move(MOVE_LEFT)

	assert := assert.New(t)
	assert.Equal(res, false)
	assert.Equal(expect, b.board)
}
