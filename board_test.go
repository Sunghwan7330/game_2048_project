package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBoard(t *testing.T) {
	width := 4
	height := 4
	b := NewBoard(width, height)

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
