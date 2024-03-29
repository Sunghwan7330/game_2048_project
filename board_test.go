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
		{4, 4, 2, 2},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	expect := [][]int{
		{8, 4, 0, 0},
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
		{2, 2, 0, 0},
		{0, 2, 2, 0},
		{0, 4, 2, 0},
	}

	expect := [][]int{
		{4, 0, 0, 0},
		{4, 0, 0, 0},
		{4, 0, 0, 0},
		{4, 2, 0, 0},
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

func TestMoveBoardLeft6(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{16, 8, 4, 2},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	expect := [][]int{
		{16, 8, 4, 2},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	res := b.move(MOVE_LEFT)

	assert := assert.New(t)
	assert.Equal(res, false)
	assert.Equal(expect, b.board)
}

func TestMoveBoardLeft7(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{0, 2, 0, 4},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	expect := [][]int{
		{2, 4, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	res := b.move(MOVE_LEFT)

	assert := assert.New(t)
	assert.Equal(res, true)
	assert.Equal(expect, b.board)
}

func TestMoveBoardRight(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{0, 0, 2, 2},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	expect := [][]int{
		{0, 0, 0, 4},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	b.move(MOVE_RIGHT)

	assert := assert.New(t)
	assert.Equal(expect, b.board)
}

func TestMoveBoardRight2(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{0, 2, 2, 2},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	expect := [][]int{
		{0, 0, 2, 4},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	b.move(MOVE_RIGHT)

	assert := assert.New(t)
	assert.Equal(expect, b.board)
}

func TestMoveBoardRight3(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{2, 2, 4, 4},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	expect := [][]int{
		{0, 0, 4, 8},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	b.move(MOVE_RIGHT)

	assert := assert.New(t)
	assert.Equal(expect, b.board)
}

func TestMoveBoardRight4(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{2, 0, 2, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	expect := [][]int{
		{0, 0, 0, 4},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	b.move(MOVE_RIGHT)

	assert := assert.New(t)
	assert.Equal(expect, b.board)
}

func TestMoveBoardRight5(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{0, 0, 2, 4},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	expect := [][]int{
		{0, 0, 2, 4},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	res := b.move(MOVE_RIGHT)

	assert := assert.New(t)
	assert.Equal(res, false)
	assert.Equal(expect, b.board)
}

func TestMoveBoardRight6(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{2, 4, 8, 16},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	expect := [][]int{
		{2, 4, 8, 16},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	res := b.move(MOVE_RIGHT)

	assert := assert.New(t)
	assert.Equal(res, false)
	assert.Equal(expect, b.board)
}

func TestMoveBoardUp(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{2, 0, 0, 0},
		{2, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	expect := [][]int{
		{4, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	b.move(MOVE_UP)

	assert := assert.New(t)
	assert.Equal(expect, b.board)
}

func TestMoveBoardUp2(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{2, 0, 0, 0},
		{2, 0, 0, 0},
		{2, 0, 0, 0},
		{0, 0, 0, 0},
	}

	expect := [][]int{
		{4, 0, 0, 0},
		{2, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	b.move(MOVE_UP)

	assert := assert.New(t)
	assert.Equal(expect, b.board)
}

func TestMoveBoardUp3(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{4, 0, 0, 0},
		{4, 0, 0, 0},
		{2, 0, 0, 0},
		{2, 0, 0, 0},
	}

	expect := [][]int{
		{8, 0, 0, 0},
		{4, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	b.move(MOVE_UP)

	assert := assert.New(t)
	assert.Equal(expect, b.board)
}

func TestMoveBoardUp4(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{0, 0, 0, 0},
		{2, 0, 0, 0},
		{0, 0, 0, 0},
		{2, 0, 0, 0},
	}

	expect := [][]int{
		{4, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	b.move(MOVE_UP)

	assert := assert.New(t)
	assert.Equal(expect, b.board)
}

func TestMoveBoardUp5(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{4, 0, 0, 0},
		{2, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	expect := [][]int{
		{4, 0, 0, 0},
		{2, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	res := b.move(MOVE_UP)

	assert := assert.New(t)
	assert.Equal(res, false)
	assert.Equal(expect, b.board)
}

func TestMoveBoardUp6(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{2, 0, 0, 0},
		{4, 0, 0, 0},
		{8, 0, 0, 0},
		{16, 0, 0, 0},
	}

	expect := [][]int{
		{2, 0, 0, 0},
		{4, 0, 0, 0},
		{8, 0, 0, 0},
		{16, 0, 0, 0},
	}

	res := b.move(MOVE_UP)

	assert := assert.New(t)
	assert.Equal(res, false)
	assert.Equal(expect, b.board)
}

func TestMoveBoardDown(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{2, 0, 0, 0},
		{2, 0, 0, 0},
	}

	expect := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{4, 0, 0, 0},
	}

	b.move(MOVE_DOWN)

	assert := assert.New(t)
	assert.Equal(expect, b.board)
}

func TestMoveBoardDown2(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{0, 0, 0, 0},
		{2, 0, 0, 0},
		{2, 0, 0, 0},
		{2, 0, 0, 0},
	}

	expect := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{2, 0, 0, 0},
		{4, 0, 0, 0},
	}

	b.move(MOVE_DOWN)

	assert := assert.New(t)
	assert.Equal(expect, b.board)
}

func TestMoveBoardDown3(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{2, 0, 0, 0},
		{2, 0, 0, 0},
		{4, 0, 0, 0},
		{4, 0, 0, 0},
	}

	expect := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{4, 0, 0, 0},
		{8, 0, 0, 0},
	}

	b.move(MOVE_DOWN)

	assert := assert.New(t)
	assert.Equal(expect, b.board)
}

func TestMoveBoardDown4(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{2, 0, 0, 0},
		{0, 0, 0, 0},
		{2, 0, 0, 0},
		{0, 0, 0, 0},
	}

	expect := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{4, 0, 0, 0},
	}

	b.move(MOVE_DOWN)

	assert := assert.New(t)
	assert.Equal(expect, b.board)
}

func TestMoveBoardDown5(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{2, 0, 0, 0},
		{4, 0, 0, 0},
	}

	expect := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{2, 0, 0, 0},
		{4, 0, 0, 0},
	}

	res := b.move(MOVE_DOWN)

	assert := assert.New(t)
	assert.Equal(res, false)
	assert.Equal(expect, b.board)
}

func TestMoveBoardDown6(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{16, 0, 0, 0},
		{8, 0, 0, 0},
		{4, 0, 0, 0},
		{2, 0, 0, 0},
	}

	expect := [][]int{
		{16, 0, 0, 0},
		{8, 0, 0, 0},
		{4, 0, 0, 0},
		{2, 0, 0, 0},
	}

	res := b.move(MOVE_DOWN)

	assert := assert.New(t)
	assert.Equal(res, false)
	assert.Equal(expect, b.board)
}

func TestNotGameOver(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{16, 0, 0, 0},
		{8, 0, 0, 0},
		{4, 0, 0, 0},
		{2, 0, 0, 0},
	}
	assert := assert.New(t)
	assert.Equal(false, b.isGameOver())
}

func TestNotGameOver2(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{2, 2, 2, 2},
		{2, 2, 2, 2},
		{2, 2, 2, 2},
		{2, 2, 2, 2},
	}
	assert := assert.New(t)
	assert.Equal(false, b.isGameOver())
}

func TestNotGameOver3(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{2, 2, 2, 2},
		{4, 8, 16, 32},
		{32, 16, 8, 4},
		{4, 8, 16, 32},
	}
	assert := assert.New(t)
	assert.Equal(false, b.isGameOver())
}

func TestGameOver(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{16, 8, 4, 2},
		{2, 4, 8, 16},
		{16, 8, 4, 2},
		{2, 4, 8, 16},
	}
	assert := assert.New(t)
	assert.Equal(true, b.isGameOver())
}
func TestScore(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height, nil)
	b.board = [][]int{
		{2, 2, 2, 2},
		{2, 2, 2, 2},
		{2, 2, 2, 2},
		{2, 2, 2, 2},
	}

	res := b.move(MOVE_DOWN)

	assert := assert.New(t)
	assert.Equal(res, true)
	assert.Equal(32, b.score)
}
