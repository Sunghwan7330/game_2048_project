package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBoard(t *testing.T) {
	b := createBoard(4, 4)

	assert := assert.New(t)
	assert.Equal(4, b.width)
	assert.Equal(4, b.height)
	assert.Equal(4, len(b.board))
	assert.Equal(4, len(b.board[0]))
}
