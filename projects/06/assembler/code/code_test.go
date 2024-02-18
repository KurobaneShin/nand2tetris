package code

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCode(t *testing.T) {
	codewriter := New("test.asm")

	defer codewriter.Close()

	assert.FileExists(t, "test.asm")
	os.Remove("test.asm")
	assert.NoFileExists(t, "test.asm")
}

func TestWriteAInstruction(t *testing.T) {
	codewriter := New("test.asm")

	defer codewriter.Close()

	WriteAInstruction(codewriter, 2)

	assert.FileExists(t, "test.asm")
	file, err := os.Open("test.asm")
	assert.NoError(t, err)

	defer file.Close()

	reader := bufio.NewReader(file)

	line, _, err := reader.ReadLine()

	assert.NoError(t, err)

	assert.Equal(t, len(line), 16)
	assert.Equal(t, "0000000000000010", string(line))

	assert.Equal(t, "0", string(line[0]))

	os.Remove("test.asm")
	assert.NoFileExists(t, "test.asm")
}

func TestWriteCInstruction(t *testing.T) {
	codewriter := New("test.asm")

	defer codewriter.Close()

	WriteCInstruction(codewriter, "D", "A", "JGT")

	assert.FileExists(t, "test.asm")
	file, err := os.Open("test.asm")
	assert.NoError(t, err)

	defer file.Close()

	reader := bufio.NewReader(file)

	line, _, err := reader.ReadLine()

	assert.NoError(t, err)

	assert.Equal(t, len(line), 16)
	assert.Equal(t, "1110110000010001", string(line))

	assert.Equal(t, "111", string(line[0:3]))

	os.Remove("test.asm")
	assert.NoFileExists(t, "test.asm")
}

func TestWriteCInstructionWithEmptyDestAndEmptyJump(t *testing.T) {
	codewriter := New("test.asm")

	defer codewriter.Close()

	WriteCInstruction(codewriter, "", "A", "")

	assert.FileExists(t, "test.asm")
	file, err := os.Open("test.asm")
	assert.NoError(t, err)

	defer file.Close()

	reader := bufio.NewReader(file)

	line, _, err := reader.ReadLine()

	assert.NoError(t, err)

	assert.Equal(t, len(line), 16)
	assert.Equal(t, "1110110000000000", string(line))

	assert.Equal(t, "111", string(line[0:3]))

	os.Remove("test.asm")
	assert.NoFileExists(t, "test.asm")
}

func TestWriteCAndAInstruction(t *testing.T) {
	codewriter := New("test.asm")

	defer codewriter.Close()

	WriteCInstruction(codewriter, "D", "A", "JGT")
	WriteAInstruction(codewriter, 2)

	assert.FileExists(t, "test.asm")
	file, err := os.Open("test.asm")
	assert.NoError(t, err)

	defer file.Close()

	reader := bufio.NewReader(file)
	assert.NoError(t, err)

	var lines []string

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		lines = append(lines, string(line))
	}

	assert.Equal(t, len(lines), 2)
	assert.Equal(t, "1110110000010001", lines[0])
	assert.Equal(t, "0000000000000010", lines[1])

	assert.NoError(t, err)

	os.Remove("test.asm")
	assert.NoFileExists(t, "test.asm")
}
