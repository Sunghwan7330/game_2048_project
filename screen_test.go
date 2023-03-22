package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateConsoleScreen(t *testing.T) {
	board := NewBoard(4, 4, nil)

	s := NewScreen(board, ConsoleScreen{})

	assert := assert.New(t)
	assert.NotEqual(s, nil)
}

func TestBlankBoard4By4(t *testing.T) {
	board := NewBoard(4, 4, nil)
	board.board = [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	s := NewScreen(board, ConsoleScreen{})

	r, w, _ := os.Pipe()
	origStdout := os.Stdout
	os.Stdout = w

	buf := make([]byte, 1024)
	s.draw()
	r.Read(buf)

	// Restore
	os.Stdout = origStdout
	expected :=
		"-------------------------\n" +
			"|     |     |     |     |\n" +
			"-------------------------\n" +
			"|     |     |     |     |\n" +
			"-------------------------\n" +
			"|     |     |     |     |\n" +
			"-------------------------\n" +
			"|     |     |     |     |\n" +
			"-------------------------\n"

	length := bytes.IndexByte(buf, 0)
	assert := assert.New(t)
	assert.Equal(string(buf)[:length], expected)
}
