package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateConsoleScreen(t *testing.T) {
	board := NewBoard(4, 4, nil)

	s := NewScreen(board, ConsoleScreen{})

	assert := assert.New(t)
	assert.NotEqual(s, nil)
}
