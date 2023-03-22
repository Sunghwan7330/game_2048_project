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
	popped := vec[len(vec)-1]
	vec = vec[:len(vec)-1]
	r.number = vec

	return popped
}

func (r *mockRand) set(n int) {
	r.number = append(r.number, n)
}

func TestCreateBoard(t *testing.T) {
	width := 4
	height := 4
	mRand := mockRand{}
	mRand.set(3)
	mRand.set(3)
	mRand.set(2)
	mRand.set(2)
	mRand.set(2)
	mRand.set(2)

	b := NewBoard(width, height, &mRand)

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
}
