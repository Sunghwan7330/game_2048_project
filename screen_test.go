package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StdoutDump struct {
	read       *os.File
	write      *os.File
	err        error
	origStdout *os.File
	buf        *[]byte
}

func (dump *StdoutDump) startDump(buffer *[]byte) {
	dump.read, dump.write, dump.err = os.Pipe()
	dump.origStdout = os.Stdout
	os.Stdout = dump.write
	dump.buf = buffer
}

func (dump *StdoutDump) endDump() {
	dump.read.Read(*dump.buf)

	// Restore
	os.Stdout = dump.origStdout
}

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

	buf := make([]byte, 1024)
	stdoutDump := StdoutDump{}
	stdoutDump.startDump(&buf)

	s.draw()
	stdoutDump.endDump()

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
